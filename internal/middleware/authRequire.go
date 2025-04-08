package middleware

import (
	"net/http"

	"github.com/DevAthhh/xvibe-chat/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthRequire(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "lack of a token",
		})
		return
	}

	if _, err := auth.ValidateJWTToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Next()
}


