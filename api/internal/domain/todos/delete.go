package todos

func (s *service) Delete(id string) error {
	exists, err := s.repo.Exists(id)
	if err != nil {
		return ErrDeleteToDo
	}

	if !exists {
		return ErrToDoNotFound
	}

	err = s.repo.Delete(id)
	if err != nil {
		return ErrDeleteToDo
	}

	return nil
}
