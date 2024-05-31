package main

import (
	"context"
	"log"

	"github.com/hardik-kansal/go-api-deploy/api"
	db "github.com/hardik-kansal/go-api-deploy/db/sqlc"
	"github.com/hardik-kansal/go-api-deploy/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	ctx:=context.Background()
	connPool, err := pgxpool.New(ctx, config.DriverURL)
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
/* 	for i:=0;i<20;i++{
		arg := db.CreateAccountParams{
			Owner:    util.RandomOwner(),
			Balance:  util.RandomMoney(),
			Currency: util.RandomCurrency(),
		}
		store.CreateAccount(context.Background(),arg)
	} */
}
