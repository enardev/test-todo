package todos

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}
		expectedTodos := []ToDo{
			{ID: "b0d94486-50a8-44d2-8384-14057c52f784", Title: "Test ToDo 1"},
			{ID: "fac8324f-a7dc-4168-a1e2-9529656f0e82", Title: "Test ToDo 2"},
		}
		mockRepo.On("FindAll").Return(expectedTodos, nil)

		todos, err := service.GetAll(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedTodos, todos)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}
		mockRepo.On("FindAll").Return(nil, errors.New("some error"))

		todos, err := service.GetAll(context.Background())

		assert.Error(t, err)
		assert.Empty(t, todos)
		assert.ErrorIs(t, err, ErrToDoNotFound)
		mockRepo.AssertExpectations(t)
	})
}
