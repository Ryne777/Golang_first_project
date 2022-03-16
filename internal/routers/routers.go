package routers

import (
	"github.com/Ryne777/Golang_first_project/internal/config"
	"github.com/Ryne777/Golang_first_project/internal/model"
	"github.com/Ryne777/Golang_first_project/internal/repository"
	"github.com/Ryne777/Golang_first_project/pkg/bdClient"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	if !*cfg.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	//TODO add routers in this group
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			db := bdClient.GetConnectionDb()
			model := model.ToDo{ToDo: "hello", Finished: false, Failed: true}
			db.AutoMigrate(&model)
			db.Create(&model)
			rep := repository.NewToDoRepository(db)
			models, err := rep.GetAll()
			if err != nil {
				c.JSON(200, gin.H{"err": err})
			}
			c.JSON(200, &models)

		})
	}

	return r
}
