package controllers

import (
	"gin-mongo-backend/middleware"
	"gin-mongo-backend/utils"
	"github.com/gin-gonic/gin"
)

// NewApp 创建服务器实例并绑定控制器
func NewApp() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.Cors())

	return app
}

//JWT 获取signKey
func SetJWTSignKey(conf utils.HTTPConfig) {
	middleware.SignKey = conf.JWT.Key
}
