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
	routerCheckRole = append(routerCheckRole, registerOrderCategoryAuthRouter)
	//routerNoCheckRole = append(routerNoCheckRole, registerOrderRouter)
}

// 注册工单类别路由
func registerOrderCategoryAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.OrderCategory{}
	r := v1.Group("/flow/category").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/list", api.GetPage)
		r.GET("/list:id", api.Get)
		r.POST("/create", api.Insert)
		r.PUT("/update", api.Update)
		r.DELETE("/delete", api.Delete)
	}
}
