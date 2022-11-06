package inputs

import "TodoList-GO/app/entities"

type GetIDSubTodoInput struct {
	SubID string
}

type CreateSubTodoInput struct {
	Title        string `form:"title" binding:"required"`
	Description  string `form:"description" binding:"required"`
	TodoID       int
	FileLocation string
}

type UpdateSubTodoInput struct {
	ID           int    `form:"id" binding:"required"`
	Title        string `form:"title" binding:"required"`
	Description  string `form:"description" binding:"required"`
	FileLocation string
}

func FormatCreateSubTodo(input CreateSubTodoInput) entities.SubTodo {
	todo := entities.SubTodo{}
	todo.Title = input.Title
	todo.Description = input.Description
	todo.File = input.FileLocation
	todo.TodoID = input.TodoID

	return todo
}

func FormatUpdateSubTodo(todo entities.SubTodo, input UpdateSubTodoInput) entities.SubTodo {
	todo.Title = input.Title
	todo.Description = input.Description
	todo.File = input.FileLocation

	return todo
}
