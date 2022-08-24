package repositorie

import (
	"errors"

	"github.com/google/uuid"

	"AntonyChR/golang-hexagonal-arch-template/internal/core/domain"
)

type MockMemoryTodoRepository struct {
	data  []domain.Todo
	Table string
}

func (repo *MockMemoryTodoRepository) GetById(id string) (domain.Todo, error) {
	for _, v := range repo.data {
		if v.ID == id {
			return v, nil
		}
	}
	err := errors.New("todo not found")
	return domain.Todo{}, err
}

func (repo *MockMemoryTodoRepository) GetAll() ([]domain.Todo, error) {
	if len(repo.data) < 1 {
		return []domain.Todo{}, nil
	}
	return repo.data, nil
}

func (repo *MockMemoryTodoRepository) Create(data domain.Todo) (string, error) {
	id := uuid.New().String()
	t := domain.Todo{
		ID:          id,
		Title:       data.Title,
		Description: data.Description,
		Date:        data.Date,
		Done:        data.Done,
	}

	repo.data = append(repo.data, t)
	return id, nil
}

func (repo *MockMemoryTodoRepository) Update(id, field string, value string) error {
	return nil
}

func (repo *MockMemoryTodoRepository) DeleteById(id string) error {
	for i, v := range repo.data {
		if id == v.ID {
			repo.data = append(repo.data[:i], repo.data[i+1:]...)
		}
	}
	return nil
}

func InitializeMemoryRepository(table string, database string) *MockMemoryTodoRepository {
	c1 := domain.Todo{Title: "title-1", ID: "id-1", Description: "desc-1", Date: "11/11/2021", Done: false}
	c2 := domain.Todo{Title: "title-2", ID: "asd-2", Description: "desc-2", Date: "11/11/2022", Done: false}
	initialData := []domain.Todo{c1, c2}
	return &MockMemoryTodoRepository{
		data:  initialData,
		Table: table,
	}
}
