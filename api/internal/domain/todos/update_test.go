package todos

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {

	toDo := ToDo{
		ID: "1",
	}

	toDoMatch := mock.MatchedBy(func(todo ToDo) bool {
		return todo.ID == "1"
	})

	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}

		mockRepo.On("Exists", toDo.ID).Return(true, nil)
		mockRepo.On("Update", toDoMatch).Return(nil)

		updatedTodo, err := service.Update(toDo)
		assert.NoError(t, err)
		assert.NotZero(t, updatedTodo.UpdatedAt)
		mockRepo.AssertExpectations(t)
	})

	t.Run("todo not found", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}

		mockRepo.On("Exists", toDo.ID).Return(false, nil)

		_, err := service.Update(toDo)
		assert.ErrorIs(t, err, ErrToDoNotFound)
		mockRepo.AssertExpectations(t)
	})

	t.Run("exists error", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}

		mockRepo.On("Exists", toDo.ID).Return(false, errors.New("exists error"))

		_, err := service.Update(toDo)
		assert.ErrorIs(t, err, ErrUpdateToDo)
		mockRepo.AssertExpectations(t)
	})

	t.Run("update error", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}

		mockRepo.On("Exists", toDo.ID).Return(true, nil)
		mockRepo.On("Update", toDoMatch).Return(errors.New("update error"))

		_, err := service.Update(toDo)
		assert.ErrorIs(t, err, ErrUpdateToDo)
		mockRepo.AssertExpectations(t)
	})
}
