package config

import "github.com/joho/godotenv"

func Load() error {
	err := godotenv.Load() // loading the .env file
	if err != nil {
		return err
	}
	return nil
}
