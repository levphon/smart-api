// @Author sunwenbo
// 2024/7/12 20:28
package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/smart/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerOrderRatingAuthRouter)
}

// 注册工单评论路由
func registerOrderRatingAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	ratingApi := apis.OrderRatingAPI{} // 假设你有一个处理评分的 API

	rating := v1.Group("/order-ratings").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		// 获取工单评分
		rating.GET("/:orderID", ratingApi.Get) // 添加评分的路由
		// 添加评分
		rating.POST("", ratingApi.Insert) // 添加评分的路由
		// 更新评分
		rating.PUT("", ratingApi.Update) // 更新评分的路由
		// 删除评分
		rating.DELETE("", ratingApi.Delete) // 删除评分的路由
	}
}
