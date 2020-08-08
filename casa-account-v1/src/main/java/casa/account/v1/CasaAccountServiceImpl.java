package casa.account.v1;

import demo.bank.Balance;
import demo.bank.CasaAccount;
import demo.bank.CasaAccountServiceGrpc;
import demo.bank.GetCasaAccountRequest;
import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import io.opencensus.trace.AttributeValue;
import io.opencensus.trace.Span;
import io.opencensus.trace.Tracer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.inject.Inject;
import javax.inject.Singleton;
import java.util.Optional;
import java.util.Set;

@Singleton
public class CasaAccountServiceImpl extends CasaAccountServiceGrpc.CasaAccountServiceImplBase {

  private static final Logger logger = LoggerFactory.getLogger(CasaAccountServiceImpl.class);

  @Inject CasaAccountRepository repository;
  @Inject Tracer tracer;

  @Override
  public void getAccount(
      GetCasaAccountRequest req, StreamObserver<demo.bank.CasaAccount> responseObserver) {
    String accountId = req.getAccountId();

    tracer
        .getCurrentSpan()
        .putAttribute(
            "casa-account-v1/get_account/account_id",
            AttributeValue.stringAttributeValue(accountId));

    retrieveAccountFromDB(accountId, responseObserver);
  }

  private void retrieveAccountFromDB(
      String accountId, StreamObserver<demo.bank.CasaAccount> responseObserver) {
    logger.info("Retrieving CasaAccount details for {}", accountId);

    Span span = tracer.spanBuilder("db.query").startSpan();
    span.putAttribute(
        "casa-account-v1/get_account/retrieve_from_db/account_id",
        AttributeValue.stringAttributeValue(accountId));

    Optional<CasaAccountEntity> result = repository.findById(accountId);

    span.end();

    if (result.isEmpty()) {
      logger.info("casa account {} not found", accountId);

      responseObserver.onError(
          Status.INVALID_ARGUMENT.withDescription("casa account not found").asException());
      return;
    }

    CasaAccountEntity row = result.get();

    Balance bal1 =
        Balance.newBuilder()
            .setCreditFlag(false)
            .setType(Balance.Type.AVAILABLE)
            .setAmount(row.getBalance())
            .build();

    Balance bal2 =
        Balance.newBuilder()
            .setCreditFlag(false)
            .setType(Balance.Type.CURRENT)
            .setAmount(row.getBalance())
            .build();

    CasaAccount casaAccount =
        demo.bank.CasaAccount.newBuilder()
            .setAccountId(row.getAccountId())
            .setNickname(row.getNickName())
            .setProdCode(row.getProdCode())
            .setProdName(row.getProdName())
            .setCurrency(row.getCurrency())
            .setStatus(CasaAccount.Status.forNumber(row.getStatus()))
            .addAllBalances(Set.of(bal1, bal2))
            .build();

    responseObserver.onNext(casaAccount);
    responseObserver.onCompleted();
  }
}
