package handlers

import (
	"net/http"
	"sports-app/backend/models"
	"sports-app/backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
    db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
    return &AuthController{db: db}
}

func (c *AuthController) Register(ctx *gin.Context) {
    var user models.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.db.Create(&user).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, user)
}

func (c *AuthController) Login(ctx *gin.Context) {
    var loginData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := ctx.ShouldBindJSON(&loginData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := c.db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
        return
    }

    if user.Password != loginData.Password {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
        return
    }

    token, err := services.GenerateToken(user.ID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "token": token,
        "user":  user,
    })
}

func (c *AuthController) Logout(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "登出成功"})
} 