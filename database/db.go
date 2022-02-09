package database

import (
	"log"

	"github.com/jadson-medeiros/gin-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Error with DB CONNECTION")
	}

	DB.AutoMigrate(&models.Student{})
}
