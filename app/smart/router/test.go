// @Author sunwenbo
// 2024/7/12 20:28
package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/smart/apis"
)

func init() {
	//routerCheckRole = append(routerCheckRole, registerTestAuthRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerTestAuthRouter)
}

// registerSysApiRouter 需要JWT认证
func registerTestAuthRouter(v1 *gin.RouterGroup) {
	api := apis.Test{}
	r := v1.Group("/hello")
	{
		r.GET("/", api.Get)
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
