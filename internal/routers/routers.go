package routers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
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
