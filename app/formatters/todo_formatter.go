package formatters

import (
	"TodoList-GO/app/entities"
	"time"
)

type TodoFormatter struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	File        string             `json:"file_path"`
	CreatedAt   time.Time          `json:"created_at"`
	SubTodo     []SubTodoFormatter `json:"sub_todos"`
}

func FormatTodo(todo entities.Todo) *TodoFormatter {
	todoFormatter := TodoFormatter{}
	todoFormatter.ID = todo.ID
	todoFormatter.Title = todo.Title
	todoFormatter.Description = todo.Description
	todoFormatter.File = todo.File
	todoFormatter.CreatedAt = todo.CreatedAt
	todoFormatter.SubTodo = FormatSubTodos(todo.SubTodo)

	return &todoFormatter
}

func FormatTodos(todos []entities.Todo) []TodoFormatter {
	if len(todos) == 0 {
		return []TodoFormatter{}
	}

	todosFormatter := []TodoFormatter{}

	for _, todo := range todos {
		todoFormatter := FormatTodo(todo)
		todosFormatter = append(todosFormatter, *todoFormatter)
	}

	return todosFormatter
}
