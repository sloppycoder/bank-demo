package casa.account.v1;

import com.datastax.oss.driver.api.core.CqlSession;
import com.datastax.oss.driver.api.core.cql.ResultSet;
import com.datastax.oss.driver.api.core.cql.Row;
import com.datastax.oss.driver.api.core.cql.SimpleStatement;
import com.datastax.oss.driver.api.core.data.UdtValue;
import demo.bank.Balance;
import demo.bank.CasaAccount;
import demo.bank.CasaAccountServiceGrpc;
import demo.bank.GetCasaAccountRequest;
import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import io.opencensus.trace.Span;
import io.opencensus.trace.Tracer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.inject.Inject;
import javax.inject.Singleton;
import java.util.Set;

import static java.util.stream.Collectors.toSet;

@Singleton
public class CasaAccountServiceImpl extends CasaAccountServiceGrpc.CasaAccountServiceImplBase {

  private static final Logger logger = LoggerFactory.getLogger(CasaAccountServiceImpl.class);

  @Inject private CqlSession session;
  @Inject Tracer tracer;

  @Override
  public void getAccount(GetCasaAccountRequest req, StreamObserver<CasaAccount> responseObserver) {
    String accountId = req.getAccountId();

    tracer
        .getCurrentSpan()
        .addAnnotation(String.format("casa-account.get_account.account_id=%s", accountId));

    retrieveAccountFromDB(accountId, responseObserver);
  }

  private void retrieveAccountFromDB(
      String accountId, StreamObserver<CasaAccount> responseObserver) {
    logger.info("Retrieving CasaAccount details for {}", accountId);

    Span span = tracer.spanBuilder("Cassandra.query").startSpan();
    span.addAnnotation(String.format("casa-account.get_account.account_id=%s", accountId));

    ResultSet rs =
        session.execute(
            SimpleStatement.builder("SELECT * FROM casa_account WHERE account_id = ?")
                .addPositionalValue(accountId)
                .build());

    Row row = rs.one();

    span.end();

    if (row == null) {
      logger.info("casa account {} not found", accountId);

      responseObserver.onError(
          Status.INVALID_ARGUMENT.withDescription("casa account not found").asException());
      return;
    }

    Set<Balance> balances =
        row.getSet("balances", UdtValue.class).stream()
            .map(
                bal ->
                    Balance.newBuilder()
                        .setAmount(bal.getFloat("amount"))
                        .setCreditFlag(bal.getBoolean("credit"))
                        .setType(Balance.Type.forNumber(bal.getShort("type")))
                        .build())
            .collect(toSet());

    CasaAccount account =
        CasaAccount.newBuilder()
            .setAccountId(row.getString("account_id"))
            .setNickname(row.getString("nickname"))
            .setProdCode(row.getString("prod_code"))
            .setProdName(row.getString("prod_name"))
            .setStatus(CasaAccount.Status.forNumber(0))
            .addAllBalances(balances)
            .build();

    responseObserver.onNext(account);
    responseObserver.onCompleted();
  }
}
