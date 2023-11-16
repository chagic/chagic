package controller

import (
	"chagic/service"

	"github.com/gin-gonic/gin"
)

func ListMessages(ctx *gin.Context) {
	_, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	var query map[string]interface{}
	ctx.BindQuery(&query)
	res := service.ListMessages(query)
	ctx.JSON(200, res)
}
