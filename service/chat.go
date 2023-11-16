package service

import (
	"chagic/model"
	"chagic/serializer"
)

func CreateChat(userIds []int) serializer.Response {
	count := len(userIds)
	var chatType model.ChatType
	if count == 2 {
		chatType = model.ChatTypePrivate
	}
	if count > 2 {
		chatType = model.ChatTypeGroup
	}
	chat := model.Chat{
		Type:             chatType,
		Name:             "",
		ParticipantCount: count,
	}
	model.DB.Create(&chat)
	var participants []model.Participant
	for _, id := range userIds {
		participant := model.Participant{
			ChatID: chat.ID,
			UserID: id,
			Role:   model.Member,
		}
		participants = append(participants, participant)
	}
	model.DB.Create(&participants)

	return serializer.ResultOK(chat)

}

func ListChats(userID float64) serializer.Response {
	var results []model.Chat
	var chat_ids []int
	model.DB.Model(&model.Participant{}).Where("user_id = ?", userID).Select("chat_id").Find(&chat_ids)
	model.DB.Model(&model.Chat{}).Where("id in (?)", chat_ids).Find(&results)

	// SQL:
	// model.DB.Model(&model.Chat{}).
	// 	Joins("INNER JOIN chagic_participant p ON chagic_chat.id = p.chat_id and p.user_id = ?", userID).
	// 	Find(&results)
	return serializer.ResultOK(results)
}
