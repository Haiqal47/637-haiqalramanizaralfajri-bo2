package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine) {
	comments := router.Group("/comments").Use(middlewares.Authentication())
	{
		comments.POST("", controllers.CreateComment)
		comments.GET("", controllers.GetComments)
		comments.PUT("/:commentId", controllers.UpdateComment)
		comments.DELETE("/:commentId", controllers.DeleteComment)
	}
}
