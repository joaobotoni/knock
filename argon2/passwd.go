package argon2

import (
	"crypto/subtle"
	"fmt"

	"github.com/joaobotoni/knock"
)

var DefaultParams = knock.Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func GeneratePasswdHash(passwd string) (string, error) {
	return GeneratePasswdHashWithParams(passwd, DefaultParams)
}

func GeneratePasswdHashWithParams(passwd string, p knock.Params) (string, error) {
	salt, err := salt(p.SaltLength)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar a senha: %w", err)
	}
	return encode(p, salt, key(passwd, salt, p)), nil
}

func ComparePasswdHash(passwd, encoded string) (bool, error) {
	p, salt, hash, err := decode(encoded)
	if err != nil {
		return false, err
	}
	return subtle.ConstantTimeCompare(hash, key(passwd, salt, p)) == 1, nil
}
