package main

import (
	"log"
	"strconv"

	"github.com/StoneMarket/registration-service/config"
	"github.com/StoneMarket/registration-service/internal/api"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func main() {
	run()
}

func run() {
	env, err := config.ReadEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(env.PortDB)
	if err != nil {
		log.Fatal(err)
	}

	cfg := pgconn.Config{
		Host:     env.Host,
		Port:     uint16(port),
		Database: env.Database,
		User:     env.UserDB,
		Password: env.PasswordDB,
	}

	httpCfg := new(pgx.ConnConfig)
	httpCfg.Config = cfg

	if err := api.Run(env, httpCfg); err != nil {
		log.Fatal(err)
	}
}
