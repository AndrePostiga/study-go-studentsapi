package domainservices

import (
	s "github.com/andrepostiga/api-go-gin/domain/entities/studentAggregate"
	c "github.com/andrepostiga/api-go-gin/domain/externalServicesContracts"
	r "github.com/andrepostiga/api-go-gin/domain/repository"
	"github.com/andrepostiga/api-go-gin/models/request"
)

type ICreateStudentService interface {
	Create(request *request.CreateStudentRequest) (*s.Student, error)
}

type CreateStudentService struct {
	repository r.IStudentRepository
	cepClient  c.ICepClient
}

func NewCreateStudentService(repository r.IStudentRepository, cepClient c.ICepClient) ICreateStudentService {
	return &CreateStudentService{
		repository: repository,
		cepClient:  cepClient,
	}
}

// Create implements CreateStudentService.
func (c *CreateStudentService) Create(request *request.CreateStudentRequest) (*s.Student, error) {
	cep, err := c.cepClient.Get(request.CreateAddresRequest.CEP)
	if err != nil {
		return nil, err
	}

	student, err := s.NewStudent(request.Name, request.CPF, cep, request.CreateAddresRequest.Number, request.CreateAddresRequest.Complement)
	if err != nil {
		return nil, err
	}

	if err := c.repository.Save(student); err != nil {
		return nil, err
	}

	return student, nil
}
