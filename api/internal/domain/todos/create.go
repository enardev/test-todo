package todos

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (s *service) Create(ctx context.Context, todo ToDo) (ToDo, error) {
	if todo.ID == "" {
		todo.ID = uuid.New().String()
	}

	exists, err := s.repo.Exists(ctx, todo.ID)
	if err != nil {
		return ToDo{}, ErrSaveToDo
	}

	if exists {
		return ToDo{}, ErrToDoAlreadyExists
	}

	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	err = s.repo.Save(ctx, todo)
	if err != nil {
		return ToDo{}, ErrSaveToDo
	}

	return todo, nil
}
