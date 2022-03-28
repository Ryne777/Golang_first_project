package controller

import (
	"net/http"

	"github.com/Ryne777/Golang_first_project/internal/model"
	"github.com/Ryne777/Golang_first_project/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var _ ToDoController = &controller{}

//TODO : add new abstraction like service
type ToDoController interface {
	GetAll(*gin.Context)
	GetOne(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

type controller struct {
	repository repository.ToDoRepository
}

func NewToDoController(db *gorm.DB) ToDoController {
	return &controller{
		repository: repository.NewToDoRepository(db),
	}
}

// Create implements ToDoController
func (c *controller) Create(g *gin.Context) {
	input := model.GetTodo()
	err := g.ShouldBindJSON(input)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	todo := model.ToDo{
		Model:    gorm.Model{},
		ToDo:     input.ToDo,
		Finished: input.Finished,
		Failed:   input.Failed,
	}
	err = c.repository.Create(todo)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.Status(http.StatusNoContent)

}

// Delete implements ToDoController
func (c *controller) Delete(g *gin.Context) {
	id := g.Param("id")
	err := c.repository.Delete(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.Status(http.StatusNoContent)
}

// GetAll implements ToDoController
func (c *controller) GetAll(g *gin.Context) {
	data, err := c.repository.GetAll()
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, data)
}

// GetOne implements ToDoController
func (c *controller) GetOne(g *gin.Context) {
	id := g.Param("id")
	data, err := c.repository.GetOneById(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, data)

}

// Update implements ToDoController
func (c *controller) Update(g *gin.Context) {
	id := g.Param("id")
	input := model.GetTodo()
	err := g.ShouldBindJSON(input)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	data, err := c.repository.GetOneById(id)

	todo := data.(*model.ToDo)
	todo.ToDo = input.ToDo
	todo.Failed = input.Failed
	todo.Finished = input.Finished
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	todo.ToDo = input.ToDo
	data, err = c.repository.Update(*todo)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, data)
}
