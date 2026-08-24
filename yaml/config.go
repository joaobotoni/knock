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

	data, err = expand(data)
	if err != nil {
		return nil, fmt.Errorf("expandir variáveis em %q: %w", path, err)
	}

	var cfg knock.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsear YAML de %q: %w", path, err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuração inválida em %q: %w", path, err)
	}
	return &cfg, nil
}

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
