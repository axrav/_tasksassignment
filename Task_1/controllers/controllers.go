package controllers

import "github.com/gofiber/fiber/v2"

func CreateBook(c *fiber.Ctx) error {
	return c.SendString("Create Book")

}

func CreateMovie(c *fiber.Ctx) error {
	return c.SendString("Create Movie")

}

func CreateMusic(c *fiber.Ctx) error {
	return c.SendString("Create Music")

}

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Get Books")

}

func GetMovies(c *fiber.Ctx) error {
	return c.SendString("Get Movies")

}

func GetAllMusic(c *fiber.Ctx) error {
	return c.SendString("Get All Music")

}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Get Book")

}

func GetMovie(c *fiber.Ctx) error {
	return c.SendString("Get Movie")
}

func GetMusic(c *fiber.Ctx) error {
	return c.SendString("Get Music")

}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete Book")

}

func DeleteMovie(c *fiber.Ctx) error {
	return c.SendString("Delete Movie")

}

func DeleteMusic(c *fiber.Ctx) error {
	return c.SendString("Delete Music")

}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("Update Book")

}

func UpdateMovie(c *fiber.Ctx) error {
	return c.SendString("Update Movie")

}

func UpdateMusic(c *fiber.Ctx) error {
	return c.SendString("Update Music")

}
