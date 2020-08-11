package main

import (
	"context"
	"load-generator/api"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

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
	log.Print("reading ids from database...")

	db := connectToDB()
	if db == nil {
		return []string{"10001000"}, 1
	}

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

	if len(lines) < 1 {
		return []string{"10001000"}, 1
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
