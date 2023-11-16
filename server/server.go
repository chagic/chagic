package server

import "chagic/log"

type Server struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewServer() *Server {
	return &Server{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

var MyServer = NewServer()

func (s *Server) Start() {
	for {
		select {
		case client := <-s.Register:
			log.Logger.Info("Client register", log.Any("register", client.Name))
			s.Clients[client.Name] = client
		case client := <-s.Unregister:
			log.Logger.Info("Client unregister", log.Any("unregister", client.Name))
			if _, ok := s.Clients[client.Name]; ok {
				close(client.Send)
				delete(s.Clients, client.Name)
			}
		case message := <-s.Broadcast:
			for _, client := range s.Clients {
				client.Send <- message
			}
		}
	}
}
