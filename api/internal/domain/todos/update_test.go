package todos

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {
	idMatch := mock.MatchedBy(func(id string) bool {
		_, err := uuid.Parse(id)
		return err == nil
	})

	t.Run("should update an existing todo", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := NewService(mockRepo)
		todo := ToDo{
			ID:    uuid.New().String(),
			Title: "Updated ToDo",
		}
		oldToDo := ToDo{
			ID:        todo.ID,
			Title:     "Old ToDo",
			CreatedAt: time.Now().Add(-time.Hour),
		}

		mockRepo.On("Exists", idMatch).Return(true, nil)
		mockRepo.On("FindByID", idMatch).Return(oldToDo, nil)
		mockRepo.On("Update", mock.Anything).Return(nil)

		updatedToDo, err := service.Update(context.Background(), todo)

		assert.NoError(t, err)
		assert.Equal(t, todo.Title, updatedToDo.Title)
		assert.WithinDuration(t, time.Now(), updatedToDo.UpdatedAt, time.Second)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if todo does not exist", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}
		todo := ToDo{
			ID:    uuid.New().String(),
			Title: "Updated ToDo",
		}

		mockRepo.On("Exists", idMatch).Return(false, nil)

		_, err := service.Update(context.Background(), todo)

		assert.ErrorIs(t, err, ErrToDoNotFound)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repo.Exists fails", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}
		todo := ToDo{
			ID:    uuid.New().String(),
			Title: "Updated ToDo",
		}

		mockRepo.On("Exists", idMatch).Return(false, errors.New("exists error"))

		_, err := service.Update(context.Background(), todo)

		assert.ErrorIs(t, err, ErrUpdateToDo)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repo.FindByID fails", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}
		todo := ToDo{
			ID:    uuid.New().String(),
			Title: "Updated ToDo",
		}

		mockRepo.On("Exists", idMatch).Return(true, nil)
		mockRepo.On("FindByID", idMatch).Return(ToDo{}, errors.New("find error"))

		_, err := service.Update(context.Background(), todo)

		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repo.Update fails", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := &service{repo: mockRepo}
		todo := ToDo{
			ID:    uuid.New().String(),
			Title: "Updated ToDo",
		}
		oldToDo := ToDo{
			ID:        todo.ID,
			Title:     "Old ToDo",
			CreatedAt: time.Now().Add(-time.Hour),
		}

		mockRepo.On("Exists", idMatch).Return(true, nil)
		mockRepo.On("FindByID", idMatch).Return(oldToDo, nil)
		mockRepo.On("Update", mock.Anything).Return(errors.New("update error"))

		_, err := service.Update(context.Background(), todo)

		assert.ErrorIs(t, err, ErrUpdateToDo)

		mockRepo.AssertExpectations(t)
	})
}

func TestCompareAndReplace(t *testing.T) {
	t.Run("should replace title if new title is not empty", func(t *testing.T) {
		oldToDo := ToDo{
			Title: "Old Title",
		}
		newToDo := ToDo{
			Title: "New Title",
		}

		updatedToDo := compareAndReplace(oldToDo, newToDo)

		assert.Equal(t, "New Title", updatedToDo.Title)
	})

	t.Run("should not replace title if new title is empty", func(t *testing.T) {
		oldToDo := ToDo{
			Title: "Old Title",
		}
		newToDo := ToDo{
			Title: "",
		}

		updatedToDo := compareAndReplace(oldToDo, newToDo)

		assert.Equal(t, "Old Title", updatedToDo.Title)
	})

	t.Run("should replace completed status if different", func(t *testing.T) {
		oldToDo := ToDo{
			Completed: false,
		}
		newToDo := ToDo{
			Completed: true,
		}

		updatedToDo := compareAndReplace(oldToDo, newToDo)

		assert.Equal(t, true, updatedToDo.Completed)
	})

	t.Run("should not replace completed status if same", func(t *testing.T) {
		oldToDo := ToDo{
			Completed: false,
		}
		newToDo := ToDo{
			Completed: false,
		}

		updatedToDo := compareAndReplace(oldToDo, newToDo)

		assert.Equal(t, false, updatedToDo.Completed)
	})

	t.Run("should update UpdatedAt field", func(t *testing.T) {
		oldToDo := ToDo{
			UpdatedAt: time.Now().Add(-time.Hour),
		}
		newToDo := ToDo{}

		updatedToDo := compareAndReplace(oldToDo, newToDo)

		assert.WithinDuration(t, time.Now(), updatedToDo.UpdatedAt, time.Second)
	})
}
