package routers

import (
	"github.com/gin-gonic/gin"

	"ass-03/controllers"
	"ass-03/middlewares"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	cors := r.Group("/").Use(middlewares.CORSMiddleware())
	{
		cors.GET("/shuffle", controllers.Shuffle)
	}

	return r
}
