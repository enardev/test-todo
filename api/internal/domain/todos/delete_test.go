package todos

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	t.Run("SuccessfulDelete", func(t *testing.T) {
		mockRepo := new(MockRepo)
		s := &service{repo: mockRepo}

		mockRepo.On("Exists", "1").Return(true, nil)
		mockRepo.On("Delete", "1").Return(nil)

		err := s.Delete("1")
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("TodoNotFound", func(t *testing.T) {
		mockRepo := new(MockRepo)
		s := &service{repo: mockRepo}

		mockRepo.On("Exists", "2").Return(false, nil)

		err := s.Delete("2")
		assert.Equal(t, ErrToDoNotFound, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("ExistsError", func(t *testing.T) {
		mockRepo := new(MockRepo)
		s := &service{repo: mockRepo}

		mockRepo.On("Exists", "3").Return(false, errors.New("exists error"))

		err := s.Delete("3")
		assert.Equal(t, ErrDeleteToDo, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteError", func(t *testing.T) {
		mockRepo := new(MockRepo)
		s := &service{repo: mockRepo}

		mockRepo.On("Exists", "4").Return(true, nil)
		mockRepo.On("Delete", "4").Return(errors.New("delete error"))

		err := s.Delete("4")
		assert.Equal(t, ErrDeleteToDo, err)

		mockRepo.AssertExpectations(t)
	})
}
