package yaml

import (
	"fmt"
	"github.com/joaobotoni/knock"
	"go.yaml.in/yaml/v4"
	"os"
)

func Load(path string) (*knock.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ler configuração %q: %w", path, err)
	}

	var cfg *knock.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsear YAML de %q: %w", path, err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuração inválida em %q: %w", path, err)
	}

	return cfg, nil
}
