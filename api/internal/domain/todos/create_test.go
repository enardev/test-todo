package todos

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	idMatch := mock.MatchedBy(func(id string) bool {
		_, err := uuid.Parse(id)
		return err == nil
	})

	t.Run("should create a new todo", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}
		todo := ToDo{
			Title: "Test ToDo",
		}

		mockRepo.On("Exists", idMatch).Return(false, nil)
		mockRepo.On("Save", mock.Anything).Return(nil)

		createdTodo, err := service.Create(todo)

		assert.NoError(t, err)
		assert.NotEmpty(t, createdTodo.ID)
		assert.Equal(t, todo.Title, createdTodo.Title)
		assert.WithinDuration(t, time.Now(), createdTodo.CreatedAt, time.Second)
		assert.WithinDuration(t, time.Now(), createdTodo.UpdatedAt, time.Second)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if todo already exists", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}
		todo := ToDo{
			ID:    uuid.New().String(),
			Title: "Test ToDo",
		}

		mockRepo.On("Exists", idMatch).Return(true, nil)

		_, err := service.Create(todo)

		assert.ErrorIs(t, err, ErrToDoAlreadyExists)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repo.Exists fails", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}
		todo := ToDo{
			ID:    uuid.New().String(),
			Title: "Test ToDo",
		}

		mockRepo.On("Exists", idMatch).Return(false, errors.New("exists error"))

		_, err := service.Create(todo)

		assert.ErrorIs(t, err, ErrSaveToDo)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repo.Save fails", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}
		todo := ToDo{
			Title: "Test ToDo",
		}

		mockRepo.On("Exists", idMatch).Return(false, nil)
		mockRepo.On("Save", mock.Anything).Return(errors.New("save error"))

		_, err := service.Create(todo)

		assert.ErrorIs(t, err, ErrSaveToDo)

		mockRepo.AssertExpectations(t)
	})
}
