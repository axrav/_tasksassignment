package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_URL = os.Getenv("DB_URL")
var Pgres *gorm.DB

func Connect() error {
	Pgres, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
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
