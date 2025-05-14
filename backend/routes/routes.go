package routes

import (
	"net/http"
	"sports-app/backend/controllers"
	"sports-app/backend/middleware"
	"sports-app/backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Manifest struct {
	Version  string `json:"version"`
	BundleURL string `json:"bundleUrl"`
	Sha256   string `json:"sha256"`
}

func manifestHandler(c *gin.Context) {
	m := Manifest{
		Version:  "1.0.0",
		BundleURL: "https://redamancy.com.cn/bundles/1.0.0.zip",
		Sha256:   "b5b0a46288e3...", // 这里需要替换为实际的 SHA256
	}
	c.JSON(http.StatusOK, m)
}

// SetupRoutes 设置所有路由
func SetupRoutes(r *gin.Engine, db *gorm.DB, logsDB *gorm.DB) {
	// 使用中间件
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.ErrorHandler())

	// 创建服务实例
	verificationService := services.NewVerificationService(logsDB)
	authService := services.NewAuthService(db, verificationService)
	recordService := services.NewRecordService(db)
	sportTypeService := services.NewSportTypeService(db)
	updateLogService := services.NewUpdateLogService(logsDB)

	// 创建控制器实例
	authController := controllers.NewAuthController(authService)
	recordController := controllers.NewRecordController(recordService)
	userController := controllers.NewUserController(db)
	sportTypeController := controllers.NewSportTypeController(sportTypeService)
	manifestController := controllers.NewManifestController(updateLogService)
	updateLogController := controllers.NewUpdateLogController(updateLogService)

	// API 路由组
	api := r.Group("/api")
	{
		// 不需要认证的路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
			auth.POST("/logout", authController.Logout)
			auth.POST("/send-reset-code", authController.SendCode)
			auth.POST("/verify-code", authController.VerifyCode)
			auth.POST("/reset-password", authController.ResetPassword)
		}

		// Manifest 相关路由 - 公开访问
		api.GET("/manifest", manifestController.GetManifest)
		api.POST("/manifest", middleware.AdminAuth(), manifestController.UpdateManifest)

		// 需要认证的路由
		authorized := api.Group("")
		authorized.Use(middleware.AuthMiddleware())
		{
			// 用户相关路由
			users := authorized.Group("/users")
			{
				users.GET("/profile", userController.GetProfile)
				users.PUT("/profile", userController.UpdateProfile)
			}

			// 记录相关路由
			records := authorized.Group("/records")
			{
				records.GET("", recordController.GetRecords)
				records.POST("", recordController.CreateRecord)
				records.PUT("/:id", recordController.UpdateRecord)
				records.DELETE("/:id", recordController.DeleteRecord)
				records.GET("/stats", recordController.GetStats)
			}

			// 运动类型相关路由
			sportTypes := authorized.Group("/sport-types")
			{
				sportTypes.GET("", sportTypeController.GetSportTypes)
				sportTypes.POST("", sportTypeController.CreateSportType)
				sportTypes.PUT("/:id", sportTypeController.UpdateSportType)
				sportTypes.DELETE("/:id", sportTypeController.DeleteSportType)
			}

			// 图片上传路由
			upload := authorized.Group("/upload")
			{
				uploadController := controllers.NewUploadController()
				upload.POST("/image", uploadController.UploadImage)
			}

			// 更新日志相关路由 - 需要管理员权限
			api.GET("/update-logs", middleware.AdminAuth(), updateLogController.GetUpdateLogs)
			api.GET("/update-stats", middleware.AdminAuth(), updateLogController.GetUpdateStats)
		}
	}

	// 测试接口 - 不需要任何认证
	r.GET("/api/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "我只是一个测试接口，不要在意我，是帅哥写的！",
			"status": "ok",
			"timestamp": c.GetInt64("timestamp"),
		})
	})

	// Live Update manifest
	r.GET("/api/manifest.json", manifestHandler)

	// TODO: 添加其他功能模块的路由
	// SetupRecordRoutes(api, db)
	// SetupCommunityRoutes(api, db)
}
