package repositories

import entities "chat/domain/entities"

type MessageRepository interface {
	GetChat(from string, to string) ([]*entities.Message, error)
	DeleteMessage(id string) error
	DeleteChat(from string, to string) error
	Save(message *entities.Message) error
}
