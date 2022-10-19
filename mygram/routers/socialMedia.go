package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func SocialMediaRoutes(router *gin.Engine) {
	comments := router.Group("/socialmedias").Use(middlewares.Authentication())
	{
		comments.POST("", controllers.CreateSocialMedia)
		comments.GET("", controllers.GetSocialMedias)
		comments.PUT("/:socialMediaId", controllers.UpdateSocialMedia)
		comments.DELETE("/:socialMediaId", controllers.DeleteSocialMedia)
	}
}
