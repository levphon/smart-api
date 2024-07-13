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
	routerCheckRole = append(routerCheckRole, registerOrderAuthRouter)
	//routerNoCheckRole = append(routerNoCheckRole, registerOrderRouter)
}

// registerSysApiRouter 需要JWT认证
func registerOrderAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.OrderItems{}
	r := v1.Group("/order/items").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/list", api.GetPage)
		r.GET("/list:id", api.Get)
		r.POST("/create", api.Insert)
		r.PUT("/update", api.Update)
		r.DELETE("/delete", api.Delete)
	}
}

//// 注册路由 不进行JWT认证的
//func registerOrderRouter(v1 *gin.RouterGroup) {
//	api := apis.Test{}
//	r := v1.Group("/app/v1")
//	{
//		r.GET("/", api.Get)
//	}
//}
