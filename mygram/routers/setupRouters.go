package routers

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	r.Use(location.Default())

	r.Static("/content", "./uploaded")

	UserRoutes(r)
	PhotoRoutes(r)
	CommentRoutes(r)
	SocialMediaRoutes(r)

	return r
}
