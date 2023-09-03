package main

import (
	"log"

	"github.com/aumb/portfolio-api/database"
	"github.com/aumb/portfolio-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Panic(err)
	}
}
