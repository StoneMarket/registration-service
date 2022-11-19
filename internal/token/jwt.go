package token

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/StoneMarket/registration-service/internal/models"
	"github.com/golang-jwt/jwt"
)

const defaultExpTime time.Duration = 98 * time.Hour

type Token string

func GetPrivateKey(privateKeyRaw string) (*rsa.PrivateKey, error) {
	pk, err := base64.StdEncoding.DecodeString(privateKeyRaw)
	if err != nil {
		return nil, fmt.Errorf("private key decode error: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(pk)
	if err != nil {
		return nil, fmt.Errorf("private key parse error: %w", err)
	}

	return key, nil
}

func GenerateNewToken(privateKey *rsa.PrivateKey, userData *models.User) (Token, error) {
	claims := jwt.MapClaims{
		"image": userData.Image,
		"name":  userData.Name,
		"login": userData.Login,
		"exp":   time.Now().Add(defaultExpTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tkn, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return Token(tkn), nil
}
