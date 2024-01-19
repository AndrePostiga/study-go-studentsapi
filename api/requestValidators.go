package api

import (
	"github.com/andrepostiga/api-go-gin/models/request"
	"github.com/asaskevich/govalidator"
)

func registryRequestValidators() {
	govalidator.TagMap["cpf"] = govalidator.Validator(func(str string) bool {
		return request.IsValidCpf(str)
	})

	govalidator.TagMap["cep"] = govalidator.Validator(func(str string) bool {
		return request.IsValidCep(str)
	})
}
