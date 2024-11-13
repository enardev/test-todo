package todos

import "context"

func (s *service) Delete(ctx context.Context, id string) error {
	exists, err := s.repo.Exists(ctx, id)
	if err != nil {
		return ErrDeleteToDo
	}

	if !exists {
		return ErrToDoNotFound
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return ErrDeleteToDo
	}

	return nil
}
