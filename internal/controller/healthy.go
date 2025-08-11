package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nuanxinqing123/flying-bird/internal/service"
	res "github.com/nuanxinqing123/flying-bird/pkg/response"
)

type HealthyController struct{}

// HealthyRouter 注册路由
func (c *HealthyController) HealthyRouter(r *gin.RouterGroup) {
	// 登录
	r.GET("/healthy", c.Healthy)
}

// Healthy 健康检查
func (c *HealthyController) Healthy(ctx *gin.Context) {
	// 业务处理
	resCode, msg := service.CheckHealth(ctx)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}
