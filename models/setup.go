package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=127.0.0.1 user=postgres password=mysecretpassword dbname=postgres port=8000 sslmode=disable timezone=Europe/London"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // change the database provider if necessary

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&System{})

	DB = database
}
