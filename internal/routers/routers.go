package routers

import (
	"github.com/Ryne777/Golang_first_project/internal/config"
	"github.com/Ryne777/Golang_first_project/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config, cnt controller.ToDoController, cnt2 controller.AuthorController) *gin.Engine {
	if !*cfg.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	//TODO add routers in this group
	todo := r.Group("/api/todo")
	{
		todo.GET("/", cnt.GetAll)
		todo.DELETE("/:id", cnt.Delete)
		todo.POST("/", cnt.Create)
		todo.GET("/:id", cnt.GetOne)
		todo.PUT("/:id", cnt.Update)

	}
	author := r.Group("/api/author")
	{
		author.GET("/:id", cnt2.GetOne)
		author.POST("/", cnt2.Create)
	}
	return r
}
