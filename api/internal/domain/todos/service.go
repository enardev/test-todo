package todos

type Repository interface {
	FindAll() ([]ToDo, error)
	Exists(string) (bool, error)
	Save(ToDo) error
	Update(ToDo) error
	Delete(string) error
}

type Service interface {
	GetAll() ([]ToDo, error)
	Create(ToDo) (ToDo, error)
	Update(ToDo) (ToDo, error)
	Delete(string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}
