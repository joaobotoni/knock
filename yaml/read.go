package yaml

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

func read(path string, out any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("ler %q: %w", path, err)
	}
	if err := yaml.Unmarshal(data, out); err != nil {
		return fmt.Errorf("parsear %q: %w", path, err)
	}
	return nil
}
