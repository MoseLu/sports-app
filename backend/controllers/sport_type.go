package controllers

import (
	"log"
	"net/http"
	"sports-app/backend/models"
	"sports-app/backend/services"

	"github.com/gin-gonic/gin"
)

type SportTypeController struct {
	sportTypeService *services.SportTypeService
}

func NewSportTypeController(sportTypeService *services.SportTypeService) *SportTypeController {
	return &SportTypeController{
		sportTypeService: sportTypeService,
	}
}

// GetSportTypes 获取所有运动类型
func (c *SportTypeController) GetSportTypes(ctx *gin.Context) {
	log.Println("收到获取运动类型请求")
	
	sportTypes, err := c.sportTypeService.GetSportTypes()
	if err != nil {
		log.Printf("获取运动类型失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("成功返回 %d 个运动类型", len(sportTypes))
	ctx.JSON(http.StatusOK, sportTypes)
}

// CreateSportType 创建运动类型
func (c *SportTypeController) CreateSportType(ctx *gin.Context) {
	var sportType models.SportType
	if err := ctx.ShouldBindJSON(&sportType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.sportTypeService.CreateSportType(&sportType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, sportType)
}

// UpdateSportType 更新运动类型
func (c *SportTypeController) UpdateSportType(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var sportType models.SportType
	if err := ctx.ShouldBindJSON(&sportType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.sportTypeService.UpdateSportType(idStr, &sportType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sportType)
}

// DeleteSportType 删除运动类型
func (c *SportTypeController) DeleteSportType(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.sportTypeService.DeleteSportType(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
} 