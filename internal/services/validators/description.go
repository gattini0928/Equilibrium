package validators

import "errors"

func ValidateDescription(description string) error{
	if description == "" {
		return errors.New("Descrição não pode ser vazia")
	}

	if len(description) < 50 || len(description) > 200 {
		return errors.New("Descrição inválida (entre 50 a 200 caracteres)")
	}

	return nil
}