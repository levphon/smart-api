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
	routerCheckRole = append(routerCheckRole, registerFlowManageAuthRouter)
}

// 注册工单模板路由
func registerFlowManageAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FlowManage{}
	r := v1.Group("/flowManage").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("/create", api.Insert)
		r.POST("/clone/:id", api.Clone)
		r.GET("/list", api.GetPage)
		r.GET("/details:id", api.Get)
		r.PUT("/update", api.Update)
		r.DELETE("/delete", api.Delete)
	}
}
