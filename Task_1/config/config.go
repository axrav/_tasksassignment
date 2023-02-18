package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() error {
	test := os.Getenv("PORT") // checking if the env variables are already loaded or not
	if test != "" {
		err := godotenv.Load() // loading the .env file
		if err != nil {
			return err
		}
	}
	return nil
}
