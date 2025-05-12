package models

// Comment 评论模型
type Comment struct {
	BaseModel
	UserID    uint64  `gorm:"not null" json:"user_id"`
	User      User    `gorm:"foreignKey:UserID" json:"user"`
	CheckInID uint64  `gorm:"not null;index" json:"check_in_id"`
	CheckIn   CheckIn `gorm:"foreignKey:CheckInID" json:"check_in"`
	Content   string  `gorm:"type:text;not null" json:"content"`
	ParentID  *uint64 `gorm:"index" json:"parent_id"`
	Parent    *Comment `gorm:"foreignKey:ParentID" json:"parent"`
}

// TableName 指定表名
func (Comment) TableName() string {
	return "comments"
} 