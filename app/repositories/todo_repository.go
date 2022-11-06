package repositories

import (
	"TodoList-GO/app/entities"

	"gorm.io/gorm"
)

type ITodoRepository interface {
	GetAll() ([]entities.Todo, error)
	GetDetailAll() ([]entities.Todo, error)
	GetByID(ID int) (entities.Todo, error)
	Create(todo entities.Todo) (bool, error)
	Update(todo entities.Todo) (bool, error)
	Delete(ID int) (bool, error)
}

type todoRepository struct {
	db *gorm.DB
}

func TodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) GetAll() ([]entities.Todo, error) {
	var todos []entities.Todo

	err := r.db.Find(&todos).Error

	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (r *todoRepository) GetDetailAll() ([]entities.Todo, error) {
	var todos []entities.Todo

	err := r.db.Preload("SubTodo").Find(&todos).Error

	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (r *todoRepository) GetByID(ID int) (entities.Todo, error) {
	var todo entities.Todo

	err := r.db.Where("id = ?", ID).Preload("SubTodo").Find(&todo).Error

	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *todoRepository) Create(todo entities.Todo) (bool, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *todoRepository) Update(todo entities.Todo) (bool, error) {
	err := r.db.Save(&todo).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *todoRepository) Delete(ID int) (bool, error) {
	err := r.db.Where("id = ?", ID).Delete(&entities.Todo{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
