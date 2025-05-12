package services

import (
	"log"
	"sports-app/backend/models"
	"strconv"

	"gorm.io/gorm"
)

// SportTypeService 运动类型服务
type SportTypeService struct {
	db *gorm.DB
}

// NewSportTypeService 创建运动类型服务实例
func NewSportTypeService(db *gorm.DB) *SportTypeService {
	return &SportTypeService{db: db}
}

// GetSportTypes 获取所有运动类型
func (s *SportTypeService) GetSportTypes() ([]models.SportType, error) {
	log.Println("开始获取运动类型...")
	var types []models.SportType
	err := s.db.Find(&types).Error
	if err != nil {
		log.Printf("获取运动类型失败: %v", err)
		return nil, err
	}
	log.Printf("成功获取到 %d 个运动类型", len(types))
	return types, nil
}

// CreateSportType 创建运动类型
func (s *SportTypeService) CreateSportType(sportType *models.SportType) error {
	return s.db.Create(sportType).Error
}

// UpdateSportType 更新运动类型
func (s *SportTypeService) UpdateSportType(id string, sportType *models.SportType) error {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	return s.db.Model(&models.SportType{}).Where("id = ?", idInt).Updates(sportType).Error
}

// DeleteSportType 删除运动类型
func (s *SportTypeService) DeleteSportType(id string) error {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	return s.db.Delete(&models.SportType{}, idInt).Error
} 