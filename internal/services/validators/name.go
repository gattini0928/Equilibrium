package validators

import (
	"strings"
	"unicode"
	"errors"
)

func ValidateName(name string) error {

	if name == "" {
		return errors.New("Seu nome não pode estar vazio")
	}

	
	for _, char := range name {
		if !unicode.IsLetter(char) && char != ' ' {
			return errors.New("Seu nome precisa conter apenas letras")
		}
	}

	parts := strings.Fields(name)
	if len(parts) < 2 {
		return errors.New("Digite seu nome e sobrenome")
	}

	if len(parts[0]) < 2 || len(parts[1]) < 2 {
		return errors.New("nome ou sobrenome incompletos")
	}

	if parts[0] == parts[1] {
		return errors.New("nome e sobrenome iguais")
	}

	return nil
}