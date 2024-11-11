package todos

import "errors"

var (
	ErrToDoAlreadyExists = errors.New("todo already exists")
	ErrToDoNotFound      = errors.New("todo not found")
	ErrSaveToDo          = errors.New("error saving todo")
	ErrUpdateToDo        = errors.New("error updating todo")
	ErrDeleteToDo        = errors.New("error deleting todo")
)
