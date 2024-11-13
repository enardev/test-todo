package todos

import "context"

type Repository interface {
	FindAll(context.Context) ([]ToDo, error)
	FindByID(context.Context, string) (ToDo, error)
	Exists(context.Context, string) (bool, error)
	Save(context.Context, ToDo) error
	Update(context.Context, ToDo) error
	Delete(context.Context, string) error
}

type Service interface {
	GetAll(context.Context) ([]ToDo, error)
	Create(context.Context, ToDo) (ToDo, error)
	Update(context.Context, ToDo) (ToDo, error)
	Delete(context.Context, string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}
