package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Name     string `gorm:"size:100;not null"`
	Users    []User `gorm:"many2many:user_chats"`
	Messages []Message
}

type Message struct {
	gorm.Model
	ChatID  uint   `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
	Content string `gorm:"type:text;not null"`
	User    User   `gorm:"foreignKey:UserID"`
	Chat    Chat   `gorm:"foreignKey:ChatID"`
}
