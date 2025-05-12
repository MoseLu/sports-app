package controllers

import (
	"net/http"
	"sports-app/backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ErrorLogController 错误日志控制器
type ErrorLogController struct {
	db *gorm.DB
}

// NewErrorLogController 创建错误日志控制器
func NewErrorLogController(db *gorm.DB) *ErrorLogController {
	return &ErrorLogController{db: db}
}

// CreateErrorLog 创建错误日志
func (c *ErrorLogController) CreateErrorLog(ctx *gin.Context) {
	var errorLog models.ErrorLog
	if err := ctx.ShouldBindJSON(&errorLog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "无效的请求数据",
			"error": err.Error(),
		})
		return
	}

	// 设置默认值
	if errorLog.ErrorType == "" {
		errorLog.ErrorType = "unknown"
	}
	if errorLog.Message == "" {
		errorLog.Message = "未知错误"
	}
	if errorLog.StatusCode == 0 {
		errorLog.StatusCode = 500
	}

	if err := c.db.Create(&errorLog).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "创建错误日志失败",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, errorLog)
} 