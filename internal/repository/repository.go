package repository

type Repository interface {
	GetAll() (interface{}, error)
	GetOneById(id interface{}) (interface{}, error)
	Create(str interface{}) error
	Update(obj interface{}) (interface{}, error)
	Delete(id interface{}) error
}
