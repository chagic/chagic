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

// ListChats godoc
// @Summary      List Chats
// @Description  get chats
// @Tags         chat
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {object}   model.Chat "{"code": 200,"data": model.Chat , "msg": "string","success": true}"
// @Router       /chat/list [get]
func ListChats(ctx *gin.Context) {
	userID, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	res := service.ListChats(userID.(float64))
	ctx.JSON(200, res)
}
