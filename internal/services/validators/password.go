package validators

import (
	"fmt"
	"unicode"
	"errors"
)

func ValidatePassword(password string) error{
	if password == "" {
		return  errors.New("A senha não pode ser vazia")
	}

	if len(password) < 8 {
		return errors.New("Sua senha deve conter no mínimo 8 caracteres")
	}

	
	specials := func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	}

	rules := map[string]func(rune) bool {
		"letra maiúscula": unicode.IsUpper,
		"letra minúscula": unicode.IsLower,
		"dígito": unicode.IsDigit,
		"caractere especial": specials,
	}

	for ruleName, ruleFunc := range rules {
		found := false
		for _, char := range password {
			if ruleFunc(char) {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("sua senha deve conter pelo menos um(a) %s", ruleName)
		}
	}

	return nil
}