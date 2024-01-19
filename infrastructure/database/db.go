package database

import (
	"log"

	s "github.com/andrepostiga/api-go-gin/domain/entities/studentAggregate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectPostgres() *gorm.DB {
	connString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connString))
	if err != nil {
		log.Panic("Error while connecting to database")
	}

	DB.AutoMigrate(&s.Student{})
	DB.AutoMigrate(&s.Address{})
	return DB
}
