package initializer

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nuanxinqing123/flying-bird/internal/middleware"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(middleware.Logger())
	Router.Use(middleware.Recovery())

	// 跨域配置
	Router.Use(cors.New(middleware.CorsConfig))

	// (可选项)
	// PID 限流基于实例的 CPU 使用率，通过拒绝一定比例的流量, 将实例的 CPU 使用率稳定在设定的阈值上。
	// 地址: https://github.com/bytedance/pid_limits
	// Router.Use(adaptive.PlatoMiddlewareGinDefault(0.8))

	// 存活检测
	Router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//ApiGroup := Router.Group("/app")
	//router.InitRouterApp(ApiGroup)

	return Router
}
