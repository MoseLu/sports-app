package middleware

import (
	"log"
	"net/http"
	"sports-app/backend/services"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("处理请求: %s %s", ctx.Request.Method, ctx.Request.URL.Path)
		
		// 获取Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("未提供认证信息")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "未提供认证信息",
				"code": "AUTH_REQUIRED",
			})
			return
		}

		// 检查Bearer token格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("无效的认证信息格式")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "无效的认证信息",
				"code": "INVALID_TOKEN_FORMAT",
			})
			return
		}

		// 解析token
		claims, err := services.ValidateToken(parts[1])
		if err != nil {
			log.Printf("token验证失败: %v", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "无效的token",
				"code": "INVALID_TOKEN",
				"details": err.Error(),
			})
			return
		}

		log.Printf("用户 %d 认证成功", claims.UserID)
		// 将用户ID存储到上下文中
		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
