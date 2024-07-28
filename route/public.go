package route

import (
	v1 "NotifierX/api/v1"
	"github.com/gin-gonic/gin"
)

// 无需登录的接口路由组
func PublicApiRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/health", v1.HealthApiHandler)   // 健康检查接口
	rg.GET("/info", v1.InfoApiHandler)       // 开发者信息接口
	rg.GET("/version", v1.VersionApiHandler) // 版本信息接口
	return rg
}
