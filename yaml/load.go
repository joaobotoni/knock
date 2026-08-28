package yaml

import (
	"fmt"
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v4"
)

func Load[T any](path string) (*T, error) {
	var inc includes
	var dir string = filepath.Dir(path)
	if err := read(path, &inc); err != nil {
		return nil, err
	}

	data, err := merge(os.DirFS(dir), inc.Files)
	if err != nil {
		return nil, err
	}

	var out T
	if err := yaml.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("parsear configuração combinada de %q: %w", path, err)
	}
	return &out, nil
}
