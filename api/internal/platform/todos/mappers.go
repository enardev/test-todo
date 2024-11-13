package todos

import "test-todo/api/internal/domain/todos"

func mapToDomain(todo ToDo) todos.ToDo {
	return todos.ToDo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func mapToEntity(todo todos.ToDo) ToDo {
	return ToDo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func mapToDomainList(todoList []ToDo) []todos.ToDo {
	var list []todos.ToDo
	for _, todo := range todoList {
		list = append(list, mapToDomain(todo))
	}
	return list
}
