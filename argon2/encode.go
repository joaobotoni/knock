package argon2

import (
	"encoding/base64"
	"fmt"

	"github.com/joaobotoni/knock"
	"golang.org/x/crypto/argon2"
)


func encode(p knock.Params, salt, hash []byte) string {
	b64 := base64.RawStdEncoding.EncodeToString
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64(salt), b64(hash))
}

