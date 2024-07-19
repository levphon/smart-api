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
	routerCheckRole = append(routerCheckRole, registerFlowTemplatesAuthRouter)
}

// 注册工单模板路由
func registerFlowTemplatesAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FlowTemplates{}
	r := v1.Group("/flow/templates").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/list", api.GetPage)
		r.GET("/list:id", api.Get)
		r.POST("/create", api.Insert)
		r.PUT("/update", api.Update)
		r.DELETE("/delete", api.Delete)
	}
}
