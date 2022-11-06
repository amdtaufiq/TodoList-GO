package routes

import (
	"TodoList-GO/app/controllers"
	"TodoList-GO/app/repositories"
	"TodoList-GO/app/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func SubTodoRoute(routerGroup *echo.Group, DB *gorm.DB) {
	subTodoRepository := repositories.SubTodoRepository(DB)
	subTodoService := services.SubTodoService(subTodoRepository)
	subTodoController := controllers.SubTodoController(subTodoService)

	routerGroup.GET("/todo/:id/sub_todo", subTodoController.GetSubTodos)
	routerGroup.GET("/todo/:id/sub_todo/:sub_id", subTodoController.GetSubTodo)
	routerGroup.POST("/todo/:id/sub_todo", subTodoController.CreateSubTodo)
	routerGroup.PUT("/todo/:id/sub_todo/:sub_id", subTodoController.UpdateSubTodo)
	routerGroup.DELETE("/todo/:id/sub_todo/:sub_id", subTodoController.DeleteSubTodo)
}
