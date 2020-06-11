package casa.account.v1;

import demo.bank.CasaAccount;
import demo.bank.CasaAccountServiceGrpc;
import demo.bank.GetCasaAccountRequest;
import io.micronaut.test.annotation.MicronautTest;
import org.junit.jupiter.api.Test;

import javax.inject.Inject;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

@MicronautTest
public class CasaAccountServiceTest {
    @Inject
    CasaAccountServiceGrpc.CasaAccountServiceBlockingStub stub;

    @Test
    void canGetDummyAccount() {
        final GetCasaAccountRequest request = GetCasaAccountRequest
                                                .newBuilder()
                                                .setAccountId("dummy")
                                                .build();

        CasaAccount response = stub.getAccount(request);

        assertNotNull(response);
        assertEquals("dummy-v1", response.getNickname());
    }
}
