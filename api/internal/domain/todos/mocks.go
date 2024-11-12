package todos

import "github.com/stretchr/testify/mock"

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) FindAll() ([]ToDo, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]ToDo), args.Error(1)
}

func (m *MockRepo) FindByID(id string) (ToDo, error) {
	args := m.Called(id)
	return args.Get(0).(ToDo), args.Error(1)
}

func (m *MockRepo) Exists(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepo) Save(todo ToDo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockRepo) Update(todo ToDo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
