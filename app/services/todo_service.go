package services

import (
	"TodoList-GO/app/entities"
	"TodoList-GO/app/inputs"
	"TodoList-GO/app/repositories"
	"errors"
	"strconv"
)

type ITodoService interface {
	GetAll() ([]entities.Todo, error)
	GetDetailAll() ([]entities.Todo, error)
	GetByID(request inputs.GetIDTodoInput) (entities.Todo, error)
	Create(request inputs.CreateTodoInput) (bool, error)
	Update(requestID inputs.GetIDTodoInput, requestData inputs.UpdateTodoInput) (bool, error)
	Delete(request inputs.GetIDTodoInput) (bool, error)
}

type todoService struct {
	todoRepository repositories.ITodoRepository
}

func TodoService(repository repositories.ITodoRepository) *todoService {
	return &todoService{repository}
}

func (s *todoService) GetAll() ([]entities.Todo, error) {
	todos, err := s.todoRepository.GetAll()
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (s *todoService) GetDetailAll() ([]entities.Todo, error) {
	todos, err := s.todoRepository.GetDetailAll()
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (s *todoService) GetByID(request inputs.GetIDTodoInput) (entities.Todo, error) {
	todo := entities.Todo{}
	todoID, _ := strconv.Atoi(request.ID)
	todo, err := s.todoRepository.GetByID(todoID)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s *todoService) Create(request inputs.CreateTodoInput) (bool, error) {
	todoFormat := inputs.FormatCreateTodo(request)

	result, err := s.todoRepository.Create(todoFormat)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *todoService) Update(requestID inputs.GetIDTodoInput, requestData inputs.UpdateTodoInput) (bool, error) {
	todoID, _ := strconv.Atoi(requestID.ID)
	todo, err := s.todoRepository.GetByID(todoID)
	if err != nil {
		return false, err
	}

	if todo.ID != requestData.ID {
		return false, errors.New("todo data not match")
	}

	todoFormat := inputs.FormatUpdateTodo(todo, requestData)

	result, err := s.todoRepository.Update(todoFormat)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *todoService) Delete(request inputs.GetIDTodoInput) (bool, error) {
	todoID, _ := strconv.Atoi(request.ID)
	_, err := s.todoRepository.GetByID(todoID)
	if err != nil {
		return false, errors.New("entities.Todo data not available")
	}

	result, err := s.todoRepository.Delete(todoID)
	if err != nil {
		return result, err
	}
	return result, nil
}
