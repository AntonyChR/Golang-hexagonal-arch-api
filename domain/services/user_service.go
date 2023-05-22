package services

import (
	entities "chat/domain/entities"
	repositories "chat/domain/repositories"
	security "chat/infrastructure/security"

	uuid "github.com/google/uuid"
)

func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{userRepository: repository}
}

type UserService struct {
	userRepository repositories.UserRepository
}

func (s *UserService) GetUserBydId(id string) (*entities.User, error) {
	return s.userRepository.GetById(id)
}

func (s *UserService) GetAllUsers() ([]*entities.User, error) {
	return s.userRepository.GetAll()
}

func (s *UserService) CreateUser(name string, email string, password string) (*entities.User, error) {
	id := uuid.New()
	user := &entities.User{
		ID:       id.String(),
		Name:     name,
		Email:    email,
		Password: security.HashPassword(password),
	}
	err := s.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *UserService) DeleteUserById(id string) error {
	return s.userRepository.DeleteById(id)
}

func (s *UserService) GetBy(field, value string) (*entities.User, error) {
	return s.userRepository.GetBy(field, value)
}
