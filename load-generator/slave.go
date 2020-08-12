package main

import (
	"context"
	"errors"
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
	if os.Getenv("USE_DUMMY_ID") == "true" {
		return []string{"10001000"}, 1
	}

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
			// log.Print("Got an empty id, perhaps there's some bugs here")
			id = "10001000"
		}
		return id
	}
}

func setupGrpcAPI(apiName string, getRandID func() string) func() {
	log.Printf("setting up gRPC API %s for test", apiName)

	return func() {
		conn, err := getConnection(apiName)
		if err != nil {
			log.Panicf("Unable to connect to ")
		}

		start := time.Now()
		loginID := getRandID()
		err = invokeApi(conn, apiName, loginID)
		elapsed := time.Since(start)

		if err != nil {
			log.Printf("account %s got error %s", loginID, err)
			boomer.RecordFailure(
				"gRPC",
				apiName,
				elapsed.Nanoseconds()/int64(time.Millisecond),
				err.Error(),
			)
		} else {
			boomer.RecordSuccess(
				"gRPC",
				apiName,
				elapsed.Nanoseconds()/int64(time.Millisecond),
				int64(10),
			)
		}
	}
}

func invokeApi(conn *grpc.ClientConn, apiName, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	switch apiName {
	case "dashboard":
		client := api.NewDashboardServiceClient(conn)
		_, err = client.GetDashboard(ctx, &api.GetDashboardRequest{LoginName: id})
	case "customer":
		client := api.NewCustomerServiceClient(conn)
		_, err = client.GetCustomer(ctx, &api.GetCustomerRequest{CustomerId: id})
	case "casa":
		client := api.NewCasaAccountServiceClient(conn)
		_, err = client.GetAccount(ctx, &api.GetCasaAccountRequest{AccountId: id})
	default:
		err = errors.New("invalid api")
	}
	return err
}

func getConnection(apiName string) (*grpc.ClientConn, error) {
	// TODO: should add some retry mechanism for wait for Transient Error and Connecting states
	if _conn != nil && _conn.GetState() != connectivity.Shutdown {
		return _conn, nil
	}

	svcAddr := os.Getenv("SVC_ADDR")
	if svcAddr == "" {
		switch apiName {
		case "customer":
			svcAddr = "customer:50051"
		case "casa":
			svcAddr = "casa-account:50051"
		default:
			svcAddr = "dashboard:50051"
		}
	}

	conn, err := grpc.Dial(svcAddr, grpc.WithInsecure())
	if err == nil {
		_conn = conn
		return _conn, nil
	}

	return nil, err
}

func main() {
	apiName := os.Getenv("TARGET_API")
	if apiName == "" {
		apiName = "dashboard"
	}

	task := &boomer.Task{
		Name:   "grpc",
		Weight: 100,
		Fn:     setupGrpcAPI(apiName, randID()),
	}

	boomer.Run(task)
}
