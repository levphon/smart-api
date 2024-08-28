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
	routerCheckRole = append(routerCheckRole, registerExecMachineAuthRouter)
}

// 执行节点管理
func registerExecMachineAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.ExecMachine{}

	r := v1.Group("/machine").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		// 查询所有的任务
		r.GET("/list", api.GetPage)
		// 根据ID 查询
		r.GET("/list/:id", api.Get)
		// 创建任务
		r.POST("/create", api.Insert)
		// 更新任务信息
		r.PUT("/update", api.Update)
		// 删除任务
		r.DELETE("/delete", api.Delete)
	}
}
