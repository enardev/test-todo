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

	oldToDo, err := s.repo.FindByID(todo.ID)
	if err != nil {
		return ToDo{}, err
	}

	todo = compareAndReplace(oldToDo, todo)

	err = s.repo.Update(todo)
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
