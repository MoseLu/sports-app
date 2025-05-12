package models

import (
	"time"

	"gorm.io/gorm"
)

// UpdateLog 更新日志模型
type UpdateLog struct {
	gorm.Model
	OldVersion string    `gorm:"type:varchar(50);not null;comment:旧版本号"`    // 旧版本号
	NewVersion string    `gorm:"type:varchar(50);not null;comment:新版本号"`    // 新版本号
	Status     string    `gorm:"type:varchar(20);not null;comment:更新状态"`    // 更新状态：success/failed
	Error      string    `gorm:"type:text;comment:错误信息"`                    // 错误信息
	UpdatedAt  time.Time `gorm:"not null;comment:更新时间"`                     // 更新时间
}

// TableName 指定表名
func (UpdateLog) TableName() string {
	return "update_logs"
} 