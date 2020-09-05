package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ApplicationJson = "application/json"
)

// 错误处理中间件
func ErrorHandler(context *gin.Context) {
	context.Next()

	// TODO
	if len(context.Errors) > 0 {
		ct := context.Request.Header.Get("Content-Type")
		if strings.Contains(ct, ApplicationJson) {
			context.JSON(http.StatusBadRequest, gin.H{"error": context.Errors})
		} else {
			context.HTML(http.StatusBadRequest, "400", gin.H{"error": context.Errors})
		}
	}
}
