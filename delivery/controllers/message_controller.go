package controllers

import (
	"chat/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewMessageController(service *services.MessageService) *MessageController {
	return &MessageController{messageService: *service}
}

type MessageController struct {
	messageService services.MessageService
}

func (c *MessageController) GetChat(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")

	if from == "" || to == "" {
		ctx.JSON(http.StatusBadRequest, "Invalid query params")
		return
	}

	userId, _ := ctx.Get("userId")

	if userId.(string) != from {
		ctx.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	data, err := c.messageService.GetChat(from, to)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Error getting chat messages")
		return
	}

	ctx.JSON(http.StatusOK, data)
}
