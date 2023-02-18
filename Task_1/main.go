package main

import (
	"os"

	"github.com/axrav/Task1/config"
	"github.com/axrav/Task1/db"
	"github.com/axrav/Task1/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// loading the .env file to get the DB_URL, PORT
	err := config.Load()
	if err != nil {
		panic(err)
	}
	// connecting to the database
	err = db.Connect()
	if err != nil {
		panic(err)
	}
	// starting the server with gofiber
	app := fiber.New()
	// setting up the routes
	router.Setup(app)
	// listening to the port
	port := os.Getenv("PORT") // getting the port from .env file
	app.Listen(":" + port)    // listening to the port

}
