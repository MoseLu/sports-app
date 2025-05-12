package migrations

import (
	"sports-app/backend/models"

	"gorm.io/gorm"
)

// CreateSportRecordsTable 创建运动记录表
func CreateSportRecordsTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.SportRecord{})
}

// DropSportRecordsTable 删除运动记录表
func DropSportRecordsTable(db *gorm.DB) error {
	return db.Migrator().DropTable("sport_records")
} 