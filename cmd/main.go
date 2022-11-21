package main

import (
	"log"

	"github.com/StoneMarket/registration-service/config"
	"github.com/StoneMarket/registration-service/internal/api"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	run()
}

func run() {
	env, err := config.ReadEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(env)

	httpCfg, err := pgx.ParseConfig(env.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := api.Run(env, httpCfg); err != nil {
		log.Fatal(err)
	}
}
