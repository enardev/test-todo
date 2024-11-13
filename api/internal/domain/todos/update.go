package todos

import (
	"context"
	"time"
)

func (s *service) Update(ctx context.Context, todo ToDo) (ToDo, error) {
	exists, err := s.repo.Exists(ctx, todo.ID)
	if err != nil {
		return ToDo{}, ErrUpdateToDo
	}

	if !exists {
		return ToDo{}, ErrToDoNotFound
	}

	oldToDo, err := s.repo.FindByID(ctx, todo.ID)
	if err != nil {
		return ToDo{}, err
	}

	todo = compareAndReplace(oldToDo, todo)

	err = s.repo.Update(ctx, todo)
	if err != nil {
		return ToDo{}, ErrUpdateToDo
	}

	return todo, nil
}

func compareAndReplace(oldToDo ToDo, newToDo ToDo) ToDo {
	if newToDo.Title != "" {
		oldToDo.Title = newToDo.Title
	}

	if newToDo.Completed != oldToDo.Completed {
		oldToDo.Completed = newToDo.Completed
	}

	oldToDo.UpdatedAt = time.Now()
	return oldToDo
}
