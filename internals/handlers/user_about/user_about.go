package userAboutHandler

import (
	"github.com/aumb/portfolio-api/database"
	"github.com/aumb/portfolio-api/internals/model"
	"github.com/aumb/portfolio-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUserAbout(c *fiber.Ctx) error {
	db := database.DB
	userAbout := new(model.UserAbout)
	user := new(model.User)

	err := c.BodyParser(userAbout)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	errs := utils.Validate(userAbout)

	if len(errs) > 0 {
		errMsgs := utils.ParseValidationErrors(errs)

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: errMsgs,
		}
	}

	db.Find(&user, "id = ?", userAbout.UserID)

	if user.ID == uuid.Nil {
		errMsg := "User does not exist"

		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: errMsg,
		}
	}

	err = db.Create(&userAbout).Error

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	return c.Status(fiber.StatusCreated).JSON(userAbout)
}

func UpdateUserAbout(c *fiber.Ctx) error {
	type updateUserAbout struct {
		Title       string `validate:"required" json:"title" form:"title"`
		Description string `validate:"required" json:"description" form:"description"`
	}

	db := database.DB
	var userAbout model.UserAbout

	id := c.Params("userAboutId")

	db.Find(&userAbout, "id = ?", id)

	if userAbout.ID == uuid.Nil {
		errMsg := "Social media profile does not exist"

		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: errMsg,
		}
	}

	var updateUserAboutData updateUserAbout

	err := c.BodyParser(&updateUserAboutData)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	errs := utils.Validate(updateUserAboutData)

	if len(errs) > 0 {
		errMsgs := utils.ParseValidationErrors(errs)

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: errMsgs,
		}
	}

	userAbout.Title = updateUserAboutData.Title
	userAbout.Description = updateUserAboutData.Description

	db.Save(&userAbout)

	return c.JSON(userAbout)
}

func DeleteUserAbout(c *fiber.Ctx) error {
	db := database.DB
	var userAbout model.UserAbout

	id := c.Params("userAboutId")

	db.Find(&userAbout, "id = ?", id)

	if userAbout.ID == uuid.Nil {
		errMsg := "User about does not exist"

		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: errMsg,
		}
	}

	err := db.Delete(&userAbout, "id = ?", id).Error

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	return c.Send(nil)
}
