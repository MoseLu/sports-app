package database

import (
	"fmt"
	"log"
	"sync"

	"sports-app/backend/config"
	"sports-app/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDB 初始化数据库连接
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.DBName,
		)

		log.Printf("尝试连接数据库: %s@%s:%s/%s", cfg.DB.User, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName)

		// 使用详细的日志配置
		gormConfig := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}

		db, err = gorm.Open(mysql.Open(dsn), gormConfig)
		if err != nil {
			err = fmt.Errorf("连接数据库失败: %v", err)
			log.Printf("数据库连接错误: %v", err)
			return
		}

		// 检查数据库连接
		sqlDB, err := db.DB()
		if err != nil {
			err = fmt.Errorf("获取数据库实例失败: %v", err)
			log.Printf("数据库实例错误: %v", err)
			return
		}

		// 测试数据库连接
		err = sqlDB.Ping()
		if err != nil {
			err = fmt.Errorf("数据库 Ping 失败: %v", err)
			log.Printf("数据库 Ping 错误: %v", err)
			return
		}

		// 自动迁移数据库表
		err = db.AutoMigrate(&models.User{}, &models.SportRecord{}, &models.SportType{})
		if err != nil {
			err = fmt.Errorf("自动迁移失败: %v", err)
			log.Printf("数据库迁移错误: %v", err)
			return
		}

		log.Println("数据库连接成功")
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetDB 获取数据库连接实例
func GetDB() *gorm.DB {
	return db
}
