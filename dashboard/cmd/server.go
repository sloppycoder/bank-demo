package main

import (
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"net"
	"os"
	"time"

	api "dashboard/api"
	app "dashboard/app"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func startServer() {
	addr := os.Getenv("GRPC_LISTEN_ADDR")
	if addr == "" {
		addr = ":50051"
	}

	log.Info("starting grpc server on ", addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	view.SetReportingPeriod(60 * time.Second)
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Warn("Unable to register views for stats ", err)
	}

	s := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	svc := &app.Server{}
	api.RegisterDashboardServiceServer(s, svc)
	health.RegisterHealthServer(s, svc)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}

func main() {
	app.InitTracing()
	startServer()
}
