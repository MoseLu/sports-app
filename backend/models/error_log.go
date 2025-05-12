package models

import (
	"time"
)

// ErrorLog 错误日志模型
type ErrorLog struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	ErrorType   string    `gorm:"not null;default:'unknown'" json:"error_type"`   // 错误类型
	Message     string    `gorm:"not null" json:"message"`                        // 错误消息
	StatusCode  int       `gorm:"not null;default:500" json:"status_code"`        // HTTP状态码
	Path        string    `gorm:"not null" json:"path"`                           // 请求路径
	Method      string    `gorm:"not null" json:"method"`                         // 请求方法
	RequestBody string    `gorm:"type:text" json:"request_body"`                  // 请求体
	ResponseBody string   `gorm:"type:text" json:"response_body"`                 // 响应体
	CreatedAt   time.Time `json:"created_at"`                                     // 创建时间
} 