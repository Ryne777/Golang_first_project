package repository

import (
	"github.com/Ryne777/Golang_first_project/internal/model"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	Repository
	GetAuthorTodos() (interface{}, error)
}
type grAuthorRepository struct {
	db *gorm.DB
}

// Create implements AuthorRepository
func (t *grAuthorRepository) Create(entity interface{}) error {
	m := entity.(model.Author)
	err := t.db.Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete implements AuthorRepository
func (*grAuthorRepository) Delete(id interface{}) error {
	panic("unimplemented")
}

// GetAll implements AuthorRepository
func (*grAuthorRepository) GetAll() (interface{}, error) {
	panic("unimplemented")
}

// GetOneById implements AuthorRepository
func (t *grAuthorRepository) GetOneById(id interface{}) (interface{}, error) {
	model := model.GetAuthor()
	err := t.db.Find(&model, id).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// Update implements AuthorRepository
func (*grAuthorRepository) Update(obj interface{}) (interface{}, error) {
	panic("unimplemented")
}

// GetAuthorTodos implements AuthorRepository
func (*grAuthorRepository) GetAuthorTodos() (interface{}, error) {
	panic("unimplemented")
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &grAuthorRepository{db: db}
}
