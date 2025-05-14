package services

import (
	"errors"
	"fmt"
	"math/rand"
	"sports-app/backend/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your-secret-key")

type Claims struct {
    UserID int64 `json:"user_id"`
    jwt.RegisteredClaims
}

type AuthService struct {
    db                  *gorm.DB
    verificationService *VerificationService
}

func NewAuthService(db *gorm.DB, verificationService *VerificationService) *AuthService {
    return &AuthService{
        db:                  db,
        verificationService: verificationService,
    }
}

func (s *AuthService) Register(user *models.User) error {
    // 检查用户名是否已存在
    var count int64
    if err := s.db.Model(&models.User{}).Where("username = ?", user.Username).Count(&count).Error; err != nil {
        return err
    }
    if count > 0 {
        return errors.New("用户名已存在")
    }
    
    // 加密密码
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    
    // 创建用户
    return s.db.Create(user).Error
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	
	if !user.CheckPassword(password) {
		return nil, errors.New("用户名或密码错误")
	}
	
	// 更新最后登录时间
	user.LastLoginAt = time.Now()
	if err := s.db.Save(&user).Error; err != nil {
		return nil, fmt.Errorf("更新登录时间失败: %v", err)
	}
	
	return &user, nil
}

func GenerateToken(userID int64) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    
    if err != nil {
        return nil, err
    }
    
    if !token.Valid {
        return nil, errors.New("无效的token")
    }
    
    return claims, nil
}

func (s *AuthService) VerifyCode(email, code string) bool {
	if s.verificationService == nil {
		return false
	}
	return s.verificationService.VerifyCode(email, code)
}

func (s *AuthService) ResetPassword(email, newPassword string) (string, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("用户不存在")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// 更新密码
	if err := s.db.Model(&user).Update("password", string(hashedPassword)).Error; err != nil {
		return "", err
	}

	return user.Username, nil
}

func (s *AuthService) SendCode(email string) error {
	// 检查邮箱是否已注册
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("该邮箱未注册")
		}
		return fmt.Errorf("系统错误")
	}

	// 生成验证码
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	
	// 保存验证码
	if err := s.verificationService.SaveCode(email, code); err != nil {
		return fmt.Errorf("保存验证码失败")
	}

	// 发送验证码邮件
	if err := s.verificationService.SendVerificationEmail(email, code); err != nil {
		return fmt.Errorf("发送验证码失败")
	}

	return nil
} 