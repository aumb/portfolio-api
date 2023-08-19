package main

import (
	"github.com/aumb/portfolio-api/database"
	"github.com/aumb/portfolio-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
