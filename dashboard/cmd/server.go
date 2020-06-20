package main

import (
	"dashboard/app"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
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

	s, _ := app.InitGrpcServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func main() {
	app.InitTracing()
	startServer()
}
