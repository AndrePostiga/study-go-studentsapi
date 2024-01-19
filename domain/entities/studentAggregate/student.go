package studentaggregate

import (
	"strings"

	valueobjects "github.com/andrepostiga/api-go-gin/domain/valueObjects"
	"gorm.io/gorm"
)

type GetAge func() int

type Student struct {
	gorm.Model
	Name    string
	CPF     string `gorm:"unique"`
	Address *Address
	// Birthdate time.Time
	// Age       GetAge
}

// func (s *Student) GetAge() int {
// 	today := time.Now().UTC()
// 	age := today.Year() - s.Birthdate.Year()
// 	if today.UTC().Month() < s.Birthdate.UTC().Month() || (today.UTC().Month() == s.Birthdate.UTC().Month() && today.UTC().Day() < s.Birthdate.UTC().Day()) {
// 		age--
// 	}
// 	return age
// }

func NewStudent(
	name string,
	cpf string,
	cep *valueobjects.Cep,
	addressNumber uint,
	addressComplement string) (*Student, error) {

	student := &Student{}
	student.Name = name
	student.setCPF(cpf)
	student.setAddress(cep, addressNumber, addressComplement)

	return student, nil
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetCPF() string {
	return s.CPF
}

func (s *Student) setCPF(cpf string) {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	s.CPF = cpf
}

func (s *Student) GetAddress() *Address {
	return s.Address
}

func (s *Student) setAddress(cep *valueobjects.Cep, addressNumber uint, addressComplement string) {
	s.Address = NewAddress(cep, addressNumber, addressComplement)
}
