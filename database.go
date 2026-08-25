package knock

import (
	"errors"
	"fmt"
)

const (
	portMin = 1
	portMax = 65535
)

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func ValidateDatabase(d *Database) error {
	var errs []error
	if d.Host == "" {
		errs = append(errs, fmt.Errorf("host é obrigatório"))
	}
	switch {
	case d.Port == 0:
		errs = append(errs, fmt.Errorf("port é obrigatório"))
	case d.Port < portMin:
		errs = append(errs, fmt.Errorf("port não pode ser negativo: %d", d.Port))
	case d.Port > portMax:
		errs = append(errs, fmt.Errorf("port acima do máximo (%d): %d", portMax, d.Port))
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
