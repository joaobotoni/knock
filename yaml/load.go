package yaml

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joaobotoni/knock"
	"go.yaml.in/yaml/v4"
)

func Load(path string) (*knock.Manifest, error) {
	var inc includes
	var dir string = filepath.Dir(path)
	if err := read(path, &inc); err != nil {
		return nil, err
	}

	data, err := merge(os.DirFS(dir), inc.Files)
	if err != nil {
		return nil, err
	}

	var manifest knock.Manifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("parsear configuração combinada de %q: %w", path, err)
	}

	if err := knock.ValidateManifest(&manifest); err != nil {
		return nil, fmt.Errorf("configuração inválida em %q: %w", path, err)
	}
	return &manifest, nil
}
