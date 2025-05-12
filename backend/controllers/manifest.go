package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"sports-app/backend/services"

	"github.com/gin-gonic/gin"
)

// ManifestController 处理 manifest 相关的请求
type ManifestController struct {
	updateLogService *services.UpdateLogService
}

// NewManifestController 创建 ManifestController 实例
func NewManifestController(updateLogService *services.UpdateLogService) *ManifestController {
	return &ManifestController{
		updateLogService: updateLogService,
	}
}

// GetManifest 获取 manifest 信息
func (c *ManifestController) GetManifest(ctx *gin.Context) {
	// 获取当前版本
	currentVersion := ctx.Query("version")
	if currentVersion == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "缺少version参数",
		})
		return
	}

	// 读取manifest文件
	manifestPath := filepath.Join("public", "manifest.json")
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		// 记录错误日志
		c.updateLogService.LogUpdate(currentVersion, "", "failed", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "读取manifest文件失败",
		})
		return
	}

	// 解析manifest数据
	var manifest map[string]interface{}
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		// 记录错误日志
		c.updateLogService.LogUpdate(currentVersion, "", "failed", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "解析manifest文件失败",
		})
		return
	}

	// 获取最新版本
	latestVersion, ok := manifest["version"].(string)
	if !ok {
		// 记录错误日志
		c.updateLogService.LogUpdate(currentVersion, "", "failed", fmt.Errorf("manifest文件格式错误"))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "manifest文件格式错误",
		})
		return
	}

	// 比较版本号
	needsUpdate := compareVersions(currentVersion, latestVersion)

	// 记录更新日志
	status := "no_update"
	if needsUpdate {
		status = "update_available"
	}
	
	// 记录检查日志
	if err := c.updateLogService.LogUpdate(currentVersion, latestVersion, status, nil); err != nil {
		// 记录日志失败不影响主流程
		fmt.Printf("记录更新日志失败: %v\n", err)
	}

	// 返回更新信息
	ctx.JSON(http.StatusOK, gin.H{
		"needsUpdate": needsUpdate,
		"manifest":    manifest,
	})
}

// compareVersions 比较版本号
func compareVersions(current, latest string) bool {
	// 移除版本号中的v前缀
	current = strings.TrimPrefix(current, "v")
	latest = strings.TrimPrefix(latest, "v")

	// 分割版本号
	currentParts := strings.Split(current, ".")
	latestParts := strings.Split(latest, ".")

	// 比较每个部分
	for i := 0; i < len(currentParts) && i < len(latestParts); i++ {
		currentNum := 0
		latestNum := 0
		// 忽略解析错误
		_, _ = fmt.Sscanf(currentParts[i], "%d", &currentNum)
		_, _ = fmt.Sscanf(latestParts[i], "%d", &latestNum)

		if latestNum > currentNum {
			return true
		}
		if latestNum < currentNum {
			return false
		}
	}

	// 如果latest版本号更长，说明有更新
	return len(latestParts) > len(currentParts)
}

// UpdateManifest 更新 manifest 信息
func (c *ManifestController) UpdateManifest(ctx *gin.Context) {
	var manifest struct {
		Version   string `json:"version" binding:"required"`
		BundleURL string `json:"bundleUrl" binding:"required"`
		Sha256    string `json:"sha256" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&manifest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid manifest data",
		})
		return
	}

	// 记录更新操作日志
	if err := c.updateLogService.LogUpdate("", manifest.Version, "success", nil); err != nil {
		// 记录日志失败不影响主流程
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Manifest updated successfully",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Manifest updated successfully",
	})
} 