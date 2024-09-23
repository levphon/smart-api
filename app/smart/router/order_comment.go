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
	routerCheckRole = append(routerCheckRole, registerOrderCommentAuthRouter)
}

// 注册工单评论路由
func registerOrderCommentAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	commentApi := apis.OrderCommentAPI{}

	r := v1.Group("/order-comments").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		// 根据ID获取评论
		r.GET("/:orderID", commentApi.Get) // 根据工单ID获取所有留言
		// 添加留言
		r.POST("", commentApi.Insert) // 添加留言的路由
		// 更新留言
		r.PUT("", commentApi.Update) // 更新留言的路由
		// 删除留言
		r.DELETE("", commentApi.Delete) // 删除留言的路由
	}
}
