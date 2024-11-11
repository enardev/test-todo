package todos

func (s *service) GetAll() ([]ToDo, error) {

	todos, err := s.repo.FindAll()
	if err != nil {
		return []ToDo{}, ErrToDoNotFound
	}

	return todos, nil
}
