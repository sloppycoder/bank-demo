package app

import (
	"context"
	pb "dashboard/api"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"os"
	"time"
)

type Server struct {
	pb.UnimplementedDashboardServiceServer
}

// GetDashboard implements DashboardService.GetDashboard
func (s *Server) GetDashboard(ctx context.Context, req *pb.GetDashboardRequest) (*pb.Dashboard, error) {
	log.Info("GetDashboard for user ", req.LoginName)

	dashboard := &pb.Dashboard{
		Customer: &pb.Customer{
			LoginName:  req.LoginName,
			CustomerId: req.LoginName,
		}}

	casa, err := getCasaAccount(ctx, req.LoginName)
	if err != nil {
		// perhaps should retry before returning dummy value?
		dashboard.Casa = []*pb.CasaAccount{}
		log.Warn("Unable to retrieve account detail")
	} else {
		dashboard.Casa = []*pb.CasaAccount{casa}
		dashboard.Customer.Name = dashboard.Casa[0].Nickname
	}

	return dashboard, nil
}

// TODO: should probably implement some kind of managed channel for better performance
func initCasaConnection(ctx context.Context) (*grpc.ClientConn, error) {
	addr := os.Getenv("CASA_SVC_ADDR")
	if addr == "" {
		addr = "casa-account:50051"
	}

	return grpc.DialContext(ctx, addr,
		grpc.WithInsecure(), grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
}

func getCasaAccount(ctx context.Context, accountId string) (*pb.CasaAccount, error) {
	conn, err := initCasaConnection(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	c := pb.NewCasaAccountServiceClient(conn)
	subctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	r, err := c.GetAccount(subctx, &pb.GetCasaAccountRequest{AccountId: accountId})
	if err != nil {
		log.Warn("unable to retrieve CasaAccount detail", err)
		return r, err
	}

	return r, nil
}

// implement grpc health check protocol.
func (s *Server) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s *Server) Watch(req *healthpb.HealthCheckRequest, ws healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
