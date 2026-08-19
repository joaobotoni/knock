package knock

import (
	"errors"
	"fmt"
)

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Database Database `yaml:"database"`
}

func (d *Database) Validate() error {
	var errs []error
	if d.Host == "" {
		errs = append(errs, fmt.Errorf("host é obrigatório"))
	}
	if d.Port == 0 {
		errs = append(errs, fmt.Errorf("port é obrigatório"))
	}
	if d.Name == "" {
		errs = append(errs, fmt.Errorf("name é obrigatório"))
	}
	if d.Username == "" {
		errs = append(errs, fmt.Errorf("username é obrigatório"))
	}
	if d.Password == "" {
		errs = append(errs, fmt.Errorf("password é obrigatório"))
	}
	return errors.Join(errs...)
}

func (c *Config) Validate() error {
	if err := c.Database.Validate(); err != nil {
		return fmt.Errorf("database: %w", err)
	}
	return nil
}

