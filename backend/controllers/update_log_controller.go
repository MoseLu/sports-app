package controllers

import (
	"net/http"
	"strconv"

	"sports-app/backend/services"

	"github.com/gin-gonic/gin"
)

// UpdateLogController 更新日志控制器
type UpdateLogController struct {
	updateLogService *services.UpdateLogService
}

// NewUpdateLogController 创建更新日志控制器实例
func NewUpdateLogController(updateLogService *services.UpdateLogService) *UpdateLogController {
	return &UpdateLogController{
		updateLogService: updateLogService,
	}
}

// GetUpdateLogs 获取更新日志列表
func (c *UpdateLogController) GetUpdateLogs(ctx *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	// 获取日志列表
	logs, total, err := c.updateLogService.GetUpdateLogs(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取更新日志失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": logs,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

// GetUpdateStats 获取更新统计信息
func (c *UpdateLogController) GetUpdateStats(ctx *gin.Context) {
	stats, err := c.updateLogService.GetUpdateStats()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取更新统计信息失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": stats,
	})
} 