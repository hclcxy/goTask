package main

import (
	"go-task4/api"
	"go-task4/config"
	"go-task4/middleware"
	"go-task4/models"
	"go-task4/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志
	logger.Init()

	// 初始化数据库
	config.ConnectDB()

	// 自动迁移模型
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	); err != nil {
		log.Fatal("failed to auto migrate models:", err)
		panic(err)
	}

	// 初始化Gin
	router := gin.Default()

	// 添加中间件
	router.Use(middleware.LoggerMiddleware())

	// 注册路由
	api.RegisterRoutes(router, config.DB)

	// 启动服务器
	if err := router.Run(":9000"); err != nil {
		log.Fatal("failed to start server:", err)
		panic(err)
	}
}
