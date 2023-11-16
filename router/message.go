package router

import (
	"chagic/controller"
	"chagic/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterMessageRouter(group *gin.RouterGroup) {
	router := group.Group("/message")
	router.Use(middleware.Jwt())
	router.Use(middleware.HeaderMiddleware())
	router.GET("/list", controller.ListMessages)
}
