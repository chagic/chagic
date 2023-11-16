package router

import (
	"chagic/log"
	"chagic/server"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upGrader = websocket.Upgrader{
	// Subprotocols: []string{"hey"},
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(c *gin.Context) {
	useId, exist := c.Get("userId")
	if !exist {
		return
	}
	user := fmt.Sprintf("%v", useId)
	log.Logger.Info("newUser", zap.String("newUser", user))
	h := http.Header{}
	for _, sub := range websocket.Subprotocols(c.Request) {
		if sub != "" {
			h.Set("Sec-Websocket-Protocol", sub)
			break
		}
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, h)
	if err != nil {
		return
	}

	client := &server.Client{
		Name: user,
		Conn: ws,
		Send: make(chan []byte),
	}

	server.MyServer.Register <- client
	go client.Read()
	go client.Write()

}
