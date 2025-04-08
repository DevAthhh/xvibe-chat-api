package database

import (
	"fmt"
	"os"

	"github.com/DevAthhh/xvibe-chat/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() error {
	var err error

	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func SyncDB() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Chat{},
		&models.Message{},
		&models.UserChat{},
	)
}
