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

	r := v1.Group("/exec-machine").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		// 查询所有节点
		r.GET("", api.GetPage)
		// 根据ID 查询
		r.GET("/:id", api.Get)
		// 创建节点
		r.POST("", api.Insert)
		// 更新节点信息
		r.PUT("", api.Update)
		// 删除节点
		r.DELETE("", api.Delete)
		// 新增连接测试接口
		r.POST("/testConnection", api.TestConnection) // 添加连接测试接口

	}
}
