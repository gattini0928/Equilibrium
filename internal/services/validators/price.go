package validators

import "errors"

func ValidatePrice(price float64) error{
	if price < 20 {
		return errors.New("Valor da consulta não pode ser menor que 20 reais")
	}

	if price > 300 {
		return errors.New("Valor da consulta muito alta")
	}

	return nil
}