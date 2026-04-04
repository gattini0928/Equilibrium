package validators

import (
	"errors"
	"strconv"
	"strings"
)

func ValidateCpf(cpf string) error {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	cpf = strings.ReplaceAll(cpf, " ", "")

	if len(cpf) != 11 {
		return errors.New("Seu CPF precisa ter 11 dígitos")
	}
	_, err := strconv.Atoi(cpf) 
	if err != nil {
		return errors.New("CPF deve conter apenas números")
	}

	if cpf == "00000000000" || cpf == "11111111111" || 
       cpf == "22222222222" || cpf == "33333333333" ||
       cpf == "44444444444" || cpf == "55555555555" ||
       cpf == "66666666666" || cpf == "77777777777" ||
       cpf == "88888888888" || cpf == "99999999999" {
        return errors.New("CPF inválido")
    }

	return nil
}