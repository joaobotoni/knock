package jwt

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Issue(key *rsa.PrivateKey, c Customer, now time.Time, ttl time.Duration) ([]byte, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, newClaims(c, now, ttl))
	s, err := t.SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("assinando token: %w", err)
	}
	return []byte(s), nil
}
