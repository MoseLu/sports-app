package routes

import (
	"sports-app/backend/controllers"
	"sports-app/backend/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRecordRoutes 注册运动记录相关路由
func RegisterRecordRoutes(r *gin.Engine, recordController *controllers.RecordController) {
	// 需要认证的路由组
	authGroup := r.Group("/api")
	authGroup.Use(middleware.AuthMiddleware())

	// 运动记录路由
	authGroup.GET("/records", recordController.GetRecords)
	authGroup.POST("/records", recordController.CreateRecord)
	authGroup.PUT("/records/:id", recordController.UpdateRecord)
	authGroup.DELETE("/records/:id", recordController.DeleteRecord)

	// 运动类型路由
	authGroup.GET("/sport-types", recordController.GetSportTypes)
	authGroup.POST("/sport-types", recordController.CreateSportType)
	authGroup.PUT("/sport-types/:id", recordController.UpdateSportType)
	authGroup.DELETE("/sport-types/:id", recordController.DeleteSportType)
} 