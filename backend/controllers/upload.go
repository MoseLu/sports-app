package controllers

import (
	"net/http"

	"sports-app/backend/services"

	"github.com/gin-gonic/gin"
)

type UploadController struct {
	uploadService *services.UploadService
}

func NewUploadController() *UploadController {
	return &UploadController{
		uploadService: &services.UploadService{},
	}
}

// UploadImage 处理图片上传
func (c *UploadController) UploadImage(ctx *gin.Context) {
	// 获取用户ID
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	// 获取上传的文件
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "获取文件失败"})
		return
	}

	// 验证文件类型
	if !isValidImageType(file.Header.Get("Content-Type")) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "不支持的文件类型"})
		return
	}

	// 验证文件大小（限制为5MB）
	if file.Size > 5*1024*1024 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "文件大小超过限制"})
		return
	}

	// 上传图片
	imageURL, err := c.uploadService.UploadImage(file, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "上传成功",
		"url":     imageURL,
	})
}

// isValidImageType 验证文件类型
func isValidImageType(contentType string) bool {
	validTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}
	return validTypes[contentType]
} 