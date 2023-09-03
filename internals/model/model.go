package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommonModelFields struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (u *CommonModelFields) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}

type Portfolio struct {
	CommonModelFields
	UserID uuid.UUID `gorm:"type:uuid"`
	User   User      `gorm:"embedded" json:"user"`
}

type User struct {
	CommonModelFields
	FirstName string `validate:"required" json:"first_name" form:"first_name"`
	LastName  string `validate:"required" json:"last_name" form:"last_name"`
	Email     string `validate:"required,email" json:"email"`
}

type UserSocialMedia struct {
	CommonModelFields
	UserID       uuid.UUID `gorm:"type:uuid" validate:"required" json:"user_id" form:"user_id"`
	FacebookUrl  string    `validate:"omitempty,http_url" json:"facebook_url" form:"facebook_url"`
	LinkedInUrl  string    `validate:"omitempty,http_url" json:"linked_in_url" form:"linked_in_url"`
	TwitterUrl   string    `validate:"omitempty,http_url" json:"twitter_url" form:"twitter_url"`
	GitHubUrl    string    `validate:"omitempty,http_url" json:"git_hub_url" form:"git_hub_url"`
	InstagramUrl string    `validate:"omitempty,http_url" json:"instagram_url" form:"instagram_url"`
}

type UserAbout struct {
	CommonModelFields
	UserID      uuid.UUID `gorm:"type:uuid" validate:"required" json:"user_id" form:"user_id"`
	Title       string    `validate:"required" json:"title" form:"title"`
	Description string    `validate:"required" json:"description" form:"description"`
}
