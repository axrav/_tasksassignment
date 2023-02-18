package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Pgres is the global variable which is used to access the database
var Pgres *gorm.DB

// Connect is used to connect to the database, Gorm is used as the ORM
func Connect() error {
	var DB_URL = os.Getenv("DB_URL")
	var err error
	Pgres, err = gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		return err
	}
	// migrate the schema
	err = Pgres.AutoMigrate(&Book{}, &Movie{}, &Music{})
	if err != nil {
		return err
	}
	return nil
}
