package formatters

import (
	"TodoList-GO/app/entities"
	"time"
)

type SubTodoFormatter struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	File        string    `json:"file_path"`
	CreatedAt   time.Time `json:"created_at"`
}

func FormatSubTodo(todo entities.SubTodo) *SubTodoFormatter {
	todoFormatter := SubTodoFormatter{}
	todoFormatter.ID = todo.ID
	todoFormatter.Title = todo.Title
	todoFormatter.Description = todo.Description
	todoFormatter.File = todo.File
	todoFormatter.CreatedAt = todo.CreatedAt

	return &todoFormatter
}

func FormatSubTodos(todos []entities.SubTodo) []SubTodoFormatter {
	if len(todos) == 0 {
		return []SubTodoFormatter{}
	}

	todosFormatter := []SubTodoFormatter{}

	for _, todo := range todos {
		todoFormatter := FormatSubTodo(todo)
		todosFormatter = append(todosFormatter, *todoFormatter)
	}

	return todosFormatter
}
