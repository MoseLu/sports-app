package config

// EmailType 邮箱类型
type EmailType string

const (
	EmailTypeQQ     EmailType = "qq"
	EmailTypeOutlook EmailType = "outlook"
)

// EmailConfig 邮箱配置
type EmailConfig struct {
	Type     EmailType `json:"type"`     // 邮箱类型
	Host     string    `json:"host"`     // SMTP服务器地址
	Port     int       `json:"port"`     // SMTP服务器端口
	Username string    `json:"username"` // 邮箱账号
	Password string    `json:"password"` // 邮箱密码或授权码
	From     string    `json:"from"`     // 发件人
}

// GetEmailConfig 获取邮箱配置
func GetEmailConfig(emailType EmailType) *EmailConfig {
	var config *EmailConfig
	
	switch emailType {
	case EmailTypeOutlook:
		config = &EmailConfig{
			Type:     EmailTypeOutlook,
			Host:     "smtp-mail.outlook.com",
			Port:     587,
			Username: "mlu@bellis-technology.cn",
			Password: "lah1999626123",
			From:     "mlu@bellis-technology.cn",
		}
	case EmailTypeQQ:
		fallthrough
	default:
		config = &EmailConfig{
			Type:     EmailTypeQQ,
			Host:     "smtp.qq.com",
			Port:     587,
			Username: "1208136885@qq.com",
			Password: "kqdhbdersfshjiae",
			From:     "1208136885@qq.com",
		}
	}
	
	return config
} 