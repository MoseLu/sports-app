// backend/main.go
package main

import (
	"log"
	"sports-app/backend/config"
	"sports-app/backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. 加载 .env 文件（CI 已经写入到后端目录中）
	if err := godotenv.Load(); err != nil {
		log.Println("未找到 .env 文件，或加载失败，将尝试使用系统环境变量")
	}

	// 2. 初始化配置（包括从环境变量读取 OSS 配置）
	config.GetConfig()

	// 3. 初始化数据库连接
	db := config.GetDB()
	logsDB := config.GetLogsDB()

	// 4. 设置 Gin 路由
	r := gin.Default()

	// 5. 把 /static/ 下所有文件映射到 ./static 目录
	r.Static("/static", "./static")

	// 同时把 /favicon.ico 直接映射到 ./static/favicon.ico
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	// 下面再注册你的 API 路由
	routes.SetupRoutes(r, db, logsDB)

	// 6. 启动 HTTP 服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
