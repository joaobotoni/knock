package yaml

import (
	"fmt"
	"io/fs"
)

func merge(fsys fs.FS, refs []map[string]string) ([]byte, error) {
	var out []byte
	for _, ref := range refs {
		for name, file := range ref {
			data, err := fs.ReadFile(fsys, file)
			if err != nil {
				return nil, fmt.Errorf("seção %q: %w", name, err)
			}

			data, err = expand(data)
			if err != nil {
				return nil, fmt.Errorf("seção %q: %w", name, err)
			}

			out = append(out, data...)
			out = append(out, '\n')
		}
	}
	return out, nil
}
