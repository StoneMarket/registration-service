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
	Host       string `envconfig:"REST_HOST" default:"0.0.0.0"`
	Port       string `envconfig:"REST_PORT" default:"8080"`
	PrivateKey string `envconfig:"PRIVATE_KEY" required:"true"`
	HostDB     string `envconfig:"HOST_DB" required:"true"`
	PortDB     string `envconfig:"PORT_DB" required:"true"`
	Database   string `envconfig:"DATABASE" required:"true"`
	UserDB     string `envconfig:"USER_DB" required:"true"`
	PasswordDB string `envconfig:"PASSWORD_DB" required:"true"`
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
		return nil, fmt.Errorf("config: failed to decode public key: %s", err)
	}

	return jwt.ParseRSAPrivateKeyFromPEM(data)
}
