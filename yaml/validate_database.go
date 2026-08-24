package yaml

import (
	"errors"
	"fmt"
)

const (
	portaMin = 1
	portaMax = 65535
)

func ValidateDatabase(d *Database) error {
	var errs []error
	if d.Host == "" {
		errs = append(errs, fmt.Errorf("host é obrigatório"))
	}
	switch {
	case d.Port == 0:
		errs = append(errs, fmt.Errorf("port é obrigatório"))
	case d.Port < portaMin:
		errs = append(errs, fmt.Errorf("port não pode ser negativo: %d", d.Port))
	case d.Port > portaMax:
		errs = append(errs, fmt.Errorf("port acima do máximo (%d): %d", portaMax, d.Port))
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
