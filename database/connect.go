package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aumb/portfolio-api/config"
	"github.com/aumb/portfolio-api/internals/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&model.Portfolio{}, &model.User{}, &model.UserSocialMedia{}, &model.UserAbout{}, &model.UserIntroduction{})
	fmt.Println("Database Migrated")
}
