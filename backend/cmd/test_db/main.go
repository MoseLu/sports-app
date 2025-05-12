package main

import (
	"fmt"
	"log"

	"sports-app/backend/config"
)

func main() {
	// 获取配置
	cfg := config.GetConfig()
	fmt.Printf("数据库配置: %+v\n", cfg.DB)

	// 获取数据库连接
	db := config.GetDB()
	if db == nil {
		log.Fatal("无法获取数据库连接")
	}

	// 测试连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库连接失败: %v", err)
	}

	// 测试ping
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}

	fmt.Println("数据库连接测试成功！")
}
