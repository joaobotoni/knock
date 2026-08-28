package jwt

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func LoadPrivateKey(pem []byte) (*rsa.PrivateKey, error) {
	k, err := jwt.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		return nil, fmt.Errorf("parseando chave privada: %w", err)
	}
	return k, nil
}

func LoadPublicKey(pem []byte) (*rsa.PublicKey, error) {
	k, err := jwt.ParseRSAPublicKeyFromPEM(pem)
	if err != nil {
		return nil, fmt.Errorf("parseando chave pública: %w", err)
	}
	return k, nil
}
