package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/aumb/portfolio-api/database"
	"github.com/aumb/portfolio-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":3000"); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
}
