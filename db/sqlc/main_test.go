package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var queries *SQLStore

const (
	driverUrl  = "postgresql://root:123@localhost:5432/simplebank?sslmode=disable"
)

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(),driverUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	queries = NewStore(connPool)
	os.Exit(m.Run())
}
