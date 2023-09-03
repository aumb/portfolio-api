package userSocialMediaRoutes

import (
	userSocialMediaHandler "github.com/aumb/portfolio-api/internals/handlers/user_social_media"
	"github.com/gofiber/fiber/v2"
)

func SetupUserSocialMediaRoutes(router fiber.Router) {
	userSocialMedia := router.Group("/user-social-media")

	userSocialMedia.Post("/", userSocialMediaHandler.CreateUserSocialMedia)
	userSocialMedia.Put("/:userSocialMediaId", userSocialMediaHandler.UpdateUserSocialMedia)
	userSocialMedia.Delete("/:userSocialMediaId", userSocialMediaHandler.DeleteUserSocialMedia)

}
