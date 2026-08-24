package yaml

import (
	"fmt"
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v4"
)

func Load(path string) (*Manifest, error) {
	var inc includes
	if err := read(path, &inc); err != nil {
		return nil, err
	}

	data, err := merge(os.DirFS(filepath.Dir(path)), inc.Files)
	if err != nil {
		return nil, err
	}

	var manifest Manifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("parsear configuração combinada de %q: %w", path, err)
	}

	if err := ValidateManifest(&manifest); err != nil {
		return nil, fmt.Errorf("configuração inválida em %q: %w", path, err)
	}
	return &manifest, nil
}
