package application

import (
	domain "AntonyChR/golang-hexagonal-arch-template/internal/core/domain"
)

type TodoService struct {
	Repo domain.TodoRepository
}

func (s *TodoService) GetById(id string) domain.Todo {
	todo, _ := s.Repo.GetById(id)
	return todo
}

func (s *TodoService) GetAll() ([]domain.Todo, error) {
	todos, err := s.Repo.GetAll()
	return todos, err
}
func (s *TodoService) MarkAsDone(id string) error {
	err := s.Repo.Update(id, "Done", "1")
	return err
}
func (s *TodoService) Create(data domain.Todo) (string, error) {
	id, err := s.Repo.Create(data)
	return id, err
}
func (s *TodoService) DeleteById(id string) error {
	err := s.Repo.DeleteById(id)
	return err
}
