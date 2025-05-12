package main

import (
	"log"
	"sports-app/backend/config"
	"sports-app/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置（必须调用一次，保证 cfg 不为 nil）
	config.GetConfig()

	// 初始化数据库连接
	db := config.GetDB()
	logsDB := config.GetLogsDB()

	// 设置路由
	r := gin.Default()
	routes.SetupRoutes(r, db, logsDB)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
