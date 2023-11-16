package server

import (
	"chagic/log"
	"chagic/service"
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
	Name string
}

func (c *Client) Read() {
	defer func() {
		MyServer.Unregister <- c
		c.Conn.Close()
	}()
	for {
		c.Conn.PongHandler()
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Logger.Error("read message error:", log.Any("err", err))
			MyServer.Unregister <- c
			c.Conn.Close()
			break
		}
		var jsonMsg map[string]interface{}
		json.Unmarshal(msg, &jsonMsg)
		// message 入库
		saveMessage(jsonMsg)
		chatId := jsonMsg["chat_id"].(float64)
		if chatId == 0 {
			MyServer.Broadcast <- msg
		} else {
			// query user or group user client
			userIds := service.GetUsersOnChat(chatId)
			for _, userId := range userIds {
				client, ok := MyServer.Clients[userId]
				if ok {
					client.Send <- msg
				}
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, message)
	}
}

func saveMessage(msg map[string]interface{}) {
	service.SaveMessage(msg)
}
