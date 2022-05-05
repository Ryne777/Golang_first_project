package controller

import (
	"net/http"

	"github.com/Ryne777/Golang_first_project/internal/model"
	"github.com/Ryne777/Golang_first_project/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var _ AuthorController = &todoController{}

//TODO : add new abstraction like service
type AuthorController interface {
	GetAll(*gin.Context)
	GetOne(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

type authorController struct {
	repository repository.AuthorRepository
}

func NewAuthorController(db *gorm.DB) ToDoController {
	return &authorController{
		repository: repository.NewAuthorRepository(db),
	}
}

// Create implements ToDoController
func (c *authorController) Create(g *gin.Context) {
	input := model.GetAuthor()
	err := g.ShouldBindJSON(input)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}
	author := model.Author{
		User: model.User{
			Model:     gorm.Model{},
			FirstName: input.FirstName,
			LastName:  input.LastName,
		},
		ToDos: input.ToDos,
	}
	err = c.repository.Create(author)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.Status(http.StatusNoContent)
}

// Delete implements ToDoController
func (*authorController) Delete(g *gin.Context) {
	panic("unimplemented")
}

// GetAll implements ToDoController
func (*authorController) GetAll(*gin.Context) {
	panic("unimplemented")
}

// GetOne implements ToDoController
func (c *authorController) GetOne(g *gin.Context) {
	id := g.Param("id")
	data, err := c.repository.GetOneById(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}
	g.JSON(http.StatusOK, data)
}

// Update implements ToDoController
func (*authorController) Update(g *gin.Context) {
	panic("unimplemented")
}
