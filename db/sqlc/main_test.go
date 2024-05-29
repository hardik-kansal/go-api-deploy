package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var queries *Queries

const (
	driverUrl  = "postgresql://root:123@localhost:5432/simplebank?sslmode=disable"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, driverUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)
	queries = New(conn)
	os.Exit(m.Run())
}
