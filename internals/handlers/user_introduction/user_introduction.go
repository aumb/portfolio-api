package userIntroductionHandler

import (
	"github.com/aumb/portfolio-api/database"
	"github.com/aumb/portfolio-api/internals/model"
	"github.com/aumb/portfolio-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUserIntroduction(c *fiber.Ctx) error {
	db := database.DB
	userIntroduction := new(model.UserIntroduction)
	user := new(model.User)

	err := c.BodyParser(userIntroduction)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	errs := utils.Validate(userIntroduction)

	if len(errs) > 0 {
		errMsgs := utils.ParseValidationErrors(errs)

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: errMsgs,
		}
	}

	db.Find(&user, "id = ?", userIntroduction.UserID)

	if user.ID == uuid.Nil {
		errMsg := "User does not exist"

		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: errMsg,
		}
	}

	err = db.Create(&userIntroduction).Error

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	return c.Status(fiber.StatusCreated).JSON(userIntroduction)
}

func UpdateUserIntroduction(c *fiber.Ctx) error {
	type updateUserIntroduction struct {
		Title       string `validate:"required" json:"title" form:"title"`
		Description string `validate:"required" json:"description" form:"description"`
	}

	db := database.DB
	var userIntroduction model.UserIntroduction

	id := c.Params("userIntroductionId")

	db.Find(&userIntroduction, "id = ?", id)

	if userIntroduction.ID == uuid.Nil {
		errMsg := "Social media profile does not exist"

		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: errMsg,
		}
	}

	var updateUserIntroductionData updateUserIntroduction

	err := c.BodyParser(&updateUserIntroductionData)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	errs := utils.Validate(updateUserIntroductionData)

	if len(errs) > 0 {
		errMsgs := utils.ParseValidationErrors(errs)

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: errMsgs,
		}
	}

	userIntroduction.Title = updateUserIntroductionData.Title
	userIntroduction.Description = updateUserIntroductionData.Description

	db.Save(&userIntroduction)

	return c.JSON(userIntroduction)
}

func DeleteUserIntroduction(c *fiber.Ctx) error {
	db := database.DB
	var userIntroduction model.UserIntroduction

	id := c.Params("userIntroductionId")

	db.Find(&userIntroduction, "id = ?", id)

	if userIntroduction.ID == uuid.Nil {
		errMsg := "User introduction about does not exist"

		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: errMsg,
		}
	}

	err := db.Delete(&userIntroduction, "id = ?", id).Error

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	return c.Send(nil)
}
