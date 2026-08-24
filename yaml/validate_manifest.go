package yaml

import "fmt"

func ValidateManifest(m *Manifest) error {
	if err := ValidateDatabase(&m.Database); err != nil {
		return fmt.Errorf("database: %w", err)
	}
	return nil
}
