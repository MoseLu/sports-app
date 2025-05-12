package controllers

import (
	"net/http"

	"sports-app/backend/models"
	"sports-app/backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// VerificationController 验证码控制器
type VerificationController struct {
	emailService       *services.EmailService
	verificationService *services.VerificationService
	db                *gorm.DB
}

// NewVerificationController 创建验证码控制器实例
func NewVerificationController(emailService *services.EmailService, verificationService *services.VerificationService, db *gorm.DB) *VerificationController {
	return &VerificationController{
		emailService:       emailService,
		verificationService: verificationService,
		db:                db,
	}
}

// SendCode 发送验证码
func (c *VerificationController) SendCode(ctx *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的邮箱地址"})
		return
	}

	// 检查用户是否存在
	var count int64
	if err := c.db.Model(&models.User{}).Where("email = ?", req.Email).Count(&count).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}

	if count == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "该邮箱未注册"})
		return
	}

	// 生成并发送验证码
	code, err := c.emailService.SendVerificationCode(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "发送验证码失败"})
		return
	}

	// 存储验证码
	c.verificationService.StoreCode(req.Email, code)

	ctx.JSON(http.StatusOK, gin.H{"message": "验证码已发送"})
}

// VerifyCode 验证验证码
func (c *VerificationController) VerifyCode(ctx *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	// 检查用户是否存在
	var count int64
	if err := c.db.Model(&models.User{}).Where("email = ?", req.Email).Count(&count).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}

	if count == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "该邮箱未注册"})
		return
	}

	// 获取存储的验证码
	storedCode, exists := c.verificationService.GetCode(req.Email)
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "验证码已过期或不存在"})
		return
	}

	// 验证验证码
	if !c.emailService.VerifyCode(req.Code, storedCode) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "验证码错误"})
		return
	}

	// 验证成功后删除验证码
	c.verificationService.DeleteCode(req.Email)

	ctx.JSON(http.StatusOK, gin.H{"message": "验证码验证成功"})
} 