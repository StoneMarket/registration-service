package v1

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/StoneMarket/registration-service/internal/models"
	"github.com/StoneMarket/registration-service/internal/token"
	"go.uber.org/zap"
)

type Controller struct {
	storage Storage
	logger  *zap.Logger
	pk      *rsa.PrivateKey
}

type Storage interface {
	SetUserData(ctx context.Context, user *models.User) error
	FindUserData(ctx context.Context, login, password string) (*models.User, error)
}

func NewContoller(stg Storage, pk *rsa.PrivateKey, logger *zap.Logger) *Controller {
	return &Controller{
		storage: stg,
		logger:  logger,
		pk:      pk,
	}
}

func (c *Controller) RegisterNewUser(ctx context.Context, user *models.User) (token.Token, error) {
	if err := c.storage.SetUserData(ctx, user); err != nil {
		c.logger.Error("registration new user", zap.Error(err))
		return "", fmt.Errorf("registration new user, error: %w", err)
	}

	token, err := token.GenerateNewToken(c.pk, user)
	if err != nil {
		c.logger.Error("registration new user", zap.Error(err))
		return "", fmt.Errorf("registration new user, error: %w", err)
	}

	return token, nil
}

func (c *Controller) Login(ctx context.Context, login, password string) (token.Token, error) {
	user, err := c.storage.FindUserData(ctx, login, password)
	if err != nil {
		c.logger.Error("login user", zap.Error(err))
		return "", fmt.Errorf("login user, error: %w", err)
	}

	token, err := token.GenerateNewToken(c.pk, user)
	if err != nil {
		c.logger.Error("login user", zap.Error(err))
		return "", fmt.Errorf("login user token, error: %w", err)
	}

	return token, nil
}
