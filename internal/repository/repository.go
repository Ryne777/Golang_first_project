package repository

type Repository interface {
	GetAll() (interface{}, error)
	Create(str interface{}) error
	Update(id uint, str interface{}) (interface{}, error)
	Delete(id uint) error
}
