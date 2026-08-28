package argon2

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

func encode(p Params, salt, hash []byte) string {
	return fmt.Sprintf("$%s$%s", encodeParams(p), encodeHash(salt, hash))
}

func encodeParams(p Params) string {
	return fmt.Sprintf("argon2id$v=%d$m=%d,t=%d,p=%d", argon2.Version, p.Memory, p.Iterations, p.Parallelism)
}

func encodeHash(salt, hash []byte) string {
	return fmt.Sprintf("%s$%s", encodeB64(salt), encodeB64(hash))
}

func encodeB64(b []byte) string {
	return base64.RawStdEncoding.EncodeToString(b)
}
