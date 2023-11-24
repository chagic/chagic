package router

import (
	"chagic/controller"
	"chagic/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(group *gin.RouterGroup) {
	router := group.Group("/user")
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)
	router.Use(middleware.Jwt())
	router.GET("/", controller.GetUserInfo)
	// router.GET("/logout", controller.Logout)
	router.GET("list", controller.ListUsers)
}
