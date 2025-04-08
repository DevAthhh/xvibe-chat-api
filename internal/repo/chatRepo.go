package repo

import (
	"errors"

	"github.com/DevAthhh/xvibe-chat/internal/database"
	"github.com/DevAthhh/xvibe-chat/internal/models"
)

type ChatRepo struct {
}

func (c *ChatRepo) CreateChat(name, memberID, creatorID string) (*models.Chat, error) {
	if memberID == "" || creatorID == "" {
		return nil, errors.New("creatorID or memberID cannot be an empty value")
	}
	chat := models.Chat{
		Name: name,
	}

	var creator models.User
	var member models.User

	if result := database.DB.First(&creator, "id = ?", creatorID); result.Error != nil {
		return nil, result.Error
	}
	if result := database.DB.First(&member, "id = ?", memberID); result.Error != nil {
		return nil, result.Error
	}

	chat.Users = append(chat.Users, creator)
	chat.Users = append(chat.Users, member)

	if result := database.DB.Create(&chat); result.Error != nil {
		return nil, result.Error
	}
	return &chat, nil
}

func (c *ChatRepo) GetMessagesByChatID(chatID string) (*[]models.Message, error) {
	if chatID == "" {
		return nil, errors.New("chatID cannot be an empty value")
	}
	var msgs []models.Message
	result := database.DB.Preload("User").Where("chat_id = ?", chatID).Order("created_at").Find(&msgs)
	if result.Error != nil {
		return nil, result.Error
	}
	return &msgs, nil
}

func (c *ChatRepo) GetListOfMembers(chatID string) (*[]models.User, error) {
	if chatID == "" {
		return nil, errors.New("chatID cannot be an empty value")
	}
	var chats models.Chat
	if result := database.DB.Preload("Users").First(&chats, chatID); result.Error != nil {
		return nil, result.Error
	}

	return &chats.Users, nil
}

func (c *ChatRepo) DeleteMemberByID(chatID, userID string) error {
	if chatID == "" || userID == "" {
		return errors.New("chatID or userID cannot be an empty value")
	}
	var userChats models.UserChat
	if result := database.DB.Delete(&userChats, "user_id = ? AND chat_id = ?", userID, chatID); result.Error != nil {
		return nil
	}

	return nil
}
