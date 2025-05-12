package services

import (
	"sports-app/backend/models"
	"time"

	"gorm.io/gorm"
)

// RecordService 运动记录服务
type RecordService struct {
	db *gorm.DB
}

// NewRecordService 创建运动记录服务实例
func NewRecordService(db *gorm.DB) *RecordService {
	return &RecordService{db: db}
}

// GetRecords 获取用户的运动记录列表
func (s *RecordService) GetRecords(userID int64) ([]models.SportRecord, error) {
	var records []models.SportRecord
	if err := s.db.Preload("SportType").Where("user_id = ?", userID).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// CreateRecord 创建运动记录
func (s *RecordService) CreateRecord(record *models.SportRecord) error {
	return s.db.Create(record).Error
}

// UpdateRecord 更新运动记录
func (s *RecordService) UpdateRecord(record *models.SportRecord) error {
	return s.db.Model(&models.SportRecord{}).Where("id = ?", record.ID).Updates(map[string]interface{}{
		"sport_type_id": record.SportTypeID,
		"exercise":      record.Exercise,
		"duration":      record.Duration,
		"calories":      record.Calories,
		"start_time":    record.StartTime,
		"end_time":      record.EndTime,
		"image_url":     record.ImageURL,
		"img_url_list":  record.ImgURLList,
		"updated_at":    time.Now(),
	}).Error
}

// DeleteRecord 删除运动记录
func (s *RecordService) DeleteRecord(id int64, userID int64) error {
	return s.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.SportRecord{}).Error
}

// GetSportTypes 获取所有运动类型
func (s *RecordService) GetSportTypes() ([]models.SportType, error) {
	var types []models.SportType
	err := s.db.Find(&types).Error
	return types, err
}

// CreateSportType 创建运动类型
func (s *RecordService) CreateSportType(sportType *models.SportType) error {
	return s.db.Create(sportType).Error
}

// UpdateSportType 更新运动类型
func (s *RecordService) UpdateSportType(sportType *models.SportType) error {
	return s.db.Save(sportType).Error
}

// DeleteSportType 删除运动类型
func (s *RecordService) DeleteSportType(id int64) error {
	return s.db.Delete(&models.SportType{}, id).Error
}

// GetStats 获取用户的运动统计信息
func (s *RecordService) GetStats(userID int64, timeRange string, sportTypeID int64) (*models.Stats, error) {
	var stats models.Stats
	
	// 构建基础查询
	baseQuery := s.db.Model(&models.SportRecord{}).Where("user_id = ?", userID)
	
	// 添加时间范围筛选
	now := time.Now()
	switch timeRange {
	case "week":
		startTime := now.AddDate(0, 0, -7)
		baseQuery = baseQuery.Where("start_time >= ?", startTime)
	case "month":
		startTime := now.AddDate(0, -1, 0)
		baseQuery = baseQuery.Where("start_time >= ?", startTime)
	case "year":
		startTime := now.AddDate(-1, 0, 0)
		baseQuery = baseQuery.Where("start_time >= ?", startTime)
	}
	
	// 添加运动类型筛选
	if sportTypeID > 0 {
		baseQuery = baseQuery.Where("sport_type_id = ?", sportTypeID)
	}
	
	// 获取总运动时长
	if err := baseQuery.Select("COALESCE(SUM(duration), 0) as total_duration").
		Scan(&stats.TotalDuration).Error; err != nil {
		return nil, err
	}
	
	// 获取运动次数
	if err := baseQuery.Count(&stats.ExerciseCount).Error; err != nil {
		return nil, err
	}
	
	// 获取平均运动时长
	if err := baseQuery.Select("COALESCE(AVG(duration), 0) as average_duration").
		Scan(&stats.AverageDuration).Error; err != nil {
		return nil, err
	}
	
	// 获取平均消耗卡路里
	if err := baseQuery.Select("COALESCE(AVG(calories), 0) as average_calories").
		Scan(&stats.AverageCalories).Error; err != nil {
		return nil, err
	}

	// 初始化每日统计数组
	stats.DailyDuration = make([]int64, 7)
	stats.DailyCount = make([]int64, 7)

	// 获取最近7天的每日统计
	for i := 0; i < 7; i++ {
		// 计算日期范围
		startDate := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		endDate := time.Now().AddDate(0, 0, -i+1).Format("2006-01-02")

		// 构建每日查询
		dailyQuery := s.db.Model(&models.SportRecord{}).
			Where("user_id = ? AND start_time >= ? AND start_time < ?", userID, startDate, endDate)
		
		// 添加运动类型筛选
		if sportTypeID > 0 {
			dailyQuery = dailyQuery.Where("sport_type_id = ?", sportTypeID)
		}

		// 获取每日运动时长
		var dailyDuration int64
		if err := dailyQuery.Select("COALESCE(SUM(duration), 0)").
			Scan(&dailyDuration).Error; err != nil {
			return nil, err
		}
		stats.DailyDuration[6-i] = dailyDuration

		// 获取每日运动次数
		var dailyCount int64
		if err := dailyQuery.Count(&dailyCount).Error; err != nil {
			return nil, err
		}
		stats.DailyCount[6-i] = dailyCount
	}
	
	return &stats, nil
} 