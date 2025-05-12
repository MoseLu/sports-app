package models

import (
	"time"
)

// SportRecord 运动记录模型
type SportRecord struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	UserID      int64     `json:"user_id"`
	SportTypeID int64     `json:"sport_type_id" gorm:"not null"`
	SportType   SportType `json:"sport_type" gorm:"foreignKey:SportTypeID"`
	Exercise    string    `json:"exercise"`
	Duration    int64     `json:"duration"`
	Calories    int64     `json:"calories"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time"`
	ImageURL    string    `json:"image_url" gorm:"size:255"`
	ImgURLList  string    `json:"img_url_list" gorm:"type:json"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName 设置表名
func (SportRecord) TableName() string {
	return "sport_records"
} 