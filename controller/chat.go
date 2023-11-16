package controller

import (
	"chagic/service"

	"github.com/gin-gonic/gin"
)

type CreateParams struct {
	Ids []int `json:"ids" binding:"required"`
}

func CreateChat(ctx *gin.Context) {
	// userID, ok := ctx.Get("userId")
	// if !ok {
	// 	ctx.JSON(401, gin.H{"error": "unauthorized"})
	// 	return
	// }
	var input CreateParams
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res := service.CreateChat(input.Ids)
	ctx.JSON(200, res)
}

func ListChats(ctx *gin.Context) {
	userID, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	res := service.ListChats(userID.(float64))
	ctx.JSON(200, res)
}
