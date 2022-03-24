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
	return &grTodoRepository{db: db}
}

func (t grTodoRepository) GetAll() (interface{}, error) {
	models := model.GetTodos()
	err := t.db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (t grTodoRepository) Create(entity interface{}) error {
	m := entity.(model.ToDo)
	err := t.db.Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}

func (t grTodoRepository) Update(id uint, str interface{}) (interface{}, error) {
	model := model.GetTodo()
	err := t.db.Find(model, id).Error
	if err != nil {
		return nil, err
	}
	err = t.db.Save(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
func (t grTodoRepository) Delete(id uint) error {
	model := model.GetTodo()
	err := t.db.Delete(model, id).Error
	if err != nil {
		return err
	}
	return nil
}
