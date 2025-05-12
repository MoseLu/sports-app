package services

import (
	"time"

	"sports-app/backend/models"

	"gorm.io/gorm"
)

// UpdateLogService 更新日志服务
type UpdateLogService struct {
	db *gorm.DB
}

// NewUpdateLogService 创建更新日志服务实例
func NewUpdateLogService(db *gorm.DB) *UpdateLogService {
	return &UpdateLogService{db: db}
}

// LogUpdate 记录更新日志
func (s *UpdateLogService) LogUpdate(oldVersion, newVersion, status string, err error) error {
	log := &models.UpdateLog{
		OldVersion: oldVersion,
		NewVersion: newVersion,
		Status:     status,
		UpdatedAt:  time.Now(),
	}

	// 如果有错误，记录错误信息
	if err != nil {
		log.Error = err.Error()
	}

	// 保存到数据库
	return s.db.Create(log).Error
}

// GetUpdateLogs 获取更新日志列表
func (s *UpdateLogService) GetUpdateLogs(page, pageSize int) ([]models.UpdateLog, int64, error) {
	var logs []models.UpdateLog
	var total int64

	// 获取总数
	if err := s.db.Model(&models.UpdateLog{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := s.db.Order("updated_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// GetUpdateStats 获取更新统计信息
func (s *UpdateLogService) GetUpdateStats() (map[string]interface{}, error) {
	var stats struct {
		TotalUpdates   int64
		SuccessUpdates int64
		FailedUpdates  int64
		LastUpdate     time.Time
	}

	// 获取总更新次数
	if err := s.db.Model(&models.UpdateLog{}).Count(&stats.TotalUpdates).Error; err != nil {
		return nil, err
	}

	// 获取成功更新次数
	if err := s.db.Model(&models.UpdateLog{}).Where("status = ?", "success").Count(&stats.SuccessUpdates).Error; err != nil {
		return nil, err
	}

	// 获取失败更新次数
	if err := s.db.Model(&models.UpdateLog{}).Where("status = ?", "failed").Count(&stats.FailedUpdates).Error; err != nil {
		return nil, err
	}

	// 获取最后更新时间
	if err := s.db.Model(&models.UpdateLog{}).Order("updated_at DESC").Limit(1).Pluck("updated_at", &stats.LastUpdate).Error; err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_updates":   stats.TotalUpdates,
		"success_updates": stats.SuccessUpdates,
		"failed_updates":  stats.FailedUpdates,
		"last_update":     stats.LastUpdate,
	}, nil
} 