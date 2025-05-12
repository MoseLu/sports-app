package controllers

import (
	"net/http"
	"sports-app/backend/models"
	"sports-app/backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RecordController 运动记录控制器
type RecordController struct {
	service *services.RecordService
}

// NewRecordController 创建运动记录控制器实例
func NewRecordController(service *services.RecordService) *RecordController {
	return &RecordController{service: service}
}

// GetRecords 获取用户的运动记录列表
func (c *RecordController) GetRecords(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	records, err := c.service.GetRecords(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, records)
}

// CreateRecord 创建运动记录
func (c *RecordController) CreateRecord(ctx *gin.Context) {
	var record models.SportRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.UserID = ctx.GetInt64("user_id")
	if err := c.service.CreateRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, record)
}

// UpdateRecord 更新运动记录
func (c *RecordController) UpdateRecord(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	userID := ctx.GetInt64("user_id")

	var record models.SportRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.ID = id
	record.UserID = userID
	if err := c.service.UpdateRecord(&record); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// DeleteRecord 删除运动记录
func (c *RecordController) DeleteRecord(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	userID := ctx.GetInt64("user_id")

	if err := c.service.DeleteRecord(id, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "记录已删除"})
}

// GetSportTypes 获取所有运动类型
func (c *RecordController) GetSportTypes(ctx *gin.Context) {
	types, err := c.service.GetSportTypes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, types)
}

// CreateSportType 创建运动类型
func (c *RecordController) CreateSportType(ctx *gin.Context) {
	var sportType models.SportType
	if err := ctx.ShouldBindJSON(&sportType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateSportType(&sportType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, sportType)
}

// UpdateSportType 更新运动类型
func (c *RecordController) UpdateSportType(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sport type ID"})
		return
	}

	var sportType models.SportType
	if err := ctx.ShouldBindJSON(&sportType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sportType.ID = id
	if err := c.service.UpdateSportType(&sportType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sportType)
}

// DeleteSportType 删除运动类型
func (c *RecordController) DeleteSportType(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sport type ID"})
		return
	}

	if err := c.service.DeleteSportType(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetStats 获取用户的运动统计信息
func (c *RecordController) GetStats(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	
	// 获取查询参数
	timeRange := ctx.DefaultQuery("time_range", "week")
	sportTypeID, _ := strconv.ParseInt(ctx.DefaultQuery("sport_type_id", "0"), 10, 64)
	
	stats, err := c.service.GetStats(userID, timeRange, sportTypeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, stats)
} 