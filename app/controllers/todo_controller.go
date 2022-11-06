package controllers

import (
	"TodoList-GO/app/formatters"
	"TodoList-GO/app/helper"
	"TodoList-GO/app/inputs"
	"TodoList-GO/app/services"
	"net/http"

	"github.com/labstack/echo"
)

type todoController struct {
	service services.ITodoService
}

func TodoController(service services.ITodoService) *todoController {
	return &todoController{service}
}

func (h *todoController) GetTodos(c echo.Context) error {
	result, err := h.service.GetAll()
	if err != nil {
		response := helper.APIResponse("Error to get todos", http.StatusBadRequest, false, err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("List of todos", http.StatusOK, true, formatters.FormatTodos(result))
	return c.JSON(http.StatusOK, response)
}

func (h *todoController) GetDetailTodos(c echo.Context) error {
	result, err := h.service.GetDetailAll()
	if err != nil {
		response := helper.APIResponse("Error to get todos", http.StatusBadRequest, false, err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("List of todos", http.StatusOK, true, formatters.FormatTodos(result))
	return c.JSON(http.StatusOK, response)
}

func (h *todoController) GetTodo(c echo.Context) error {
	var request inputs.GetIDTodoInput
	request.ID = c.Param("id")

	result, err := h.service.GetByID(request)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of todo", http.StatusBadRequest, false, err)
		return c.JSON(http.StatusBadRequest, response)

	}

	if result.ID == 0 {
		response := helper.APIResponse("Todo data", http.StatusNotFound, true, nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := helper.APIResponse("Todo data", http.StatusOK, true, formatters.FormatTodo(result))
	return c.JSON(http.StatusOK, response)
}

func (h *todoController) CreateTodo(c echo.Context) error {
	var request inputs.CreateTodoInput

	err := c.Bind(&request)

	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to create todo", http.StatusBadRequest, false, errors)
		return c.JSON(http.StatusBadRequest, response)

	}

	file, _ := c.FormFile("file")
	if file != nil {
		request.FileLocation, err = helper.UploadFile(file)
		if err != nil {
			response := helper.APIResponse("Failed to upload file", http.StatusBadRequest, false, err)
			return c.JSON(http.StatusBadRequest, response)

		}
	}

	result, err := h.service.Create(request)
	if err != nil {
		response := helper.APIResponse("Failed to create todo", http.StatusBadRequest, false, err)
		return c.JSON(http.StatusBadRequest, response)

	}

	response := helper.APIResponse("Success to create todo", http.StatusOK, true, result)
	return c.JSON(http.StatusOK, response)
}

func (h *todoController) UpdateTodo(c echo.Context) error {
	var inputID inputs.GetIDTodoInput
	inputID.ID = c.Param("id")

	var inputData inputs.UpdateTodoInput

	err := c.Bind(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to update todo", http.StatusBadRequest, false, errors)
		return c.JSON(http.StatusBadRequest, response)

	}

	file, _ := c.FormFile("file")
	if file != nil {
		inputData.FileLocation, err = helper.UploadFile(file)
		if err != nil {
			response := helper.APIResponse("Failed to upload file", http.StatusBadRequest, false, err)
			return c.JSON(http.StatusBadRequest, response)

		}
	}

	result, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update todo", http.StatusBadRequest, false, err)
		return c.JSON(http.StatusBadRequest, response)

	}

	response := helper.APIResponse("Success to update todo", http.StatusOK, true, result)
	return c.JSON(http.StatusOK, response)
}

func (h *todoController) DeleteTodo(c echo.Context) error {
	var request inputs.GetIDTodoInput
	request.ID = c.Param("id")

	result, err := h.service.Delete(request)
	if err != nil {
		response := helper.APIResponse("Failed to delete of todo", http.StatusBadRequest, false, result)
		return c.JSON(http.StatusBadRequest, response)

	}

	response := helper.APIResponse("Todo deleted", http.StatusOK, true, result)
	return c.JSON(http.StatusOK, response)
}
