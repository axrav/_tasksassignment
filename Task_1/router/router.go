package router

import (
	"github.com/axrav/Task1/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// checking if the server is alive or not
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Online")
	})
	// grouping all the routes with /api prefix
	api := app.Group("/api")
	// grouping all the routes with their names to make it more readable
	book := api.Group("/book")
	music := api.Group("/music")
	movie := api.Group("/movie")
	// Creating Book, Movie and Music
	book.Post("/createBook", controllers.CreateBook)
	movie.Post("/createMovie", controllers.CreateMovie)
	music.Post("/createMusic", controllers.CreateMusic)
	// Getting all the Books, Movies and Music
	book.Get("/getBooks", controllers.GetBooks)
	movie.Get("/getMovies", controllers.GetMovies)
	music.Get("/getMusic", controllers.GetAllMusic)
	// Getting a single Book, Movie and Music
	book.Get("/getBook/:id", controllers.GetBook)
	movie.Get("/getMovie/:id", controllers.GetMovie)
	music.Get("/getMusic/:id", controllers.GetMusic)
	// Deleting a single Book, Movie and Music
	book.Delete("/deleteBook/:id", controllers.DeleteBook)
	movie.Delete("/deleteMovie/:id", controllers.DeleteMovie)
	music.Delete("/deleteMusic/:id", controllers.DeleteMusic)
	// Updating a single Book, Movie and Music
	book.Put("/updateBook/:id", controllers.UpdateBook)
	movie.Put("/updateMovie/:id", controllers.UpdateMovie)
	music.Put("/updateMusic/:id", controllers.UpdateMusic)

}
