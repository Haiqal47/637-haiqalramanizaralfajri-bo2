package routers

import (
	"github.com/gin-gonic/gin"

	"ass-03/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/shuffle", controllers.Shuffle)

	return r
}
