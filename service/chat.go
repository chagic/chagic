package service

import (
	"chagic/conf"
	"chagic/model"
	"chagic/serializer"
	"errors"
)

func CreateChat(userIds []int) serializer.Response {
	count := len(userIds)
	var chatType model.ChatType
	if count == 2 {
		chat, err := chatExists(userIds[0], userIds[1])
		if err == nil {
			return serializer.ResultOK(chat)
		}
		chatType = model.ChatTypePrivate
	}
	if count > 2 {
		chatType = model.ChatTypeGroup
	}

	chat := model.Chat{
		Type:             chatType,
		Name:             "",
		ParticipantCount: count,
		// ActiveUsers:      fmt.Sprintf("%d,%d", userIds...),
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
	var chats []model.Chat
	var participants []model.Participant
	tablePrefix := conf.GetConfig().MySQL.TablePrefix
	model.DB.Where("user_id = ?", userID).Joins("JOIN " + tablePrefix + "chat ON " + tablePrefix + "participant.chat_id = " + tablePrefix + "chat.id").Find(&participants)
	for _, participant := range participants {
		var chat model.Chat
		model.DB.First(&chat, participant.ChatID)

		if chat.Type == model.ChatTypePrivate {
			anotherUser := GetChatAnotherUser(chat.ID, int(userID))
			chat.Name = anotherUser.Username
		}

		// Append the modified chat entry to the result list
		chats = append(chats, chat)
	}

	return serializer.ResultOK(chats)
}

func chatExists(user1ID, user2ID int) (model.Chat, error) {
	var chat model.Chat
	tablePrefix := conf.GetConfig().MySQL.TablePrefix

	result := model.DB.Where("type = ?", model.ChatTypePrivate).
		Where("participant_count = ?", 2).
		Joins("JOIN "+tablePrefix+"participant p1 ON p1.chat_id = "+tablePrefix+"chat.id").
		Joins("JOIN "+tablePrefix+"participant p2 ON p2.chat_id = "+tablePrefix+"chat.id").
		Where("p1.user_id = ?", user1ID).
		Where("p2.user_id = ?", user2ID).
		First(&chat)

	if result.RowsAffected > 0 {
		return chat, nil // Chat already exists
	}
	return chat, errors.New("Chat does not exist")
}

func GetChatRoomUserList(chatID uint) ([]model.User, error) {
	var participants []model.Participant
	tablePrefix := conf.GetConfig().MySQL.TablePrefix
	if err := model.DB.Preload(tablePrefix+"user").Where("chat_id = ?", chatID).Find(&participants).Error; err != nil {
		return nil, err
	}

	var users []model.User
	for _, participant := range participants {
		users = append(users, participant.User)
	}

	return users, nil
}

func GetChatAnotherUser(chatID int, userID int) model.User {
	tablePrefix := conf.GetConfig().MySQL.TablePrefix
	var participant model.Participant

	model.DB.Table(tablePrefix+"participant").
		Preload("User").
		Where("chat_id = ?", chatID).
		Not("user_id = ?", userID).
		Find(&participant)

	return participant.User
}
