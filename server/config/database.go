package config

import (
	"fmt"
	"os"
	"share_movie/server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	// 从环境变量获取数据库配置
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "123456")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "share_movie")

	// 数据库连接配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 自动迁移数据库表结构
	err = db.AutoMigrate(
		&models.Movie{},
		&models.Link{},
		&models.Tag{},
		&models.User{},
	)
	if err != nil {
		return fmt.Errorf("数据库迁移失败: %v", err)
	}

	DB = db
	return nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
