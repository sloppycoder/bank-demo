package casa.account.v1;

import com.datastax.oss.driver.api.core.CqlSession;
import com.datastax.oss.driver.api.core.cql.ResultSet;
import com.google.common.collect.ImmutableMap;
import io.micronaut.context.annotation.Bean;
import io.micronaut.context.annotation.Factory;
import io.micronaut.context.annotation.Value;
import io.opencensus.exporter.trace.jaeger.JaegerExporterConfiguration;
import io.opencensus.exporter.trace.jaeger.JaegerTraceExporter;
import io.opencensus.exporter.trace.stackdriver.StackdriverTraceConfiguration;
import io.opencensus.exporter.trace.stackdriver.StackdriverTraceExporter;
import io.opencensus.trace.AttributeValue;
import io.opencensus.trace.Tracer;
import io.opencensus.trace.Tracing;
import io.opencensus.trace.config.TraceConfig;
import io.opencensus.trace.samplers.Samplers;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Map;

// TODO: need a way to destroy cassandra connection during shutdown

@Factory
public class Initializers {
  private static final Logger logger = LoggerFactory.getLogger(Initializers.class);

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

  @Value("${cassandra.password:secret}")
  private String dbPassword;

  @Value("${cassandra.host:localhost}")
  private String localCassandraHost;

  @Value("${cassandra.port:9042}")
  private int localCassandraPort;

  @Bean
  public CqlSession initializeCassandraSession() {
    logger.info(
        "cassandra config instance={}, host={}, port={}, username={}, password={}****",
        cassandraInstance,
        localCassandraHost,
        localCassandraPort,
        dbUsername,
        dbPassword.substring(0, 1));

    String cwd = Paths.get(".").toAbsolutePath().normalize().toString();
    logger.info("CWD={}", cwd);

    // the secure bundle is in the parent directory when running locally outside
    // K8S environment.
    // when running inside K8s it is at the /astra directory mounted from a secret
    String secureBundle = "../secure-connect-vino9.zip";
    if (!Files.exists(Paths.get(secureBundle).toAbsolutePath())) {
        secureBundle = "/astra/secure-connect-vino9.zip";
    }

    CqlSession session;
    if ("astra".equalsIgnoreCase(cassandraInstance)) {
      session =
          CqlSession.builder()
              .withCloudSecureConnectBundle(Paths.get(secureBundle))
              .withAuthCredentials(dbUsername, dbPassword)
              .withKeyspace("vino9")
              .build();
    } else if ("local".equalsIgnoreCase(cassandraInstance)) {
      session =
          CqlSession.builder()
              .addContactPoint(new InetSocketAddress(localCassandraHost, localCassandraPort))
              .withAuthCredentials(dbUsername, dbPassword)
              .withKeyspace("vino9")
              .withLocalDatacenter("Cassandra")
              .build();
    } else {
      logger.info("cassandra not available");
      return null;
    }

    ResultSet rs = session.execute("select release_version from system.local");
    logger.info("connected to cassandra version {}", rs.one().getString("release_version"));

    return session;
  }

  @Bean
  public Tracer initializeTracing() {
    boolean exporterInitialized = false;

    exporterInitialized = initializeStackdriverExporter();

    if (!exporterInitialized) {
      exporterInitialized = initializeJaegerExporter();
    }

    if (!exporterInitialized) {
      logger.info("no exporter available, tracing not initialized");
    }

    TraceConfig traceConfig = Tracing.getTraceConfig();
    traceConfig.updateActiveTraceParams(
        traceConfig.getActiveTraceParams().toBuilder().setSampler(Samplers.alwaysSample()).build());

    return Tracing.getTracer();
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
        ImmutableMap.of(
            "service", AttributeValue.stringAttributeValue(appName),
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
