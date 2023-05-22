package ws

import (
	"chat/domain/entities"
	"fmt"

	websocket "github.com/gorilla/websocket"
)

func NewWesocketStore() *ClientStore {
	return &ClientStore{
		clients: make(map[string]*websocket.Conn),
		send:    make(chan entities.Message),
	}
}

type ClientStore struct {
	clients map[string]*websocket.Conn
	send    chan entities.Message
}

func (c *ClientStore) RegisterClient(clientConn *websocket.Conn, id string) {
	if _, exists := c.clients[id]; exists {
		delete(c.clients, id)
	}
	c.clients[id] = clientConn
}

func (c *ClientStore) Run() {
	for {
		msg := <-c.send
		if conn, exists := c.clients[msg.To]; exists {
			err := conn.WriteJSON(msg)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
