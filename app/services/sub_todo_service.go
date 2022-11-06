package services

import (
	"TodoList-GO/app/entities"
	"TodoList-GO/app/inputs"
	"TodoList-GO/app/repositories"
	"errors"
	"strconv"
)

type ISubTodoService interface {
	GetAll(request inputs.GetIDTodoInput) ([]entities.SubTodo, error)
	GetByID(request inputs.GetIDSubTodoInput) (entities.SubTodo, error)
	Create(request inputs.CreateSubTodoInput) (bool, error)
	Update(requestID inputs.GetIDSubTodoInput, requestData inputs.UpdateSubTodoInput) (bool, error)
	Delete(request inputs.GetIDSubTodoInput) (bool, error)
}

type subTodoService struct {
	subTodoRepository repositories.ISubTodoRepository
}

func SubTodoService(repository repositories.ISubTodoRepository) *subTodoService {
	return &subTodoService{repository}
}

func (s *subTodoService) GetAll(request inputs.GetIDTodoInput) ([]entities.SubTodo, error) {
	todoID, _ := strconv.Atoi(request.ID)
	subTodos, err := s.subTodoRepository.GetAll(todoID)
	if err != nil {
		return subTodos, err
	}
	return subTodos, nil
}

func (s *subTodoService) GetByID(request inputs.GetIDSubTodoInput) (entities.SubTodo, error) {
	subTodo := entities.SubTodo{}
	subID, _ := strconv.Atoi(request.SubID)
	subTodo, err := s.subTodoRepository.GetByID(subID)
	if err != nil {
		return subTodo, err
	}
	return subTodo, nil
}

func (s *subTodoService) Create(request inputs.CreateSubTodoInput) (bool, error) {
	subTodoFormat := inputs.FormatCreateSubTodo(request)

	result, err := s.subTodoRepository.Create(subTodoFormat)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *subTodoService) Update(requestID inputs.GetIDSubTodoInput, requestData inputs.UpdateSubTodoInput) (bool, error) {
	subID, _ := strconv.Atoi(requestID.SubID)
	subTodo, err := s.subTodoRepository.GetByID(subID)
	if err != nil {
		return false, err
	}

	if subTodo.ID != requestData.ID {
		return false, errors.New("subTodo data not match")
	}

	subTodoFormat := inputs.FormatUpdateSubTodo(subTodo, requestData)

	result, err := s.subTodoRepository.Update(subTodoFormat)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *subTodoService) Delete(request inputs.GetIDSubTodoInput) (bool, error) {
	subID, _ := strconv.Atoi(request.SubID)
	_, err := s.subTodoRepository.GetByID(subID)
	if err != nil {
		return false, errors.New("entities.SubTodo data not available")
	}

	result, err := s.subTodoRepository.Delete(subID)
	if err != nil {
		return result, err
	}
	return result, nil
}
