package repositories

import entities "chat/domain/entities"

type UserRepository interface {
	GetById(id string) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Save(user *entities.User) error
	DeleteById(id string) error
	GetBy(field string, value string) (*entities.User, error)
}
