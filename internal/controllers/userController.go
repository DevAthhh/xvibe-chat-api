package controllers

import (
	"net/http"

	"github.com/DevAthhh/xvibe-chat/internal/auth"
	"github.com/DevAthhh/xvibe-chat/internal/repo"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var requestJson struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	if err := c.BindJSON(&requestJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var userRepo repo.UserRepo
	user, err := userRepo.Create(requestJson.Username, requestJson.Email, requestJson.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := auth.GenerateJWTToken(int(user.ID), user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"user_id": user.ID,
	})
}

func Profile(c *gin.Context) {

}
