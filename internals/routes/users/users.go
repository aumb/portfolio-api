package usersRoutes

import (
	usersHandler "github.com/aumb/portfolio-api/internals/handlers/users"
	"github.com/gofiber/fiber/v2"
)

func SetupUsersRoutes(router fiber.Router) {
	users := router.Group("/users")

	users.Post("/", usersHandler.CreateUser)
	users.Put("/:userId", usersHandler.UpdateUser)
	users.Delete("/:userId", usersHandler.DeleteUser)
}
