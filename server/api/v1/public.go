package v1

import (
	"NotifierX/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 健康检测接口
func HealthApiHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

// 开发者信息接口
func InfoApiHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"developer": "DK",
		"email":     "ezopscn@gmail.com",
	})
}

// 系统版本接口
func VersionApiHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version": common.Version,
	})
}
