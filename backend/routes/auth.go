package routes

import (
	"sports-app/backend/controllers"
	"sports-app/backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupAuthRoutes 设置认证相关路由
func SetupAuthRoutes(r *gin.Engine, db *gorm.DB) {
	verificationService := services.NewVerificationService(db)
	authService := services.NewAuthService(db, verificationService)
	authController := controllers.NewAuthController(authService)

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/logout", authController.Logout)
		auth.POST("/send-reset-code", authController.SendCode)
		auth.POST("/verify-code", authController.VerifyCode)
		auth.POST("/reset-password", authController.ResetPassword)
	}
}
