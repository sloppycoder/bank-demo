package demo.bank.customer;

import brave.Tracing;
import brave.grpc.GrpcTracing;
import net.devh.boot.grpc.server.interceptor.GlobalServerInterceptorConfigurer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class GlobalInterceptorConfiguration {

    @Bean
    public GlobalServerInterceptorConfigurer globalInterceptorConfigurerAdapter() {
        return registry -> registry.addServerInterceptors(new LogGrpcInterceptor());
    }

    // enable using grpc-trace-bin header in order to be
    // compatible with opencensus based tracer
    // https://github.com/openzipkin/brave/blob/master/instrumentation/grpc/README.md
    @Bean
    public GrpcTracing grpcTracing(final Tracing tracing) {
        return GrpcTracing.newBuilder(tracing)
                .grpcPropagationFormatEnabled(true).build();
    }
}