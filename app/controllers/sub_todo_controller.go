package controllers

import (
	"TodoList-GO/app/formatters"
	"TodoList-GO/app/helper"
	"TodoList-GO/app/inputs"
	"TodoList-GO/app/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type subTodoController struct {
	service services.ISubTodoService
}

func SubTodoController(service services.ISubTodoService) *subTodoController {
	return &subTodoController{service}
}

func (h *subTodoController) GetSubTodos(c echo.Context) error {
	var request inputs.GetIDTodoInput
	request.ID = c.Param("id")
	result, err := h.service.GetAll(request)
	if err != nil {
		response := helper.APIResponse("Error to get subTodos", http.StatusBadRequest, false, err)
		return c.JSON(http.StatusBadRequest, response)

	}

	response := helper.APIResponse("List of subTodos", http.StatusOK, true, formatters.FormatSubTodos(result))
	return c.JSON(http.StatusOK, response)
}

func (h *subTodoController) GetSubTodo(c echo.Context) error {
	var request inputs.GetIDSubTodoInput
	request.SubID = c.Param("sub_id")

	result, err := h.service.GetByID(request)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of subTodo", http.StatusBadRequest, false, err)
		return c.JSON(http.StatusBadRequest, response)
	}

	if result.ID == 0 {
		response := helper.APIResponse("Sub Todo data", http.StatusNotFound, true, nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := helper.APIResponse("Sub Todo Data", http.StatusOK, true, formatters.FormatSubTodo(result))
	return c.JSON(http.StatusOK, response)
}

func (h *subTodoController) CreateSubTodo(c echo.Context) error {
	var request inputs.CreateSubTodoInput

	err := c.Bind(&request)

	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to create subTodo", http.StatusBadRequest, false, errors)
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

	request.TodoID, _ = strconv.Atoi(c.Param("id"))

	result, err := h.service.Create(request)
	if err != nil {
		response := helper.APIResponse("Failed to create subTodo", http.StatusBadRequest, false, err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success to create subTodo", http.StatusOK, true, result)
	return c.JSON(http.StatusOK, response)
}

func (h *subTodoController) UpdateSubTodo(c echo.Context) error {
	var inputID inputs.GetIDSubTodoInput
	inputID.SubID = c.Param("sub_id")

	var inputData inputs.UpdateSubTodoInput

	err := c.Bind(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to update subTodo", http.StatusBadRequest, false, errors)
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
		response := helper.APIResponse("Failed to update subTodo", http.StatusBadRequest, false, err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success to update subTodo", http.StatusOK, true, result)
	return c.JSON(http.StatusOK, response)
}

func (h *subTodoController) DeleteSubTodo(c echo.Context) error {
	var request inputs.GetIDSubTodoInput
	request.SubID = c.Param("sub_id")

	result, err := h.service.Delete(request)
	if err != nil {
		response := helper.APIResponse("Failed to delete of subTodo", http.StatusBadRequest, false, result)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Todo deleted", http.StatusOK, true, result)
	return c.JSON(http.StatusOK, response)
}
