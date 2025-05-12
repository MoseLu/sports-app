package services

import (
	"fmt"
	"sports-app/backend/models"
	"time"

	"gorm.io/gorm"
)

// VerificationCode 验证码信息
type VerificationCode struct {
	Code      string
	Email     string
	ExpiresAt time.Time
}

// VerificationService 验证码服务
type VerificationService struct {
	db *gorm.DB
}

// NewVerificationService 创建验证码服务实例
func NewVerificationService(db *gorm.DB) *VerificationService {
	return &VerificationService{
		db: db,
	}
}

// SaveCode 保存验证码到数据库
func (s *VerificationService) SaveCode(email, code string) error {
	verification := models.Verification{
		Email:     email,
		Code:      code,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}

	if err := s.db.Create(&verification).Error; err != nil {
		return fmt.Errorf("保存验证码失败: %v", err)
	}

	return nil
}

// SendVerificationEmail 发送验证码邮件
func (s *VerificationService) SendVerificationEmail(email, code string) error {
	// TODO: 实现发送邮件的逻辑
	// 这里先模拟发送成功
	fmt.Printf("发送验证码 %s 到邮箱 %s\n", code, email)
	return nil
}

// VerifyCode 验证验证码
func (s *VerificationService) VerifyCode(email, code string) bool {
	if s.db == nil {
		return false
	}

	var verification models.Verification
	if err := s.db.Where("email = ? AND code = ? AND expires_at > ?", email, code, time.Now()).First(&verification).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Printf("验证码验证错误: %v\n", err)
		}
		return false
	}

	// 验证成功后删除验证码
	if err := s.db.Delete(&verification).Error; err != nil {
		fmt.Printf("删除验证码错误: %v\n", err)
	}

	return true
}

// StoreCode 存储验证码
func (s *VerificationService) StoreCode(email, code string) {
	// This method is removed as per the new implementation instructions
}

// GetCode 获取验证码
func (s *VerificationService) GetCode(email string) (string, bool) {
	// This method is removed as per the new implementation instructions
	return "", false
}

// DeleteCode 删除验证码
func (s *VerificationService) DeleteCode(email string) {
	// This method is removed as per the new implementation instructions
}

// cleanupExpiredCodes 清理过期的验证码
func (s *VerificationService) cleanupExpiredCodes() {
	// This method is removed as per the new implementation instructions
} 