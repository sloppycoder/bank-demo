package app

import (
	"context"
	api "dashboard/api"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server struct {
	api.UnimplementedDashboardServiceServer
}

// helper for logging with trace and span id
func info(ctx context.Context, args ...interface{}) {
	if span := trace.FromContext(ctx); span != nil {
		log.WithFields(log.Fields{
			"traceId": span.SpanContext().TraceID.String(),
			"spanId":  span.SpanContext().SpanID.String(),
		}).Info(args)
	}
}

func warn(ctx context.Context, args ...interface{}) {
	if span := trace.FromContext(ctx); span != nil {
		log.WithFields(log.Fields{
			"traceId": span.SpanContext().TraceID.String(),
			"spanId":  span.SpanContext().SpanID.String(),
		}).Warn(args)
	}
}

// GetDashboard implements DashboardService.GetDashboard.
func (s *Server) GetDashboard(ctx context.Context, req *api.GetDashboardRequest) (*api.Dashboard, error) {
	span := trace.FromContext(ctx)

	user := req.LoginName
	info(ctx, "2. GetDashboard for user ", user)

	span.AddAttributes(
		trace.StringAttribute("get_dashboard.login_name", user))

	dashboard := &api.Dashboard{
		Customer: &api.Customer{
			LoginName:  req.LoginName,
			CustomerId: req.LoginName,
		}}

	casa, err := getCasaAccount(ctx, req.LoginName)
	if err != nil {
		// perhaps should retry before returning dummy value?
		dashboard.Casa = []*api.CasaAccount{}

		warn(ctx, "Unable to retrieve account detail")
	} else {
		dashboard.Casa = []*api.CasaAccount{casa}
		dashboard.Customer.Name = dashboard.Casa[0].Nickname
	}

	return dashboard, nil
}

// TODO: should probably implement some kind of managed channel for better performance.
func initCasaConnection(ctx context.Context) (*grpc.ClientConn, error) {
	addr := os.Getenv("CASA_SVC_ADDR")
	if addr == "" {
		addr = "casa-account:50051"
	}

	return grpc.DialContext(ctx, addr,
		grpc.WithInsecure(), grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
}

func getCasaAccount(ctx context.Context, accountId string) (*api.CasaAccount, error) {
	conn, err := initCasaConnection(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	c := api.NewCasaAccountServiceClient(conn)
	subctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	r, err := c.GetAccount(subctx, &api.GetCasaAccountRequest{AccountId: accountId})
	if err != nil {
		warn(subctx, "unable to retrieve CasaAccount detail", err)
		return r, err
	}

	return r, nil
}

func InitGrpcServer() (*grpc.Server, error) {
	view.SetReportingPeriod(60 * time.Second)

	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Warn("Unable to register views for stats ", err)
	}

	s := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	svc := &Server{}
	api.RegisterDashboardServiceServer(s, svc)
	health.RegisterHealthServer(s, svc)
	reflection.Register(s)

	return s, nil
}

// implement grpc health check protocol.
func (s *Server) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}

func (s *Server) Watch(req *health.HealthCheckRequest, ws health.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
