package env

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Load(path string) error {
	if err := godotenv.Load(path); err != nil {
		return fmt.Errorf("carregar %q: %w", path, err)
	}
	return nil
}
