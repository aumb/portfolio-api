package router

import (
	portfolioRoutes "github.com/aumb/portfolio-api/internals/routes/portfolio"
	usersRoutes "github.com/aumb/portfolio-api/internals/routes/users"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	portfolioRoutes.SetupPortfolioRoutes(api)
	usersRoutes.SetupUsersRoutes(api)
}
