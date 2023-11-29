package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-restful-api-mysql/entities"
	"go-restful-api-mysql/models"
	"gorm.io/gorm"
	"net/http"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := entities.Book{}

	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book."})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "create book successfully"})
	return nil

}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id connot be empty."})
		return nil
	}

	err := r.DB.Delete(bookModel, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete book."})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "delete book successfully"})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "books fatched books",
			"data":    bookModels,
		})
	return nil
}

func (r *Repository) GetBookID(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := &models.Books{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id connot empty"})
		return nil
	}

	fmt.Sprintln("the ID is", id)
	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book fatcher by id successfully",
		"data":    bookModel,
	})
	return nil
}
