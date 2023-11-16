package controller

import (
	"chagic/service"

	"github.com/gin-gonic/gin"
)

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var service service.UserService
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res := service.Login()
	ctx.JSON(200, res)

}

func Register(ctx *gin.Context) {
	var service service.UserService
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 注册
	res := service.Register()
	ctx.JSON(200, res)
}

func GetUserInfo(ctx *gin.Context) {
	userID, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	res := service.GetUserInfo(userID.(float64))
	ctx.JSON(200, res)

}
