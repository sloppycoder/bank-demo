package app

import (
	"os"
	"runtime"
	"strings"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/zipkin"
	openzipkin "github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

func initZipkinExporter() *zipkin.Exporter {
	collectorUrl := os.Getenv("ZIPKIN_COLLECTOR_URL")
	if collectorUrl == "" {
		log.Info("zipkin tracing cannot be initailzed as ZIPKIN_COLLECTOR_URL is not set")
		return nil
	}

	localEndpoint, _ := openzipkin.NewEndpoint("dashboard", "dashboard:0")
	reporter := httpreporter.NewReporter(collectorUrl)
	zipkinExporter := zipkin.NewExporter(reporter, localEndpoint)

	if zipkinExporter != nil {
		log.Info("zipkin tracing initialized")
		return zipkinExporter
	}

	log.Info("zipkin tracing not initialized")
	return nil
}

func initStats(exporter *stackdriver.Exporter) {
	view.SetReportingPeriod(60 * time.Second)
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

func InitTracing() {
	exporterAvailable := false

	// try stackdriver first
	sdExporter := initStackdriverExporter()
	if sdExporter != nil {
		trace.RegisterExporter(sdExporter)
		initStats(sdExporter)

		exporterAvailable = true
	}

	// if stackdriver is not available, then jaeger
	if !exporterAvailable {
		zipkinExporter := initZipkinExporter()
		if zipkinExporter != nil {
			trace.RegisterExporter(zipkinExporter)

			exporterAvailable = true
		}
	}

	if exporterAvailable {
		trace.ApplyConfig(trace.Config{
			DefaultSampler: trace.AlwaysSample(),
		})
	}

	log.Info("tracing initialized")
}
