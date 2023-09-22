package pkg

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

func ValidateDocument(fl validator.FieldLevel) bool {
	document := fl.Field().String()
	if !OnlyNumber(document) {
		return false
	}

	if len(document) == 11 {
		return validateCPF(document)
	}

	if len(document) == 14 {
		return validateCNPJ(document)
	}

	return false
}

func validateCPF(cpf string) bool {
	firstDigit := calculateCPFDigit(cpf[:9])
	secondDigit := calculateCPFDigit(cpf[:10])

	return cpf[9] == firstDigit && cpf[10] == secondDigit
}

func validateCNPJ(cnpj string) bool {
	firstDigit := calculateCNPJDigits(cnpj[:12], 5)
	secondDigit := calculateCNPJDigits(cnpj[:13], 6)

	return cnpj[12] == firstDigit && cnpj[13] == secondDigit
}

func calculateCPFDigit(subString string) byte {
	sum := 0
	multiply := 10

	for _, char := range subString {
		sum += int(char-'0') * multiply
		multiply--
	}

	rest := sum % 11
	if rest < 2 {
		return '0'
	}
	return byte(11 - rest + '0')
}

func calculateCNPJDigits(subString string, weight int) byte {
	sum := 0
	for _, char := range subString {
		sum += int(char-'0') * weight
		weight--
		if weight == 1 {
			weight = 9
		}
	}
	rest := sum % 11
	if rest < 2 {
		return '0'
	}
	return byte(11 - rest + '0')
}

func OnlyNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
