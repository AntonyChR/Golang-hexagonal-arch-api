package domain

type TodoRepository interface {
	GetById(id string) (Todo, error)
	GetAll() ([]Todo, error)
	Create(data Todo) (string, error)
	Update(id string, field string, value string) error
	DeleteById(id string) error
}
