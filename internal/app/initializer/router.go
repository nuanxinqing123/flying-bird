package initializer

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nuanxinqing123/flying-bird/internal/controller"
	"github.com/nuanxinqing123/flying-bird/internal/middleware"
	ginprometheus "github.com/zsais/go-gin-prometheus"
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

	// 初始化 Prometheus 中间件
	p := ginprometheus.NewPrometheus("gin")
	p.Use(Router)

	// 存活检测
	Router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 健康检查
	HealthyGroup := Router.Group("/")
	HealthyCon := controller.NewHealthyController()
	HealthyCon.HealthyRouter(HealthyGroup)

	// 用户管理户管理
	UserGroup := Router.Group("/user")
	UserCon := controller.NewUserController()
	UserCon.UserRouter(UserGroup)

	return Router
}
