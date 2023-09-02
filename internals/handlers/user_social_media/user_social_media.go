package userSocialMediaHandler

import (
	"github.com/aumb/portfolio-api/database"
	"github.com/aumb/portfolio-api/internals/model"
	"github.com/aumb/portfolio-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUserSocialMedia(c *fiber.Ctx) error {
	db := database.DB
	userSocialMedia := new(model.UserSocialMedia)
	user := new(model.User)

	err := c.BodyParser(userSocialMedia)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	errs := utils.Validate(userSocialMedia)

	if len(errs) > 0 {
		errMsgs := utils.ParseValidationErrors(errs)

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: errMsgs,
		}
	}

	db.Find(&user, "id = ?", userSocialMedia.UserID)

	if user.ID == uuid.Nil {
		errMsg := "User does not exist"

		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: errMsg,
		}
	}

	err = db.Create(&userSocialMedia).Error

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	return c.Status(fiber.StatusCreated).JSON(userSocialMedia)
}

func UpdateUserSocialMedia(c *fiber.Ctx) error {
	type updateUserSocialMedia struct {
		FacebookUrl  string `validate:"omitempty,http_url" json:"facebook_url" form:"facebook_url"`
		LinkedInUrl  string `validate:"omitempty,http_url" json:"linked_in_url" form:"linked_in_url"`
		TwitterUrl   string `validate:"omitempty,http_url" json:"twitter_url" form:"twitter_url"`
		GitHubUrl    string `validate:"omitempty,http_url" json:"git_hub_url" form:"git_hub_url"`
		InstagramUrl string `validate:"omitempty,http_url" json:"instagram_url" form:"instagram_url"`
	}

	db := database.DB
	var userSocialMedia model.UserSocialMedia

	id := c.Params("userSocialMediaId")

	db.Find(&userSocialMedia, "id = ?", id)

	if userSocialMedia.ID == uuid.Nil {
		errMsg := "Social media profile does not exist"

		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: errMsg,
		}
	}

	var updateUserSocialMediaData updateUserSocialMedia

	err := c.BodyParser(&updateUserSocialMediaData)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	errs := utils.Validate(updateUserSocialMediaData)

	if len(errs) > 0 {
		errMsgs := utils.ParseValidationErrors(errs)

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: errMsgs,
		}
	}

	userSocialMedia.FacebookUrl = updateUserSocialMediaData.FacebookUrl
	userSocialMedia.LinkedInUrl = updateUserSocialMediaData.LinkedInUrl
	userSocialMedia.TwitterUrl = updateUserSocialMediaData.TwitterUrl
	userSocialMedia.GitHubUrl = updateUserSocialMediaData.GitHubUrl
	userSocialMedia.InstagramUrl = updateUserSocialMediaData.InstagramUrl

	db.Save(&userSocialMedia)

	return c.JSON(userSocialMedia)
}

func DeleteUserSocialMedia(c *fiber.Ctx) error {
	db := database.DB
	var userSocialMedia model.UserSocialMedia

	id := c.Params("userSocialMediaId")

	db.Find(&userSocialMedia, "id = ?", id)

	if userSocialMedia.ID == uuid.Nil {
		errMsg := "Social media profile does not exist"

		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: errMsg,
		}
	}

	err := db.Delete(&userSocialMedia, "id = ?", id).Error

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	return c.Send(nil)
}
