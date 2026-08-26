package argon2

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/joaobotoni/knock"
	"golang.org/x/crypto/argon2"
)

var ErrInvalidHash = "hash inválido"

func decode(encoded string) (knock.Params, []byte, []byte, error) {
	var p knock.Params

	parts := strings.Split(encoded, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return p, nil, nil, fmt.Errorf("%s", ErrInvalidHash)
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return p, nil, nil, fmt.Errorf("%s: %w", ErrInvalidHash, err)
	}
	if version != argon2.Version {
		return p, nil, nil, fmt.Errorf("%s: versão %d incompatível", ErrInvalidHash, version)
	}

	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism); err != nil {
		return p, nil, nil, fmt.Errorf("%s: %w", ErrInvalidHash, err)
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(parts[4])
	if err != nil {
		return p, nil, nil, fmt.Errorf("%s: %w", ErrInvalidHash, err)
	}

	hash, err := base64.RawStdEncoding.Strict().DecodeString(parts[5])
	if err != nil {
		return p, nil, nil, fmt.Errorf("%s: %w", ErrInvalidHash, err)
	}

	p.SaltLength = uint32(len(salt))
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}
