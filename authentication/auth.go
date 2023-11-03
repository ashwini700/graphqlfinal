package authentication

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Key int

const Ctxkey Key = 1

type Auth struct {
	privatekey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*Auth, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("private key or public key cannot be nil")
	}
	return &Auth{
		privatekey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func (a *Auth) GenerateToken(claims jwt.RegisteredClaims) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tknstr, err := tkn.SignedString(a.privatekey)
	if err != nil {
		return "", fmt.Errorf("signing token %w", err)
	}
	return tknstr, nil
}

func (a *Auth) ValidateToken(token string) (jwt.RegisteredClaims, error) {
	var c jwt.RegisteredClaims
	tkn, err := jwt.ParseWithClaims(token, &c, func(t *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})
	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("parsing token %w", err)
	}
	if !tkn.Valid {
		return jwt.RegisteredClaims{}, fmt.Errorf("not a valid token %w", err)
	}
	return c, nil

}
