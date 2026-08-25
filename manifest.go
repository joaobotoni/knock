package knock

import "fmt"

type Manifest struct {
	Database Database `yaml:"database"`
}

func ValidateManifest(m *Manifest) error {
	if err := ValidateDatabase(&m.Database); err != nil {
		return fmt.Errorf("database: %w", err)
	}
	return nil
}
