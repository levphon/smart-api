// @Author sunwenbo
// 2024/7/12 20:28
package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"smart-api/app/smart/apis"
	"smart-api/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerOrderCategoryAuthRouter)
}

// 注册工单类别路由
func registerOrderCategoryAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.OrderCategory{}
	r := v1.Group("/order-category").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}
}
