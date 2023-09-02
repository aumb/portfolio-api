package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Portfolio struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid"`
	UserID uuid.UUID `gorm:"type:uuid"`
	User   User      `json:"user"`
}

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	FirstName string    `validate:"required" json:"first_name"`
	LastName  string    `validate:"required" json:"last_name"`
	Email     string    `validate:"required" json:"email"`
}

type UserSocialMedia struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid"`
	UserID       uuid.UUID `gorm:"type:uuid"`
	FacebookUrl  string    `json:"facebook_url"`
	LinkedInUrl  string    `json:"linked_in_url"`
	TwitterUrl   string    `json:"twitter_url"`
	GitHubUrl    string    `json:"git_hub_url"`
	InstagramUrl string    `json:"instagram_url"`
}
