// @Author sunwenbo
// 2024/7/12 20:28
package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"smart-api/app/smart/apis"
	"smart-api/common/actions"
	"smart-api/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFlowManageAuthRouter)
}

// 注册工单模板路由
func registerFlowManageAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FlowManage{}
	r := v1.Group("/flow-manage").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.POST("", api.Insert)
		r.POST("/:id", api.Clone)
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}
}
