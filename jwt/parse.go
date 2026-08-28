package jwt

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var methods = []string{"RS256"}

func keyfunc(key *rsa.PublicKey) jwt.Keyfunc {
	return func(token *jwt.Token) (any, error) { return key, nil }
}

func parse(key *rsa.PublicKey, raw string, c *claims) error {
	if _, err := jwt.ParseWithClaims(raw, c, keyfunc(key), jwt.WithValidMethods(methods)); err != nil {
		return fmt.Errorf("parseando token: %w", err)
	}
	return nil
}

func decodeClaims(key *rsa.PublicKey, raw string) (*claims, error) {
	var c claims
	if err := parse(key, raw, &c); err != nil {
		return nil, fmt.Errorf("parseando claims: %w", err)
	}
	return &c, nil
}
