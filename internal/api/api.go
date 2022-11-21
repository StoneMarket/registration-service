package api

import (
	"context"
	"fmt"

	"github.com/StoneMarket/registration-service/config"
	v1 "github.com/StoneMarket/registration-service/internal/controller/http/v1"
	"github.com/StoneMarket/registration-service/internal/storage"
	"github.com/jackc/pgx/v4"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

const path = ":8080"

func Run(env *config.Environment, cfg *pgx.ConnConfig) error {
	conn, err := storage.Connect(context.Background(), cfg)
	if err != nil {
		return err
	}

	pk, err := env.MakeRSAPrivateKey()
	if err != nil {
		return fmt.Errorf("rsa private key, error :%w", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("logger creation, error :%w", err)
	}

	controller := v1.NewContoller(storage.NewStorage(conn), pk, logger)
	rest := v1.NewV1(controller)

	if err := fasthttp.ListenAndServe(path, rest.RestV1); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
