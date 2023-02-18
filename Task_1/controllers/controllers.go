package controllers

import (
	"sync"

	"github.com/axrav/Task1/db"
	"github.com/gofiber/fiber/v2"
)

// creating a mutex and waitgroup to make sure that only one goroutine is accessing the database at a time, avoiding race conditions
var mu sync.Mutex
var wg sync.WaitGroup

// create handlers for books, movies and music
func CreateBook(c *fiber.Ctx) error {
	book := new(db.Book)
	err := c.BodyParser(book)
	if err != nil {
		return err
	}

	wg.Add(1) // incrementing the waitgroup counter by 1 to make sure that the goroutine is done before the next one starts

	b, err := db.CreateBookinDb(*book, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(b)

}

func CreateMovie(c *fiber.Ctx) error {
	movie := new(db.Movie)
	err := c.BodyParser(movie)
	if err != nil {
		return err
	}
	wg.Add(1)
	m, err := db.CreateMovieinDb(*movie, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(m)

}

func CreateMusic(c *fiber.Ctx) error {
	music := new(db.Music)
	err := c.BodyParser(music)
	if err != nil {
		return err
	}
	wg.Add(1)
	m, err := db.CreateMusicinDb(*music, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(m)

}

// get all handlers for books, movies and music

func GetBooks(c *fiber.Ctx) error {
	wg.Add(1)
	books, err := db.GetBooksfromDb(&mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(books)

}

func GetMovies(c *fiber.Ctx) error {
	wg.Add(1)
	movies, err := db.GetMoviesfromDb(&mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(movies)

}

func GetAllMusic(c *fiber.Ctx) error {
	wg.Add(1)
	musics, err := db.GetAllMusicfromDb(&mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(musics)

}

// get handlers for each model
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	wg.Add(1)
	book, err := db.GetBookfromDb(id, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(book)

}

func GetMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	wg.Add(1)
	movie, err := db.GetMoviefromDb(id, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(movie)
}

func GetMusic(c *fiber.Ctx) error {
	id := c.Params("id")
	wg.Add(1)
	music, err := db.GetMusicfromDb(id, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(music)

}

// delete handlers for each model

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	wg.Add(1)
	err := db.DeleteBookfromDb(id, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendString("Deleted Book")

}

func DeleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	wg.Add(1)
	err := db.DeleteMoviefromDb(id, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendString("Deleted Movie")

}

func DeleteMusic(c *fiber.Ctx) error {
	id := c.Params("id")
	wg.Add(1)
	err := db.DeleteMusicfromDb(id, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendString("Deleted Music")

}

// updating handlers for each model

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(db.Book)
	err := c.BodyParser(book)
	if err != nil {
		return err
	}
	wg.Add(1)
	b, err := db.UpdateBookinDb(id, *book, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(b)

}

func UpdateMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	movie := new(db.Movie)
	err := c.BodyParser(movie)
	if err != nil {
		return err
	}
	wg.Add(1)
	m, err := db.UpdateMovieinDb(id, *movie, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(m)
}

func UpdateMusic(c *fiber.Ctx) error {
	id := c.Params("id")
	music := new(db.Music)
	err := c.BodyParser(music)
	if err != nil {
		return err
	}
	wg.Add(1)
	m, err := db.UpdateMusicinDb(id, *music, &mu, &wg)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(m)

}
