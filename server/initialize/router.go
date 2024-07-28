package initialize

import (
	"NotifierX/common"
	"NotifierX/middleware"
	"NotifierX/route"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 路由初始化
func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	// 创建一个没中间件的路由引擎
	r := gin.New()

	// 中间件
	r.Use(middleware.AccessLog) // 请求日志中间件
	r.Use(middleware.Cors)      // 跨域访问中间件

	// 无需登录的接口路由
	publicApis := r.Group(common.APIPrefix)
	publicApis.Use(middleware.Exception)
	route.PublicApiRoutes(publicApis)

	// 其它
	r.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/404")
	})

	log.Println("系统路由初始化完成")
	return r
}
