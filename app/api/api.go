package api

import (
	"github.com/gin-gonic/gin"
)

func RunHTTPServer(engine *gin.Engine) {
	InitAPI(engine)
	ToolsAPI(engine)
}
