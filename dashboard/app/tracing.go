package app

import (
	"contrib.go.opencensus.io/exporter/jaeger"
	"contrib.go.opencensus.io/exporter/stackdriver"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"os"
	"strings"
	"time"
)

func initJaegerExporter() *jaeger.Exporter {
	svcAddr := os.Getenv("JAEGER_THRIFT_ENDPOINT")
	if svcAddr == "" {
		log.Info("jaeger exporter not initialized")
		return nil
	}

	exporter, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: svcAddr,
		Process: jaeger.Process{
			ServiceName: "dashboard",
			Tags: []jaeger.Tag{
				jaeger.StringTag("service", "dashboard"),
			},
		},
	})
	if err != nil {
		log.Warn("error trying to initialize jaeger exporter", err)
		return nil
	}

	log.Info("jaeger exporter initialized")

	return exporter
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
	if env != "yes" || env != "true" {
		log.Info("stackdriver disabled by environment variable")
		return nil
	}

	for i := 1; i <= 3; i++ {
		exporter, err := stackdriver.NewExporter(stackdriver.Options{})
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
	jaegerExporter := initJaegerExporter()
	if jaegerExporter != nil {
		trace.RegisterExporter(jaegerExporter)
	}

	sdExporter := initStackdriverExporter()
	if sdExporter != nil {
		trace.RegisterExporter(sdExporter)
		initStats(sdExporter)
	}

	if jaegerExporter != nil || sdExporter != nil {

		trace.ApplyConfig(trace.Config{
			DefaultSampler: trace.AlwaysSample(),
		})
	}

	log.Info("tracing initialized")
}
