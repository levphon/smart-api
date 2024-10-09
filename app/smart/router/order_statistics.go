// @Author sunwenbo
// 2024/9/25 20:28
package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"smart-api/app/smart/apis"
	"smart-api/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerOrderStatisticsAuthRouter)
}

// registerSysApiRouter 需要JWT认证
func registerOrderStatisticsAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.OrderStatistics{}
	r := v1.Group("/statistics").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/orders", api.GetOrderStatistics)          // 获取工单统计
		r.GET("/orders/count", api.GetOrderCountByPeriod) // 根据周或月统计工单数量，以及个人提交排行榜

		r.GET("/ratings", api.GetOrderRatingsByPeriod) // 根据周或月统计获取评分统计
	}
}
