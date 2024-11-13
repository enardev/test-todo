package todos

import "context"

func (s *service) GetAll(ctx context.Context) ([]ToDo, error) {

	todos, err := s.repo.FindAll(ctx)
	if err != nil {
		return []ToDo{}, ErrToDoNotFound
	}

	return todos, nil
}
