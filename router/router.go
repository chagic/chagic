package router

import (
	"chagic/conf"
	"chagic/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	gin.SetMode(conf.GetConfig().Server.Mode)
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "Success")
		})
		RegisterUserRouter(v1)
		RegisterChatRouter(v1)
		RegisterMessageRouter(v1)
		// RegisterSocketIORouter(r)

		v1.GET("ws", middleware.Jwt(), WsHandler)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
