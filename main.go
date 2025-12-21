package main

import (
	"chunli/config"
	"chunli/global"
	"chunli/models"
	"chunli/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化数据库
	global.Connect(cfg)

	// 自动迁移数据库表
	if err := global.DB.AutoMigrate(&models.UserModel{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
	fmt.Println("Database migration completed")

	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode)

	// 初始化 Gin 引擎
	r := gin.Default()

	// 设置路由
	routes.SetupRoutes(r)

	// 启动服务器
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Server starting on %s\n", addr)
	if err := r.Run(addr); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
