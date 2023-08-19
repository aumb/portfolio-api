package portfolioRoutes

import (
	portfolioHandler "github.com/aumb/portfolio-api/internals/handlers/portfolio"
	"github.com/gofiber/fiber/v2"
)

func SetupPortfolioRoutes(router fiber.Router) {
	portfolio := router.Group("/portfolio")

	portfolio.Get("/", portfolioHandler.GetPortfolio)
}
