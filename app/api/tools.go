package api

import (
	"gin-mongo-backend/app/handler"
	"github.com/gin-gonic/gin"
)

func ToolsAPI(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	v1.POST("/getHTML", handler.GetHTML)
}