package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotoRoutes(router *gin.Engine) {
	photos := router.Group("/photos").Use(middlewares.Authentication())
	{
		photos.POST("", controllers.CreatePhoto)
		photos.GET("", controllers.GetPhotos)
		photos.PUT("/:photoId", controllers.UpdatePhoto)
		photos.DELETE("/:photoId", controllers.DeletePhoto)
	}
}
