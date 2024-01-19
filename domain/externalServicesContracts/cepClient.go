package externalservicescontracts

import valueobjects "github.com/andrepostiga/api-go-gin/domain/valueObjects"

type ICepClient interface {
	Get(cep string) (*valueobjects.Cep, error)
}
