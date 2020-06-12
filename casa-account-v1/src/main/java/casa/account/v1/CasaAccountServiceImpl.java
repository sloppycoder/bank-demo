package casa.account.v1;

import com.datastax.oss.driver.api.core.CqlSession;
import com.datastax.oss.driver.api.core.cql.ResultSet;
import com.datastax.oss.driver.api.core.cql.Row;
import demo.bank.CasaAccount;
import demo.bank.CasaAccountServiceGrpc;
import demo.bank.GetCasaAccountRequest;
import io.grpc.stub.StreamObserver;
import io.micronaut.context.annotation.Value;
import io.opencensus.exporter.trace.jaeger.JaegerExporterConfiguration;
import io.opencensus.exporter.trace.jaeger.JaegerTraceExporter;
import io.opencensus.exporter.trace.stackdriver.StackdriverTraceConfiguration;
import io.opencensus.exporter.trace.stackdriver.StackdriverTraceExporter;
import io.opencensus.trace.AttributeValue;
import io.opencensus.trace.Tracing;
import io.opencensus.trace.config.TraceConfig;
import io.opencensus.trace.samplers.Samplers;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;
import javax.inject.Singleton;
import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.file.Paths;
import java.util.Map;

@Singleton
public class CasaAccountServiceImpl extends CasaAccountServiceGrpc.CasaAccountServiceImplBase {

  private static final Logger logger = LoggerFactory.getLogger(CasaAccountServiceImpl.class);

  @Value("${micronaut.application.name}")
  private String appName;

  @Value("${tracing.jaeger-endpoint:off}")
  private String jaegerThriftEndpoint;

  @Value("${tracing.use-stackdriver:false}")
  private String stackdriverFlag;

  @Value("${cassandra.instance:astra}")
  private String cassandraInstance;

  @Value("${cassandra.username:vino9}")
  private String dbUsername;

  @Value("${cassandra.password}")
  private String dbPassword;

  private CqlSession session;

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
    initializeTracing();
    initializeCassandraSession();
  }

  @PreDestroy
  public void cleanup() {
    if (this.session != null) {
      logger.info("closing CsqlSession");
      this.session.close();
    }
  }

  private void initializeCassandraSession() {
    CqlSession session;

    if ("astra".equalsIgnoreCase(cassandraInstance)) {
      session =
          CqlSession.builder()
              .withCloudSecureConnectBundle(Paths.get("secure-connect-vino9.zip"))
              .withAuthCredentials(dbUsername, dbPassword)
              .withKeyspace("vino9")
              .build();
    } else {
      session =
          CqlSession.builder()
              .addContactPoint(new InetSocketAddress("127.0.0.1", 9042))
              .withKeyspace("bank")
              .withLocalDatacenter("Cassandra")
              .build();
    }

    ResultSet rs = session.execute("select release_version from system.local");
    Row row = rs.one();
    String releaseVersion = row.getString("release_version");

    this.session = session;

    logger.info("connected to cassandra version {}", releaseVersion);
  }

  private void initializeTracing() {
    boolean exporterInitialized = false;

    exporterInitialized = initializeStackdriverExporter();

    if (!exporterInitialized) {
      exporterInitialized = initializeJaegerExporter();
    }

    if (!exporterInitialized) {
      logger.info("no exporter available, tracing not initialized");
      return;
    }

    TraceConfig traceConfig = Tracing.getTraceConfig();
    traceConfig.updateActiveTraceParams(
        traceConfig.getActiveTraceParams().toBuilder().setSampler(Samplers.alwaysSample()).build());
  }

  boolean initializeJaegerExporter() {
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

  boolean isStackdriverEnabled() {
    return "yes".equalsIgnoreCase(stackdriverFlag) || "true".equalsIgnoreCase(stackdriverFlag);
  }

  boolean initializeStackdriverExporter() {
    if (!isStackdriverEnabled()) {
      return false;
    }

    Map<String, AttributeValue> attributes =
        Map.of(
            "service", AttributeValue.stringAttributeValue("casa-account-v1"),
            "runtime", AttributeValue.stringAttributeValue(System.getProperty("java.version")));

    try {
      StackdriverTraceExporter.createAndRegister(
          StackdriverTraceConfiguration.builder()
              .setProjectId("vino9-276317")
              .setFixedAttributes(attributes)
              .build());
    } catch (IOException e) {
      logger.warn("unable to initialize stackdriver exporter", e);
      return false;
    }

    return true;
  }
}
