package casa.account.v1;

import demo.bank.Balance;
import demo.bank.CasaAccount;
import demo.bank.CasaAccountServiceGrpc;
import demo.bank.GetCasaAccountRequest;
import io.micronaut.test.annotation.MicronautTest;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import javax.inject.Inject;
import java.util.Set;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

@MicronautTest
public class CasaAccountServiceTest {
  @Inject CasaAccountServiceGrpc.CasaAccountServiceBlockingStub stub;

  // if needs a database setup to work, ignore for now
  @Disabled
  @Test
  void canGetDummyAccount() {
    String sureFireAccountId = "10001000";

    final GetCasaAccountRequest request =
        GetCasaAccountRequest.newBuilder().setAccountId(sureFireAccountId).build();

    CasaAccount response = stub.getAccount(request);

    assertNotNull(response);
    assertEquals(sureFireAccountId, response.getAccountId());
    assertEquals(CasaAccount.Status.ACTIVE, response.getStatus());

    Set.of(Balance.Type.values()).contains(response.getBalances(0).getType());
    Set.of(Balance.Type.values()).contains(response.getBalances(1).getType());
  }
}
