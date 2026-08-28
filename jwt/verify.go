package jwt

import (
	"crypto/rsa"
	"fmt"
)

func Verify(key *rsa.PublicKey, raw string) (*Customer, error) {
	c, err := decodeClaims(key, raw)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar o token: %w", err)
	}
	return &c.Customer, nil
}
