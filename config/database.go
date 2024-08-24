package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	sl "homeschooledu/logger"
	"os"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		//panic("Error loading .env file")
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
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

	sl.Log.Info("Database connected successfully.")
}

func GetDB() *gorm.DB {
	return DB
}