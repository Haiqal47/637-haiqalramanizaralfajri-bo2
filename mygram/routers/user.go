package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
		secured := users.Use(middlewares.Authentication())
		{
			secured.PUT("/:userId", controllers.UpdateUser)
			secured.DELETE("/:userId", controllers.DeleteUser)
		}
	}
}
