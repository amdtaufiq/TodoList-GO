package routes

import (
	"TodoList-GO/app/controllers"
	"TodoList-GO/app/repositories"
	"TodoList-GO/app/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func TodoRoute(routerGroup *echo.Group, DB *gorm.DB) {
	todoRepository := repositories.TodoRepository(DB)
	todoService := services.TodoService(todoRepository)
	todoController := controllers.TodoController(todoService)

	routerGroup.GET("/todo", todoController.GetTodos)
	routerGroup.GET("/todo/detail", todoController.GetDetailTodos)
	routerGroup.GET("/todo/:id", todoController.GetTodo)
	routerGroup.POST("/todo", todoController.CreateTodo)
	routerGroup.PUT("/todo/:id", todoController.UpdateTodo)
	routerGroup.DELETE("/todo/:id", todoController.DeleteTodo)
}
