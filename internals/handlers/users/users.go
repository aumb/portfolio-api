package usersHandler

import (
	"github.com/aumb/portfolio-api/database"
	"github.com/aumb/portfolio-api/internals/model"
	"github.com/aumb/portfolio-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	err := c.BodyParser(user)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	errs := utils.Validate(user)

	if len(errs) > 0 {
		errMsgs := utils.ParseValidationErrors(errs)

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: errMsgs,
		}
	}

	user.ID = uuid.New()

	err = db.Create(&user).Error

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	id := c.Params("userId")

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		errMsg := "User does not exist"

		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: errMsg,
		}
	}

	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	return c.Send(nil)
}
