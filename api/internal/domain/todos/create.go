package todos

import (
	"time"

	"github.com/google/uuid"
)

func (s *service) Create(todo ToDo) (ToDo, error) {
	if todo.ID == "" {
		todo.ID = uuid.New().String()
	}

	exists, err := s.repo.Exists(todo.ID)
	if err != nil {
		return ToDo{}, ErrSaveToDo
	}

	if exists {
		return ToDo{}, ErrToDoAlreadyExists
	}

	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	err = s.repo.Save(todo)
	if err != nil {
		return ToDo{}, ErrSaveToDo
	}

	return todo, nil
}
