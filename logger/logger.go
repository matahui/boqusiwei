package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/lestrrat-go/file-rotatelogs"
	"gorm.io/gorm/logger"
	"time"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	// 配置 rotatelogs
	path := "app.log"
	writer, err := rotatelogs.New(
		path+".%Y%m%d",                       // 按天切割日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour), // 日志文件保留时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	if err != nil {
		Log.Fatalf("Failed to initialize log file rotator: %v", err)
	}

	// 设置 logrus 输出
	Log.SetOutput(writer)

	// 配置 logrus 的日志格式和日志级别
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.InfoLevel)

	// 设置日志格式
	Log.SetFormatter(&logrus.JSONFormatter{})
}


type LogrusLogger struct {
	Log *logrus.Logger
}

// LogMode 设置日志级别
func (l *LogrusLogger) LogMode(level logger.LogLevel) logger.Interface {
	return &LogrusLogger{Log: l.Log}
}

// Info 记录信息级别的日志
func (l *LogrusLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	l.Log.WithFields(logrus.Fields{
		"component": "gorm",
		"level":     "info",
	}).Infof(msg, args...)
}

// Warn 记录警告级别的日志
func (l *LogrusLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	l.Log.WithFields(logrus.Fields{
		"component": "gorm",
		"level":     "warn",
	}).Warnf(msg, args...)
}

// Error 记录错误级别的日志
func (l *LogrusLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	l.Log.WithFields(logrus.Fields{
		"component": "gorm",
		"level":     "error",
	}).Errorf(msg, args...)
}

// Trace 记录追踪级别的日志
func (l *LogrusLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	duration := time.Since(begin)

	if err != nil {
		l.Log.WithFields(logrus.Fields{
			"component": "gorm",
			"level":     "error",
			"duration":  duration,
			"sql":       sql,
			"rows_affected": rowsAffected,
			"error":     err,
		}).Errorf("SQL error: %v", err)
	} else {
		l.Log.WithFields(logrus.Fields{
			"component": "gorm",
			"level":     "info",
			"duration":  duration,
			"sql":       sql,
			"rows_affected": rowsAffected,
		}).Info("SQL executed successfully")
	}
}