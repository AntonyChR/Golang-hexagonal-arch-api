package ws

import (
	"chat/domain/entities"
	"chat/domain/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Server(store *ClientStore, ctx *gin.Context, msgService *services.MessageService) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer conn.Close()
	clientId, _ := ctx.Get("userId")
	store.RegisterClient(conn, clientId.(string))
	//store.RegisterClient(conn, uuid.NewString())
	for {
		var msg entities.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		msgService.SaveMessage(&msg)

		store.send <- msg

	}
}
