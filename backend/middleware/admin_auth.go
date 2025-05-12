package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminAuth 管理员认证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户信息
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "未登录",
			})
			c.Abort()
			return
		}

		// 检查用户是否是管理员
		userMap, ok := user.(map[string]interface{})
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "用户信息无效",
			})
			c.Abort()
			return
		}

		isAdmin, ok := userMap["is_admin"].(bool)
		if !ok || !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "需要管理员权限",
			})
			c.Abort()
			return
		}

		c.Next()
	}
} 