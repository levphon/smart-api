package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"smart-api/app/system/apis"
	"smart-api/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysNotifyRouter)
}

// 需认证的路由代码
func registerSysNotifyRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysNotify{}

	r := v1.Group("/notify").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("", api.NotifySend)
	}
}
