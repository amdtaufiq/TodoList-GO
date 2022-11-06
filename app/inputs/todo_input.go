package inputs

import "TodoList-GO/app/entities"

type GetIDTodoInput struct {
	ID string
}

type CreateTodoInput struct {
	Title        string `form:"title" binding:"required"`
	Description  string `form:"description" binding:"required"`
	FileLocation string
}

type UpdateTodoInput struct {
	ID           int    `form:"id" binding:"required"`
	Title        string `form:"title" binding:"required"`
	Description  string `form:"description" binding:"required"`
	FileLocation string
}

func FormatCreateTodo(input CreateTodoInput) entities.Todo {
	todo := entities.Todo{}
	todo.Title = input.Title
	todo.Description = input.Description
	todo.File = input.FileLocation

	return todo
}

func FormatUpdateTodo(todo entities.Todo, input UpdateTodoInput) entities.Todo {
	todo.Title = input.Title
	todo.Description = input.Description
	if input.FileLocation != "" {
		todo.File = input.FileLocation
	}

	return todo
}
