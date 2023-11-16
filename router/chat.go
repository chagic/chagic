package router

import (
	"chagic/controller"
	"chagic/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterChatRouter(group *gin.RouterGroup) {
	router := group.Group("/chat")
	router.Use(middleware.Jwt())
	router.POST("/create", controller.CreateChat)
	router.GET("/list", controller.ListChats)
}
