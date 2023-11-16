package service

import "chagic/model"

func GetUsersOnChat(chatId float64) []string {
	var result []string
	model.DB.Model(&model.Participant{}).
		Where("chat_id = ?", chatId).
		Select("user_id").
		Find(&result)
	return result
}
