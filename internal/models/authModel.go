package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string `gorm:"unique"`
	IsBanned bool
	Chats    []Chat `gorm:"many2many:user_chats"`
}

type UserChat struct {
	UserID uint `gorm:"primaryKey"`
	ChatID uint `gorm:"primaryKey"`
}
