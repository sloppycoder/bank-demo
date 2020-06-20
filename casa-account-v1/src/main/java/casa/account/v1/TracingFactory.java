package casa.account.v1;

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
import java.util.Map;

@Factory
public class TracingFactory {
  private static final Logger logger = LoggerFactory.getLogger(TracingFactory.class);

  @Value("${micronaut.application.name}")
  private String appName;

  @Value("${tracing.jaeger-endpoint:off}")
  private String jaegerThriftEndpoint;

  @Value("${tracing.use-stackdriver:false}")
  private String stackdriverFlag;

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
              .setFixedAttributes(attributes)
              .build());
    } catch (IOException e) {
      logger.warn("unable to initialize stackdriver exporter", e);
      return false;
    }

    return true;
  }
}
