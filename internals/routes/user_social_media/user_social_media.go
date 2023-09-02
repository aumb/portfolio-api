package userSocialMediaRoutes

import (
	userSocialMediaHandler "github.com/aumb/portfolio-api/internals/handlers/user_social_media"
	"github.com/gofiber/fiber/v2"
)

func SetupUserSocialMediaRoutes(router fiber.Router) {
	users := router.Group("/user-social-media")

	users.Post("/", userSocialMediaHandler.CreateUserSocialMedia)
	users.Put("/:userSocialMediaId", userSocialMediaHandler.UpdateUserSocialMedia)
	users.Delete("/:userSocialMediaId", userSocialMediaHandler.DeleteUserSocialMedia)

}
