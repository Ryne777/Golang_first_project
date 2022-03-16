package routers

import (
	"github.com/Ryne777/Golang_first_project/internal/config"
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
			c.Status(200)
		})
	}

	return r
}
