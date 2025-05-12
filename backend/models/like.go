package models

// Like 点赞模型
type Like struct {
	BaseModel
	UserID    uint64  `gorm:"not null;uniqueIndex:uk_user_check_in" json:"user_id"`
	User      User    `gorm:"foreignKey:UserID" json:"user"`
	CheckInID uint64  `gorm:"not null;uniqueIndex:uk_user_check_in" json:"check_in_id"`
	CheckIn   CheckIn `gorm:"foreignKey:CheckInID" json:"check_in"`
}

// TableName 指定表名
func (Like) TableName() string {
	return "likes"
} 