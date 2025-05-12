package routes

import (
	"sports-app/backend/controllers"
	"sports-app/backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUploadRoutes(router *gin.Engine) {
	uploadController := controllers.NewUploadController()

	// 图片上传路由组
	uploadGroup := router.Group("/api/upload")
	uploadGroup.Use(middleware.AuthMiddleware())
	{
		uploadGroup.POST("/image", uploadController.UploadImage)
	}
} 