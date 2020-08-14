package main

import (
	"dashboard/app"
	"github.com/pkg/profile"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

func initLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	level, err := log.ParseLevel(os.Getenv("APP_LOG_LEVEL"))
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
}

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

	s, _ := app.InitGrpcServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func main() {
	if os.Getenv("ENABLE_PROFILING") == "true" {
		defer profile.Start().Stop()
	}

	initLogging()
	app.InitTracing()
	startServer()
}
