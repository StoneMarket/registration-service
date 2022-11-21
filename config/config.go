package config

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"net"

	"github.com/golang-jwt/jwt"
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Host        string `envconfig:"REST_HOST" default:"0.0.0.0"`
	Port        string `envconfig:"REST_PORT" default:"8080"`
	PrivateKey  string `envconfig:"PRIVATE_KEY" required:"true"`
	PostgresDSN string `envconfig:"POSTGRES_DSN" required:"true"`
}

func ReadEnvironment() (*Environment, error) {
	var env Environment
	if err := envconfig.Process("", &env); err != nil {
		return nil, fmt.Errorf("config: failed to get registration config: %w", err)
	}

	return &env, nil
}

func (env *Environment) Addr() string {
	return net.JoinHostPort(env.Host, env.Port)
}

func (env *Environment) MakeRSAPrivateKey() (*rsa.PrivateKey, error) {
	data, err := base64.StdEncoding.DecodeString(env.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("config: failed to decode private key: %s", err)
	}

	return jwt.ParseRSAPrivateKeyFromPEM(data)
}
