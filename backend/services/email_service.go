package services

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"sports-app/backend/config"

	"gopkg.in/gomail.v2"
)

// EmailService 邮箱服务
type EmailService struct {
}

// NewEmailService 创建邮箱服务实例
func NewEmailService() *EmailService {
	rand.Seed(time.Now().UnixNano())
	return &EmailService{}
}

// isOutlookEmail 检查是否是 Outlook 邮箱
func isOutlookEmail(email string) bool {
	// 只检查 Outlook 相关后缀
	email = strings.ToLower(email)
	if strings.HasSuffix(email, "@outlook.com") || 
	   strings.HasSuffix(email, "@hotmail.com") || 
	   strings.HasSuffix(email, "@msn.com") {
		return true
	}

	// 对于其他域名，尝试连接 Outlook 服务器
	config := config.GetEmailConfig(config.EmailTypeOutlook)
	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	
	// 尝试连接
	s, err := d.Dial()
	if err != nil {
		fmt.Printf("连接 Outlook 服务器失败: %v\n", err)
		return false
	}
	defer s.Close()
	
	return true
}

// getEmailConfigByAddress 根据邮箱地址获取对应的配置
func getEmailConfigByAddress(email string) *config.EmailConfig {
	
	if isOutlookEmail(email) {
		return config.GetEmailConfig(config.EmailTypeOutlook)
	}
	
	return config.GetEmailConfig(config.EmailTypeQQ)
}

// generateVerificationCode 生成6位数字验证码
func (s *EmailService) generateVerificationCode() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// SendVerificationCode 发送验证码
func (s *EmailService) SendVerificationCode(email string) (string, error) {
	// 生成6位随机验证码
	code := s.generateVerificationCode()
	
	// 根据邮箱地址获取对应的配置
	config := getEmailConfigByAddress(email)
	
	// 创建邮件
	m := gomail.NewMessage()
	m.SetHeader("From", config.From)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Re Sports - 验证码")
	m.SetBody("text/html", fmt.Sprintf(`
		<div style="max-width: 600px; margin: 0 auto; padding: 20px; font-family: Arial, sans-serif;">
			<h2 style="color: #1976D2;">Re Sports</h2>
			<p>您的验证码是：</p>
			<div style="background: #f5f5f5; padding: 10px; text-align: center; font-size: 24px; letter-spacing: 5px; margin: 20px 0;">
				<strong>%s</strong>
			</div>
			<p>验证码有效期为5分钟，请尽快使用。</p>
			<p>如果这不是您的操作，请忽略此邮件。</p>
		</div>
	`, code))

	// 发送邮件
	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("发送邮件失败: %v\n", err)
		return "", err
	}

	return code, nil
}

// VerifyCode 验证验证码
func (s *EmailService) VerifyCode(inputCode, storedCode string) bool {
	return inputCode == storedCode
} 