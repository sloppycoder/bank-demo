package casa.account.v1;

import demo.bank.CasaAccountServiceGrpc;
import io.grpc.ManagedChannel;
import io.micronaut.context.annotation.Bean;
import io.micronaut.context.annotation.Factory;
import io.micronaut.grpc.annotation.GrpcChannel;
import io.micronaut.grpc.server.GrpcServerChannel;

@Factory
public class TestClients {
    @Bean
    CasaAccountServiceGrpc.CasaAccountServiceBlockingStub blockingStub(
        @GrpcChannel(GrpcServerChannel.NAME) ManagedChannel channel) {
            return CasaAccountServiceGrpc.newBlockingStub(channel);
    }
}
