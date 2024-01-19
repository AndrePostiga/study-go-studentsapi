package request

import "strings"

type CreateStudentRequest struct {
	Name                string               `json:"name" binding:"required"`
	CPF                 string               `json:"cpf" valid:"cpf" binding:"required"`
	CreateAddresRequest CreateAddressRequest `json:"address" binding:"required"`
}

func IsValidCpf(cpf string) bool {
	var sum int
	var rest int

	// Cleaning the CPF string
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 || strings.Count(cpf, string(cpf[0])) == 11 {
		return false
	}

	for i := 1; i <= 9; i++ {
		sum += int(cpf[i-1]-'0') * (11 - i)
	}
	rest = sum * 10 % 11

	if rest == 10 || rest == 11 {
		rest = 0
	}
	if rest != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 1; i <= 10; i++ {
		sum += int(cpf[i-1]-'0') * (12 - i)
	}
	rest = sum * 10 % 11

	if rest == 10 || rest == 11 {
		rest = 0
	}
	if rest != int(cpf[10]-'0') {
		return false
	}

	return true
}
