package app

import (
	"context"
	api "dashboard/api"
	"golang.org/x/sync/errgroup"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/connectivity"
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

var (
	_casaConn, _custConn *grpc.ClientConn
)

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

func getCustomerConnection(ctx context.Context) (*grpc.ClientConn, error) {
	// TODO: should add some retry mechansim for wait for Transient Error and Connecting states
	if _custConn != nil && _custConn.GetState() != connectivity.Shutdown {
		info(ctx, "_custConn state = ", _custConn.GetState().String())
		return _custConn, nil
	}

	conn, err := newCustomerConnection(ctx)
	if err == nil {
		_custConn = conn
		return _custConn, nil
	}

	warn(ctx, "unable to create connection to customer service", err)
	return nil, err
}

func newCustomerConnection(ctx context.Context) (*grpc.ClientConn, error) {
	info(ctx, "Creating new connection for Customer Service")

	addr := os.Getenv("CUSTOMER_SVC_ADDR")
	if addr == "" {
		addr = "customer:50051"
	}

	return grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))

}

func getCasaConnection(ctx context.Context) (*grpc.ClientConn, error) {
	// TODO: should add some retry mechansim for wait for Transient Error and Connecting states
	if _casaConn != nil && _casaConn.GetState() != connectivity.Shutdown {
		info(ctx, "_casaConn state = ", _casaConn.GetState().String())
		return _casaConn, nil
	}

	conn, err := newCasaConnection(ctx)
	if err == nil {
		_casaConn = conn
		return _casaConn, nil
	}

	warn(ctx, "unable to create connection to casa account service", err)
	return nil, err
}

func newCasaConnection(ctx context.Context) (*grpc.ClientConn, error) {
	info(ctx, "Creating new connection for Casa Service")

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
			trace.StringAttribute("get_dashboard.login_name", user))
	}

	dashboard := &api.Dashboard{}

	errs, ctx := errgroup.WithContext(ctx)
	errs.Go(func() error {
		return getCustomer(ctx, req.LoginName, dashboard)
	})
	errs.Go(func() error {
		return getCasaAccount(ctx, req.LoginName, dashboard)
	})
	if err := errs.Wait(); err != nil {
		return nil, status.New(codes.Code(code.Code_NOT_FOUND), "unable to load dashboard").Err()
	}

	return dashboard, nil
}

func getCasaAccount(ctx context.Context, accountId string, dashboard *api.Dashboard) error {
	conn, err := getCasaConnection(ctx)
	if err != nil {
		return err
	}

	c := api.NewCasaAccountServiceClient(conn)
	subctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	casa, err := c.GetAccount(subctx, &api.GetCasaAccountRequest{AccountId: accountId})
	if err == nil {
		dashboard.Casa = []*api.CasaAccount{casa}
	} else {
		info(ctx, "error retrieving casa account detail", err)
	}

	return err
}

func getCustomer(ctx context.Context, custId string, dashboard *api.Dashboard) error {
	conn, err := getCustomerConnection(ctx)
	if err != nil {
		return err
	}

	c := api.NewCustomerServiceClient(conn)
	subctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	cust, err := c.GetCustomer(subctx, &api.GetCustomerRequest{CustomerId: custId})
	if err == nil {
		dashboard.Customer = cust
	} else {
		info(ctx, "error retrieving customer detail", err)
	}
	return err
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
