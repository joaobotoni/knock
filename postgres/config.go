package postgres

import (
	"errors"
	"fmt"
)

const (
	portMin = 1
	portMax = 65535
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func validate(c *Config) error {
	var errs []error
	if c.Host == "" {
		errs = append(errs, fmt.Errorf("host é obrigatório"))
	}
	switch {
	case c.Port == 0:
		errs = append(errs, fmt.Errorf("port é obrigatório"))
	case c.Port < portMin:
		errs = append(errs, fmt.Errorf("port não pode ser negativo: %d", c.Port))
	case c.Port > portMax:
		errs = append(errs, fmt.Errorf("port acima do máximo (%d): %d", portMax, c.Port))
	}
	if c.Name == "" {
		errs = append(errs, fmt.Errorf("name é obrigatório"))
	}
	if c.Username == "" {
		errs = append(errs, fmt.Errorf("username é obrigatório"))
	}
	if c.Password == "" {
		errs = append(errs, fmt.Errorf("password é obrigatório"))
	}
	return errors.Join(errs...)
}
