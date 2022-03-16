package repository

import (
	"github.com/Ryne777/Golang_first_project/internal/model"
	"gorm.io/gorm"
)

type ToDoRepository interface {
	Repository
}

type grTodoRepository struct {
	db *gorm.DB
}

func NewToDoRepository(db *gorm.DB) ToDoRepository {
	return grTodoRepository{db: db}
}

func (t grTodoRepository) GetAll() (interface{}, error) {
	var model model.ToDo
	err := t.db.Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (t grTodoRepository) Create(entity interface{}) error {
	err := t.db.Create(entity.(model.ToDo)).Error
	if err != nil {
		return err
	}
	return nil
}

func (t grTodoRepository) Update(id uint, str interface{}) (interface{}, error) {
	var model model.ToDo
	err := t.db.Find(&model, id).Error
	if err != nil {
		return nil, err
	}
	err = t.db.Save(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
func (t grTodoRepository) Delete(id uint) error {
	var model model.ToDo
	err := t.db.Find(&model, id).Error
	if err != nil {
		return err
	}
	err = t.db.Delete(&model).Error
	if err != nil {
		return err
	}
	return nil
}
