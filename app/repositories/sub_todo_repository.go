package repositories

import (
	"TodoList-GO/app/entities"

	"gorm.io/gorm"
)

type ISubTodoRepository interface {
	GetAll(ID int) ([]entities.SubTodo, error)
	GetByID(ID int) (entities.SubTodo, error)
	Create(subTodo entities.SubTodo) (bool, error)
	Update(subTodo entities.SubTodo) (bool, error)
	Delete(ID int) (bool, error)
}

type subTodoRepository struct {
	db *gorm.DB
}

func SubTodoRepository(db *gorm.DB) *subTodoRepository {
	return &subTodoRepository{db}
}

func (r *subTodoRepository) GetAll(ID int) ([]entities.SubTodo, error) {
	var subTodos []entities.SubTodo

	err := r.db.Where("todo_id = ?", ID).Find(&subTodos).Error

	if err != nil {
		return subTodos, err
	}
	return subTodos, nil
}

func (r *subTodoRepository) GetByID(ID int) (entities.SubTodo, error) {
	var subTodo entities.SubTodo

	err := r.db.Where("id = ?", ID).Find(&subTodo).Error

	if err != nil {
		return subTodo, err
	}
	return subTodo, nil
}

func (r *subTodoRepository) Create(subTodo entities.SubTodo) (bool, error) {
	err := r.db.Create(&subTodo).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *subTodoRepository) Update(subTodo entities.SubTodo) (bool, error) {
	err := r.db.Save(&subTodo).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *subTodoRepository) Delete(ID int) (bool, error) {
	err := r.db.Where("id = ?", ID).Delete(&entities.SubTodo{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
