package app

import (
	"errors"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"contrib.go.opencensus.io/exporter/prometheus"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/zipkin"
	openzipkin "github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

const SampleRatio = 0.1

func initZipkinExporter() *zipkin.Exporter {
	collectorURL := os.Getenv("ZIPKIN_COLLECTOR_URL")
	if collectorURL == "" {
		log.Info("zipkin tracing cannot be initailzed as ZIPKIN_COLLECTOR_URL is not set")
		return nil
	}

	localEndpoint, _ := openzipkin.NewEndpoint("dashboard", "dashboard:0")
	reporter := httpreporter.NewReporter(collectorURL)
	zipkinExporter := zipkin.NewExporter(reporter, localEndpoint)

	if zipkinExporter != nil {
		log.Info("zipkin tracing initialized")
		return zipkinExporter
	}

	log.Info("zipkin tracing not initialized")
	return nil
}

func initStats(exporter view.Exporter) {
	view.SetReportingPeriod(StatsReportingPeriod * time.Second)
	view.RegisterExporter(exporter)

	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Warn("Error registering default server views")
	} else {
		log.Info("Registered default server views")
	}
}

func initStackdriverExporter() *stackdriver.Exporter {
	env := strings.ToLower(os.Getenv("USE_STACKDRIVER"))
	if env != "yes" && env != "true" {
		log.Info("stackdriver disabled by environment variable")
		return nil
	}

	for i := 1; i <= 3; i++ {
		exporter, err := stackdriver.NewExporter(stackdriver.Options{
			DefaultTraceAttributes: map[string]interface{}{
				"runtime": runtime.Version(),
				"service": "dashboard",
			},
		})
		if err == nil {
			log.Info("stackdriver exporter initialized")
			return exporter
		}

		log.Info("error trying to initialize stackdriver export, wait 5s and retry...")
		time.Sleep(time.Second * 5 * time.Duration(i))
	}

	log.Warn("unable to initialize stackdriver exporter after retrying, giving up")

	return nil
}

func initPrometheusExporter() (*prometheus.Exporter, error) {
	metricsPort := os.Getenv("METRIC_HTTP_ADDR")
	if metricsPort == "" {
		log.Info("stackdriver disabled by environment variable")
		return nil, errors.New("METRIC_HTTP_ADDR not set")
	}

	promExporter, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		log.Warnf("Failed to create Prometheus exporter: %v", err)
		return nil, err
	}

	// TODO: how to shut this down gracefully?
	go func() {
		log.Info("starting metrics http server at ", metricsPort)

		mux := http.NewServeMux()
		mux.Handle("/metrics", promExporter)
		if err := http.ListenAndServe(metricsPort, mux); err != nil {
			log.Warnf("Failed to run Prometheus /metrics endpoint: %v", err)
		}
	}()

	return promExporter, nil
}

func InitTracing() {
	exporterAvailable := false

	// try stackdriver first
	sdExporter := initStackdriverExporter()
	if sdExporter != nil {
		trace.RegisterExporter(sdExporter)
		initStats(sdExporter)

		exporterAvailable = true
	}

	// if stackdriver is not available, then zipkin
	if !exporterAvailable {
		zipkinExporter := initZipkinExporter()
		if zipkinExporter != nil {
			trace.RegisterExporter(zipkinExporter)
			if promExporter, err := initPrometheusExporter(); err == nil {
				initStats(promExporter)
			}

			exporterAvailable = true
		}
	}

	if exporterAvailable {
		trace.ApplyConfig(trace.Config{
			DefaultSampler: trace.ProbabilitySampler(SampleRatio),
		})
	}

	log.Info("tracing initialized")
}
