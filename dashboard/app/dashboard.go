package app

import (
	"context"
	"dashboard/api"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"golang.org/x/sync/errgroup"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	StatsReportingPeriod = 60
)

type ServerContext struct {
	casaSvcConn, custSvcConn *grpc.ClientConn
	mockCustSvc, mockCasaSvc bool
	timeout                  time.Duration
}

type Server struct {
	context *ServerContext
}

func newServerContext() *ServerContext {
	serverCtx := &ServerContext{
		mockCustSvc: true,
		mockCasaSvc: true,
		timeout:     5 * time.Second,
	}

	if os.Getenv("USE_CUST_SVC") != "false" {
		conn, err := newCustomerConnection()
		if err != nil {
			log.Warnf("unable to create connection to customer service, %v", err)
		} else {
			serverCtx.custSvcConn = conn
			serverCtx.mockCustSvc = false
		}
	}

	if os.Getenv("USE_CASA_SVC") != "false" {
		conn, err := newCasaConnection()
		if err != nil {
			log.Warnf("unable to create connection to casa-account service, %v", err)
		} else {
			serverCtx.casaSvcConn = conn
			serverCtx.mockCasaSvc = false
		}
	}

	return serverCtx
}

// helper for logging with trace and span id
func info(ctx context.Context, args ...interface{}) {
	if span := trace.FromContext(ctx); span != nil {
		log.WithFields(log.Fields{
			"traceId": span.SpanContext().TraceID.String(),
			"spanId":  span.SpanContext().SpanID.String(),
		}).Info(args...)
	}
}

func newCustomerConnection() (*grpc.ClientConn, error) {
	log.Info("Creating new connection for Customer Service")

	addr := os.Getenv("CUSTOMER_SVC_ADDR")
	if addr == "" {
		addr = "customer:50051"
	}

	return grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
}

func newCasaConnection() (*grpc.ClientConn, error) {
	log.Info("Creating new connection for Casa Service")

	addr := os.Getenv("CASA_SVC_ADDR")
	if addr == "" {
		addr = "casa-account:50051"
	}

	return grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
}

// GetDashboard implements DashboardService.GetDashboard.
func (s *Server) GetDashboard(ctx context.Context, req *api.GetDashboardRequest) (*api.Dashboard, error) {
	user := req.LoginName
	info(ctx, "GetDashboard for user ", user)

	if span := trace.FromContext(ctx); span != nil {
		span.AddAttributes(
			trace.StringAttribute("dashboard/get_dashboard/login_name", user))
	}

	dashboard := &api.Dashboard{}
	serverCtx := s.context

	errs, ctx := errgroup.WithContext(ctx)
	errs.Go(func() error {
		if serverCtx.mockCustSvc {
			dashboard.Customer = &api.Customer{LoginName: "skip"}
			return nil
		}
		return getCustomer(ctx, s.context, req.LoginName, dashboard)
	})
	errs.Go(func() error {
		if serverCtx.mockCasaSvc {
			dashboard.Casa = []*api.CasaAccount{{AccountId: "skip"}}
			return nil
		}
		return getCasaAccount(ctx, s.context, req.LoginName, dashboard)
	})
	if err := errs.Wait(); err != nil {
		return nil, status.New(codes.Code(code.Code_NOT_FOUND), "unable to load dashboard").Err()
	}

	return dashboard, nil
}

func getCasaAccount(ctx context.Context, serverCtx *ServerContext, accountID string, dashboard *api.Dashboard) error {
	c := api.NewCasaAccountServiceClient(serverCtx.casaSvcConn)
	subctx, cancel := context.WithTimeout(ctx, serverCtx.timeout)
	defer cancel()

	casa, err := c.GetAccount(subctx, &api.GetCasaAccountRequest{AccountId: accountID})
	if err == nil {
		dashboard.Casa = []*api.CasaAccount{casa}
	} else {
		info(ctx, "error retrieving casa account detail", err)
	}

	return err
}

func getCustomer(ctx context.Context, serverCtx *ServerContext, custID string, dashboard *api.Dashboard) error {
	c := api.NewCustomerServiceClient(serverCtx.custSvcConn)
	subctx, cancel := context.WithTimeout(ctx, serverCtx.timeout)
	defer cancel()

	cust, err := c.GetCustomer(subctx, &api.GetCustomerRequest{CustomerId: custID})
	if err == nil {
		dashboard.Customer = cust
	} else {
		info(ctx, "error retrieving customer detail", err)
	}
	return err
}

func InitGrpcServer() (*grpc.Server, error) {
	view.SetReportingPeriod(StatsReportingPeriod * time.Second)

	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Warn("Unable to register views for server stats ", err)
	}

	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		log.Warn("Unable to register views for client stats ", err)
	}

	serverCtx := newServerContext()
	svc := &Server{serverCtx}

	s := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
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
