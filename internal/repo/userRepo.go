package repo

import (
	"github.com/DevAthhh/xvibe-chat/internal/database"
	"github.com/DevAthhh/xvibe-chat/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
}

func (ur *UserRepo) Create(username, email, password string) (*models.User, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: username,
		Password: string(passwordHash),
		Email:    email,
		IsBanned: false,
	}
	if result := database.DB.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
