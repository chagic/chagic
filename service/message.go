package service

import (
	"chagic/model"
	"chagic/serializer"
)

func SaveMessage(msg map[string]interface{}) {
	message := model.Message{
		Msg:    msg["msg"].(string),
		ChatID: int(msg["chat_id"].(float64)),
		Type:   msg["type"].(string),
		UserID: int(msg["user_id"].(float64)),
	}
	model.DB.Create(&message)
}

func ListMessages(query map[string]interface{}) serializer.Response {
	var result []model.Message

	model.DB.Model(&model.Message{}).Where(query).Find(&result)
	return serializer.ResultOK(result)
}
