package validators

import "errors"

func ValidateSpecialty(specialty string) error{
	if specialty == "" {
		return errors.New("Especialidade não pode ser vazia")
	}

	if len(specialty) < 3 || len(specialty) > 20 {
		return errors.New("Especialidade inválida (entre 3 a 20 caracteres)")
	}

	return nil
}