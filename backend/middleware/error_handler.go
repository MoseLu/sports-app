package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 全局错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否有错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			statusCode := http.StatusInternalServerError

			// 根据错误类型设置状态码
			switch err.(type) {
			case *gin.Error:
				if e, ok := err.(*gin.Error); ok {
					if e.Type == gin.ErrorTypeBind {
						statusCode = http.StatusBadRequest
					}
				}
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
				"message": "服务器内部错误",
				"status": statusCode,
			})
		}
	}
} 