package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-restful-api-mysql/controllers"
)

type Repository struct {
	Controller *controllers.Repository
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.Controller.CreateBook)
	api.Delete("/delete_book/:id", r.Controller.DeleteBook)
	api.Get("/get_book/:id", r.Controller.GetBookID)
	api.Get("/books", r.Controller.GetBooks)
}
