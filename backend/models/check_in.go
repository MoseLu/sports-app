package models

// CheckIn 打卡记录模型
type CheckIn struct {
	BaseModel
	UserID      uint64    `gorm:"not null;index" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	SportTypeID uint64    `gorm:"not null" json:"sport_type_id"`
	SportType   SportType `gorm:"foreignKey:SportTypeID" json:"sport_type"`
	Duration    int       `gorm:"not null" json:"duration"` // 运动时长（分钟）
	Images      string    `gorm:"type:text" json:"images"`  // 图片URL，多个用逗号分隔
	Description string    `gorm:"type:text" json:"description"`
	IsShared    bool      `gorm:"default:false" json:"is_shared"`
}

// TableName 指定表名
func (CheckIn) TableName() string {
	return "check_ins"
} 