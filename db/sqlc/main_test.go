package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/hardik-kansal/go-api-deploy/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var queries *SQLStore
func TestMain(m *testing.M) {
	config,err:=util.LoadConfig("../../")
	if err!=nil{
		log.Fatalf("error grabbing envs %v",err)
	}
	connPool, err := pgxpool.New(context.Background(),config.DriverURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	queries = NewStore(connPool)
	os.Exit(m.Run())
}
