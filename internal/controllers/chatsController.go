package controllers

import (
	"net/http"

	"github.com/DevAthhh/xvibe-chat/internal/repo"
	"github.com/gin-gonic/gin"
)

var chatRepo repo.ChatRepo

func CreateChat(c *gin.Context) {
	var requestData struct {
		MemberID string `json:"member_id"`
		NameChat string `json:"name_chat"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	author, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	chat, err := chatRepo.CreateChat(requestData.NameChat, requestData.MemberID, author)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": chat.ID,
	})
}

func GetListOfMembers(c *gin.Context) {
	var requestData struct {
		ChatID string `json:"chat_id"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	members, err := chatRepo.GetListOfMembers(requestData.ChatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"members": *members,
	})

}

func DeleteMember(c *gin.Context) {
	var requestData struct {
		UserID string `json:"user_id"`
		ChatID string `json:"chat_id"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := chatRepo.DeleteMemberByID(requestData.ChatID, requestData.UserID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "the user has b een deleted",
	})
}
