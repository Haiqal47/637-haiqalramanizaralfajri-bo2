package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ass-02/controllers"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders)
	r.GET("/orders/:orderId", controllers.GetOrder)
	r.PUT("/orders/:orderId", controllers.UpdateOrder)
	r.DELETE("/orders/:orderId", controllers.DeleteOrder)

	return r
}
