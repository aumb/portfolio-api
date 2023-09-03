package userAboutRoutes

import (
	userAboutHandler "github.com/aumb/portfolio-api/internals/handlers/user_about"
	"github.com/gofiber/fiber/v2"
)

func SetupUserAboutRoutes(router fiber.Router) {
	userAbout := router.Group("/user-about")

	userAbout.Post("/", userAboutHandler.CreateUserAbout)
	userAbout.Put("/:userAboutId", userAboutHandler.UpdateUserAbout)
	userAbout.Delete("/:userAboutId", userAboutHandler.DeleteUserAbout)

}
