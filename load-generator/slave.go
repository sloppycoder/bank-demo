package main

import (
	"context"
	"io/ioutil"
	"load-generator/api"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	humanize "github.com/dustin/go-humanize"
	"github.com/myzhan/boomer"
)

const (
	IDFile = "ids.txt"
)

var (
	_conn *grpc.ClientConn
)

func idsFromFile(fname string) ([]string, int) {
	log.Print("reading ids from file...")

	content, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Print("cannot open ids.txt file")
		return nil, 0
	}

	lines := strings.Split(string(content), "\n")
	return lines, len(lines)
}

func randID() func() string {
	ids, size := idsFromFile(IDFile)
	log.Printf("using a pool of %s ids", humanize.Comma(int64(size)))

	rs := rand.NewSource(time.Now().UnixNano())
	rr := rand.New(rs)

	return func() string {
		id := strings.TrimSpace(ids[rr.Intn(size)])
		if id == "" {
			log.Print("Got an empty id, perhaps there's some bugs here")
			id = "0000"
		}
		return id
	}
}

func setupGrpcAPI(name string, getRandID func() string) func() {
	log.Print("setting up gRPC API test")

	return func() {
		start := time.Now()
		loginID := getRandID()
		err := callDashboardSvc(loginID)
		elapsed := time.Since(start)

		if err != nil {
			log.Printf("account %s got error %s", loginID, err)
			boomer.RecordFailure(
				"http",
				name,
				elapsed.Nanoseconds()/int64(time.Millisecond),
				err.Error(),
			)
		} else {
			boomer.RecordSuccess(
				"http",
				name,
				elapsed.Nanoseconds()/int64(time.Millisecond),
				int64(10),
			)
		}
	}
}

func callDashboardSvc(loginName string) error {
	conn, err := getConnection()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := api.NewDashboardServiceClient(conn)
	if _, err = client.GetDashboard(ctx, &api.GetDashboardRequest{
		LoginName: loginName,
	}); err != nil {
		return err
	}

	return nil
}

func getConnection() (*grpc.ClientConn, error) {
	// TODO: should add some retry mechanism for wait for Transient Error and Connecting states
	if _conn != nil && _conn.GetState() != connectivity.Shutdown {
		return _conn, nil
	}

	svcAddr := os.Getenv("DASHBOARD_SVC_ADDR")
	if svcAddr == "" {
		svcAddr = "dashboard:50051"
	}

	conn, err := grpc.Dial(svcAddr, grpc.WithInsecure())
	if err == nil {
		_conn = conn
		return _conn, nil
	}

	return nil, err
}

func main() {
	task := &boomer.Task{
		Name:   "dash",
		Weight: 100,
		Fn:     setupGrpcAPI("gRPC dashboard api", randID()),
	}

	boomer.Run(task)
}
