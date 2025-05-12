package models

import (
	"time"

	"gorm.io/gorm"
)

// SportType 运动类型模型
type SportType struct {
	ID          int64          `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:50;not null;index" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Icon        string         `gorm:"size:255" json:"icon"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (SportType) TableName() string {
	return "sport_types"
} 