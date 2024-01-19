package repositories

import (
	s "github.com/andrepostiga/api-go-gin/domain/entities/studentAggregate"
	"github.com/andrepostiga/api-go-gin/domain/repository"
	"gorm.io/gorm"
)

type StudentRepositoryConcrete struct {
	Conn *gorm.DB
}

func NewStudentRepository(conn *gorm.DB) repository.IStudentRepository {
	return &StudentRepositoryConcrete{Conn: conn}
}

// Get implements repository.StudentRepository.
func (repository *StudentRepositoryConcrete) Get(id string) (*s.Student, error) {
	var response *s.Student
	if err := repository.Conn.Preload("Address").First(&response, id).Error; err != nil {
		return nil, err
	}

	return response, nil
}

// GetStudents implements repository.StudentRepository.
func (repository *StudentRepositoryConcrete) GetStudents() ([]s.Student, error) {
	var response []s.Student
	if err := repository.Conn.Preload("Address").Find(&response).Error; err != nil {
		return nil, err
	}

	return response, nil
}

// Save implements repository.StudentRepository.
func (repository *StudentRepositoryConcrete) Save(student *s.Student) error {
	if err := repository.Conn.Save(&student).Error; err != nil {
		return err
	}

	return nil
}
