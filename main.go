package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	config2 "go-restful-api-mysql/config"
	"go-restful-api-mysql/controllers"
	"go-restful-api-mysql/models"
	"go-restful-api-mysql/routes"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &config2.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := config2.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db.")
	}

	repository := &controllers.Repository{
		DB: db,
	}

	routesRepository := &routes.Repository{
		Controller: repository,
	}

	app := fiber.New()
	routesRepository.SetupRoutes(app)
	app.Listen(":8080")

}
