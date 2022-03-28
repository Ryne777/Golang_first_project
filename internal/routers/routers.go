package routers

import (
	"github.com/Ryne777/Golang_first_project/internal/config"
	"github.com/Ryne777/Golang_first_project/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config, cnt controller.ToDoController) *gin.Engine {
	if !*cfg.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	//TODO add routers in this group
	api := r.Group("/api")
	{
		api.GET("/", cnt.GetAll)
		api.DELETE("/:id", cnt.Delete)
		api.POST("/", cnt.Create)
		api.GET("/:id", cnt.GetOne)
		api.PUT("/:id", cnt.Update)

	}

	return r
}
