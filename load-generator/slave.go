package main

import (
	"context"
	"errors"
	"load-generator/api"
	"load-generator/grpcpool"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"google.golang.org/grpc"

	humanize "github.com/dustin/go-humanize"
	"github.com/myzhan/boomer"
)

var (
	_conn *grpc.ClientConn
)

func connectToDB() *sql.DB {
	connStr := os.Getenv("MYSQL_CONN_STRING")
	if connStr == "" {
		connStr = "demo:demo@tcp(192.168.39.1:3306)/demo?tls=skip-verify&autocommit=true"
	}
	log.Printf("MYSQL conn string is %s", connStr)

	for i := 0; i < 3; i++ {
		db, _ := sql.Open("mysql", connStr)
		_, err := db.Query("select version()")
		if err == nil {
			return db
		}

		log.Print("Unable to connect to database, retrying...")
		time.Sleep(2 * time.Second)
	}

	log.Print("Unable to connect to database")
	return nil
}

func idsFromDB() ([]string, int) {
	if os.Getenv("USE_DUMMY_ID") == "true" {
		return []string{"10001000"}, 1
	}

	log.Print("reading ids from database...")
	db := connectToDB()
	rows, err := db.Query("select account_id from demo.casa_account")
	defer rows.Close()

	var lines []string
	var id string
	if err == nil {
		for rows.Next() {
			err := rows.Scan(&id)
			if err != nil {
				log.Fatal(err)
			}
			lines = append(lines, id)
		}
	}

	return lines, len(lines)
}

func randID() func() string {
	ids, size := idsFromDB()
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

func setupGrpcAPI(api string, getRandID func() string) func() {
	log.Printf("setting up gRPC API %s for test", api)

	pool, err := NewConnectionPool(apiName)
	if err != nil {

	}

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		conn, err := pool.Get(ctx)
		defer conn.Close()

		start := time.Now()
		loginID := getRandID()
		err = invokeApi(ctx, conn.ClientConn, apiName, loginID)
		elapsed := time.Since(start)

		if err != nil {
			log.Printf("account %s got error %s", loginID, err)
			boomer.RecordFailure(
				"gRPC",
				api,
				elapsed.Nanoseconds()/int64(time.Millisecond),
				err.Error(),
			)
		} else {
			boomer.RecordSuccess(
				"gRPC",
				api,
				elapsed.Nanoseconds()/int64(time.Millisecond),
				int64(10),
			)
		}
	}
}

func invokeApi(ctx context.Context, conn *grpc.ClientConn, apiName, id string) error {
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

func NewConnectionPool(apiName string) (*grpcpool.Pool, error) {
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

	var factory grpcpool.Factory
	factory = func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(svcAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to start gRPC connection: %v", err)
		}
		log.Printf("Connected to %s", svcAddr)
		return conn, err
	}

	size ,err := strconv.Atoi(os.Getenv("POOLSIZE"))
	if err != nil {
		size = 6
	}

	return grpcpool.New(factory, size, size, time.Second)
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
