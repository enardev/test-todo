package todos

import "time"

func (s *service) Update(todo ToDo) (ToDo, error) {
	exists, err := s.repo.Exists(todo.ID)
	if err != nil {
		return ToDo{}, ErrUpdateToDo
	}

	if !exists {
		return ToDo{}, ErrToDoNotFound
	}

	todo.UpdatedAt = time.Now()

	err = s.repo.Update(todo)
	if err != nil {
		return ToDo{}, ErrUpdateToDo
	}

	return todo, nil
}
