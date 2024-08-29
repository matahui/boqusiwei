package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	sl "homeschooledu/logger"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		//panic("Error loading .env file")
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	//开启日志

	gormLogger := &sl.LogrusLogger{Log: sl.Log}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:gormLogger,
	})

	if err != nil {
		//panic("failed to connect to database")
		sl.Log.Fatal("Database connected failed", err)
		return
	}

	sqlDB, err := DB.DB()
	if err != nil {
		sl.Log.Fatal("Database getDB failed", err)
		return
	}

	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)

	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(100)

	// 设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	sl.Log.Info("Database connected successfully.")

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			printDBStats(sqlDB)
		}
	}()
}

func GetDB() *gorm.DB {
	return DB
}

func printDBStats(db *sql.DB) {
	stats := db.Stats()
	sl.Log.Infof("连接池状态 - 空闲连接数: %d, 打开连接数: %d, 最大打开连接数: %d, 等待中的连接数: %d\n",
		stats.Idle,
		stats.OpenConnections,
		stats.MaxOpenConnections,
		stats.WaitCount)
}