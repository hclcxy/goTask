package config

import (
	"go-task4/pkg/logger"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.NewGormLogger(),
	})

	if err != nil {
		log.Fatal("failed to connect database:", err)
		panic("failed to connect database")
	}

	DB = db
}
