package services

import (
	"chat/domain/entities"
	"chat/domain/repositories"
)

func NewMessageService(repository repositories.MessageRepository) *MessageService {
	return &MessageService{messageRepository: repository}
}

type MessageService struct {
	messageRepository repositories.MessageRepository
}

func (s *MessageService) GetChat(from, to string) ([]*entities.Message, error) {
	return s.messageRepository.GetChat(from, to)
}

func (s *MessageService) SaveMessage(m *entities.Message) error {
	return s.messageRepository.Save(m)
}
