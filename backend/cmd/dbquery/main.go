package main

import (
	"fmt"
	"log"

	"sports-app/backend/config"
	"sports-app/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("查询用户失败: %v", err)
	}

	fmt.Println("\n当前用户列表:")
	fmt.Println("ID\t用户名\t\t邮箱\t\t\t创建时间")
	fmt.Println("------------------------------------------------")
	for _, user := range users {
		fmt.Printf("%d\t%s\t\t%s\t\t%s\n",
			user.ID,
			user.Username,
			user.Email,
			user.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
