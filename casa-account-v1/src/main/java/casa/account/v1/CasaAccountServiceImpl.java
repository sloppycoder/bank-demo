package casa.account.v1;

import demo.bank.CasaAccount;
import demo.bank.CasaAccountServiceGrpc;
import demo.bank.GetCasaAccountRequest;
import io.grpc.stub.StreamObserver;
import io.micronaut.context.annotation.Value;
import io.opencensus.exporter.trace.jaeger.JaegerExporterConfiguration;
import io.opencensus.trace.Tracing;
import io.opencensus.trace.config.TraceConfig;
import io.opencensus.trace.samplers.Samplers;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.annotation.PostConstruct;
import javax.inject.Singleton;

import io.opencensus.exporter.trace.jaeger.JaegerTraceExporter;

@Singleton
public class CasaAccountServiceImpl extends CasaAccountServiceGrpc.CasaAccountServiceImplBase {

  private static final Logger logger = LoggerFactory.getLogger(CasaAccountServiceImpl.class);

  @Value("${micronaut.application.name}")
  private String appName;

  @Value("${tracing.jaeger-thrift-endpoint:off}")
  private String jaegerThriftEndpoint;

  @Value("${tracing.stackdriver:false}")
  private String stackdriverFlag;

  public void getAccount(GetCasaAccountRequest req, StreamObserver<CasaAccount> responseObserver) {
    String accountId = req.getAccountId();

    CasaAccount account =
        CasaAccount.newBuilder()
            .setAccountId(accountId)
            .setNickname("dummy-v1")
            .setStatus(CasaAccount.Status.DORMANT)
            .build();
    responseObserver.onNext(account);
    responseObserver.onCompleted();

    logger.info("Retrieving CasaAccount details for {}", accountId);
  }

  // this will only work if there is only one gRPC service in the application
  // it's probably better to move this out to a Factory class
  @PostConstruct
  public void initialize() {

    boolean jaegerAvailable = initializeJaegerExporter();
    boolean stackdriverAvailable = false;

    if (jaegerAvailable || stackdriverAvailable) {
      TraceConfig traceConfig = Tracing.getTraceConfig();
      traceConfig.updateActiveTraceParams(
          traceConfig
              .getActiveTraceParams()
              .toBuilder()
              .setSampler(Samplers.alwaysSample())
              .build());
    }
  }

  boolean  initializeJaegerExporter() {
    if (jaegerThriftEndpoint != null && jaegerThriftEndpoint.startsWith("http")) {
      JaegerTraceExporter.createAndRegister(
          JaegerExporterConfiguration.builder()
              .setServiceName(appName)
              .setThriftEndpoint(jaegerThriftEndpoint)
              .build());

      logger.info("jaeger exporter initialized");

      return true;
    }

    return false;
  }
}
