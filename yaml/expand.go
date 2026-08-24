package yaml

import (
	"fmt"
	"os"
)

func expand(data []byte) ([]byte, error) {
	var err error
	out := os.Expand(string(data), func(key string) string {
		v, ok := os.LookupEnv(key)
		if !ok && err == nil {
			err = fmt.Errorf("variável não definida: %s", key)
		}
		return v
	})
	return []byte(out), err
}
