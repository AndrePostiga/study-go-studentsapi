package repository

import studentaggregate "github.com/andrepostiga/api-go-gin/domain/entities/studentAggregate"

type IStudentRepository interface {
	Get(id string) (*studentaggregate.Student, error)
	GetStudents() ([]studentaggregate.Student, error)
	Save(student *studentaggregate.Student) error
}
