package todos

import "test-todo/api/internal/domain/todos"

type Repository interface {
	FindAll() ([]todos.ToDo, error)
	Exists(string) (bool, error)
	Save(todos.ToDo) error
	Update(todos.ToDo) error
	Delete(string) error
}
