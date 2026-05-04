package validators

import (
	"errors"
	"unicode"
)

func ValidateCrm(crm string) error {
    if crm == "" {
        return errors.New("CRM não pode ser vazio")
    }

    if len(crm) != 9 {
        return errors.New("CRM precisa ter 9 caracteres (ex: 123456-SP)")
    }

    for _, c := range crm[0:6] {
        if !unicode.IsDigit(c) {
            return errors.New("Os primeiros 6 caracteres do CRM precisam ser dígitos")
        }
    }

    if crm[6] != '-' {
        return errors.New("CRM inválido, separador esperado: -")
    }

    var validUFs = map[string]bool{
        "AC": true, "AL": true, "AP": true, "AM": true, "BA": true,
        "CE": true, "DF": true, "ES": true, "GO": true, "MA": true,
        "MT": true, "MS": true, "MG": true, "PA": true, "PB": true,
        "PR": true, "PE": true, "PI": true, "RJ": true, "RN": true,
        "RS": true, "RO": true, "RR": true, "SC": true, "SP": true,
        "SE": true, "TO": true,
    }

    uf := crm[7:9]
    if !validUFs[uf] {
        return errors.New("Sigla de estado inválida")
    }

    return nil
}