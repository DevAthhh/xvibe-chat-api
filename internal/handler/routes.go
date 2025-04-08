package handler

import (
	"github.com/DevAthhh/xvibe-chat/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/u")
			{
				auth.POST("/sign-up", controllers.SignUp)
				auth.GET("/:id", controllers.Profile)
			}

			chats := v1.Group("/chat")
			{
				chats.POST("/", controllers.CreateChat)
				chats.POST("/members", controllers.GetListOfMembers)
				chats.DELETE("/members", controllers.DeleteMember)
			}
		}
	}

	return router
}
