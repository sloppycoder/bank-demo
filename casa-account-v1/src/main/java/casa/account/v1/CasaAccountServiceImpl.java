package casa.account.v1;

import demo.bank.CasaAccount;
import demo.bank.CasaAccountServiceGrpc;
import demo.bank.GetCasaAccountRequest;
import io.grpc.stub.StreamObserver;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.inject.Singleton;

@Singleton
public class CasaAccountServiceImpl extends CasaAccountServiceGrpc.CasaAccountServiceImplBase {

    private static final Logger logger = LoggerFactory.getLogger(CasaAccountServiceImpl.class);

    public void getAccount(GetCasaAccountRequest req, StreamObserver<CasaAccount> responseObserver) {
        String accountId = req.getAccountId();

        CasaAccount account = CasaAccount.newBuilder()
                .setAccountId(accountId)
                .setNickname("dummy-v1")
                .setStatus(CasaAccount.Status.DORMANT)
                .build();
        responseObserver.onNext(account);
        responseObserver.onCompleted();

        logger.info("Retrieving CasaAccount details for {}", accountId);
    }
}

