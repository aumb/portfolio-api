package portfolioHandler

import (
	"github.com/aumb/portfolio-api/database"
	"github.com/aumb/portfolio-api/internals/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetPortfolio(c *fiber.Ctx) error {
	db := database.DB
	var portfolio model.Portfolio

	db.Find(&portfolio)

	if portfolio.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	return c.JSON(portfolio)
}
