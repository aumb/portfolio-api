package userIntroductionRoutes

import (
	userIntroductionHandler "github.com/aumb/portfolio-api/internals/handlers/user_introduction"
	"github.com/gofiber/fiber/v2"
)

func SetupUserIntroductionRoutes(router fiber.Router) {
	userIntroduction := router.Group("/user-introduction")

	userIntroduction.Post("/", userIntroductionHandler.CreateUserIntroduction)
	userIntroduction.Put("/:userIntroductionId", userIntroductionHandler.UpdateUserIntroduction)
	userIntroduction.Delete("/:userIntroductionId", userIntroductionHandler.DeleteUserIntroduction)

}
