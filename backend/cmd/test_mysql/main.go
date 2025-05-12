package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 不指定数据库的 DSN
	dsn := "root:123456@tcp(localhost:3306)/"

	// 连接 MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("无法创建数据库连接: %v", err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("无法连接到 MySQL: %v", err)
	}
	fmt.Println("成功连接到 MySQL 服务器")

	// 尝试创建数据库
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS sports_app")
	if err != nil {
		log.Fatalf("创建数据库失败: %v", err)
	}
	fmt.Println("成功创建数据库 sports_app")

	// 尝试使用数据库
	_, err = db.Exec("USE sports_app")
	if err != nil {
		log.Fatalf("使用数据库失败: %v", err)
	}
	fmt.Println("成功切换到数据库 sports_app")

	// 创建用户表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			deleted_at DATETIME NULL,
			username VARCHAR(20) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			role VARCHAR(20) DEFAULT 'user',
			timezone VARCHAR(50) DEFAULT 'Asia/Shanghai',
			INDEX idx_users_deleted_at (deleted_at)
		)
	`)
	if err != nil {
		log.Fatalf("创建用户表失败: %v", err)
	}
	fmt.Println("成功创建用户表")
}
