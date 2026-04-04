package validators

import (
	"errors"
)

func ValidateAge(age int, role string) error {
	if age < 21 {
		return errors.New("Você precisa ser maior de 21 anos para exercer a função de terapeuta")
	}
	return nil
}