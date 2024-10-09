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
	routerCheckRole = append(routerCheckRole, registerUserFavoriteAuthRouter)
}

// registerSysApiRouter 需要JWT认证
func registerUserFavoriteAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.UserFavorites{}
	r := v1.Group("/favorite").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("", api.AddFavorite)
		r.DELETE("", api.RemoveFavorite)
		r.GET("", api.GetUserFavorites)
	}
}
