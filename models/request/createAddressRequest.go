package request

import "regexp"

type CreateAddressRequest struct {
	Number     uint   `json:"number" binding:"required"`
	CEP        string `json:"cep" valid:"cep" binding:"required"`
	Complement string `json:"complement"`
}

func IsValidCep(cep string) bool {
	// Regex pattern for CEP (12345-678 or 12345678)
	pattern := `^\d{5}-?\d{3}$`
	matched, _ := regexp.MatchString(pattern, cep)
	return matched
}
