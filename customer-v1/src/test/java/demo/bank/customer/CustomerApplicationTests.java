package demo.bank.customer;

import demo.bank.*;
import io.grpc.health.v1.HealthCheckRequest;
import io.grpc.health.v1.HealthCheckResponse;
import io.grpc.health.v1.HealthGrpc;
import io.grpc.testing.GrpcCleanupRule;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.junit.Rule;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.annotation.DirtiesContext;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

@SpringBootTest(
        properties = {
                "grpc.server.inProcessName=test", // Enable inProcess server
                "grpc.server.port=-1", // Disable external server
                "grpc.client.inProcess.address=in-process:test" // Configure the client to connect to the
                // inProcess server
        })
@DirtiesContext
class CustomerApplicationTests {

    @Rule public final GrpcCleanupRule grpcCleanup = new GrpcCleanupRule();

    @GrpcClient("inProcess")
    private CustomerServiceGrpc.CustomerServiceBlockingStub stub;

    @GrpcClient("inProcess")
    private HealthGrpc.HealthBlockingStub healthService;

    @Test
    void context_loads() {}

    @Test
    void customerService_running() {
        String testCustomerId = "1234";
        GetCustomerRequest request = GetCustomerRequest.newBuilder().setCustomerId(testCustomerId).build();
        Customer response = stub.getCustomer(request);

        assertNotNull(response);
        assertEquals("1234", response.getCustomerId());
    }

    @Test
    void healthService_running() {
        HealthCheckRequest request = HealthCheckRequest.newBuilder().build();
        HealthCheckResponse response = healthService.check(request);

        assertNotNull(response);
        assertEquals("SERVING", response.getStatus().toString());
    }
}
