package controllers

import (
	"net/http"
	"sports-app/backend/models"
	"sports-app/backend/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := c.service.Register(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	user, err := c.service.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	
	token, err := services.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "生成token失败"})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "登出成功"})
}

// ResetPassword 发送重置密码邮件
func (c *AuthController) ResetPassword(ctx *gin.Context) {
	var req struct {
		Email       string `json:"email" binding:"required,email"`
		Code        string `json:"code" binding:"required"`
		NewPassword string `json:"newPassword" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	// 验证验证码
	if !c.service.VerifyCode(req.Email, req.Code) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "验证码错误或已过期"})
		return
	}

	// 重置密码
	username, err := c.service.ResetPassword(req.Email, req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "重置密码失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "密码重置成功",
		"data": gin.H{
			"username": username,
		},
	})
}

func (c *AuthController) SendCode(ctx *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	if err := c.service.SendCode(req.Email); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "验证码已发送"})
}

func (c *AuthController) VerifyCode(ctx *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	valid := c.service.VerifyCode(req.Email, req.Code)
	if !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "验证码错误或已过期"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "验证码验证成功"})
} 