package argon2

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func key(passwd string, salt []byte, p Params) []byte {
	return argon2.IDKey([]byte(passwd), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)
}

func salt(size uint32) ([]byte, error) {
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("erro ao gerar o salt: %w", err)
	}
	return b, nil
}
