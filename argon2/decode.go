package argon2

import (
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

var ErrInvalidHash = errors.New("hash inválido")

type decoded struct {
	Params Params
	Salt   []byte
	Hash   []byte
}

func decode(encoded string) (decoded, error) {
	var d decoded

	parts := strings.Split(encoded, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return d, fmt.Errorf("%w: formato inesperado", ErrInvalidHash)
	}

	if err := decodeVersion(parts[2]); err != nil {
		return d, err
	}

	p, err := decodeParams(parts[3])
	if err != nil {
		return d, err
	}

	salt, err := decodeB64(parts[4])
	if err != nil {
		return d, fmt.Errorf("salt: %w", err)
	}

	hash, err := decodeB64(parts[5])
	if err != nil {
		return d, fmt.Errorf("hash: %w", err)
	}

	p.SaltLength = uint32(len(salt))
	p.KeyLength = uint32(len(hash))

	return decoded{Params: p, Salt: salt, Hash: hash}, nil
}

func decodeVersion(s string) error {
	var v int
	if _, err := fmt.Sscanf(s, "v=%d", &v); err != nil {
		return fmt.Errorf("%w: versão ilegível: %v", ErrInvalidHash, err)
	}
	if v != argon2.Version {
		return fmt.Errorf("%w: versão %d incompatível", ErrInvalidHash, v)
	}
	return nil
}

func decodeParams(s string) (Params, error) {
	var p Params
	if _, err := fmt.Sscanf(s, "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism); err != nil {
		return p, fmt.Errorf("%w: parâmetros ilegíveis: %v", ErrInvalidHash, err)
	}
	return p, nil
}

func decodeB64(s string) ([]byte, error) {
	b, err := base64.RawStdEncoding.Strict().DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("%w: base64 inválido: %v", ErrInvalidHash, err)
	}
	return b, nil
}
