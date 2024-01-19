package studentaggregate

import (
	valueobjects "github.com/andrepostiga/api-go-gin/domain/valueObjects"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	valueobjects.Cep
	Number     uint
	Complement string
	Country    string `gorm:"default:BR"`
	StudentId  uint
}

func NewAddress(cep *valueobjects.Cep, number uint, complement string) *Address {
	address := &Address{}
	address.Cep = *cep
	address.Number = uint(number)
	address.Complement = complement
	return address
}
