package api

import (
	"gin-mongo-backend/app/handler"
	"gin-mongo-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitAPI(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	v1.POST("/configs", handler.GetAppConfigs)

	v1.Use(middleware.JWTAuth())

	v1.PUT("/configs/:_id", handler.UpdateAppConfigs)
}
