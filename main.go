package main

import (
	"github.com/gin-gonic/gin"
	"homeschooledu/config"
	"homeschooledu/consts"
	"homeschooledu/logger"
	"homeschooledu/routers"
	"os"
)

func main() {
	// 初始化日志记录器
	logger.InitLogger()

	// 初始化数据库
	config.InitDB()

	env := os.Getenv("Environment")
	if env == consts.EnvPrd {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化Gin引擎
	r := gin.New()

	gin.DefaultWriter = logger.Log.Writer()
	gin.DefaultErrorWriter = logger.Log.Writer()

	// 使用 Gin 的 Logger 和 Recovery 中间件
	r.Use(gin.LoggerWithWriter(logger.Log.Writer()), gin.RecoveryWithWriter(logger.Log.Writer()))
	r.Use(gin.Recovery())


	// 加载路由
	routers.SetupRouters(r)

	// 启动服务器
	logger.Log.Info("Starting server on port 9123")
	if err := r.Run(":9123"); err != nil {
		logger.Log.Fatal("Server failed to start: ", err)
	}

	logger.Log.Info("Starting server on port 9123 success")
}
