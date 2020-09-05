package app

import (
	"gin-mongo-backend/app/api"
	"gin-mongo-backend/app/controllers"
	"gin-mongo-backend/models"
	"gin-mongo-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func initService(config utils.Config)  {
	// 初始化数据库
	if err := models.InitDB(&config.DB); err != nil {
		panic(err)
	}

	// 初始化 JWT
	controllers.SetJWTSignKey(config.HTTP)
}

func Run(configPath string) {
	// 初始化日志
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// 读取配置
	var config utils.Config
	config.LoadConf(configPath)
	if config.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else {
		gin.SetMode(gin.DebugMode)
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// 初始化各种服务
	initService(config)

	// 启动服务器
	app := controllers.NewApp()

	api.RunHTTPServer(app)



	if err := app.Run(":" + config.HTTP.Port); err != nil {
		panic(err)
	}
}
