package router

import (
	"mime"
	"smart-api/app/system/apis"

	"github.com/go-admin-team/go-admin-core/sdk/config"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/ws"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"

	"smart-api/common/middleware"
	"smart-api/common/middleware/handler"
	_ "smart-api/docs/smart"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	sysBaseRouter(g)
	// 静态文件
	sysStaticFileRouter(g)
	// swagger；注意：生产环境可以注释掉
	if config.ApplicationConfig.Mode != "prod" {
		sysSwaggerRouter(g)
	}
	// 需要认证
	sysCheckRoleRouterInit(g, authMiddleware)

	return g
}

func sysBaseRouter(r *gin.RouterGroup) {
	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendAllService()

	if config.ApplicationConfig.Mode != "prod" {
		r.GET("/", apis.GoAdmin)
	}
	r.GET("/info", handler.Ping)
}

func sysStaticFileRouter(r *gin.RouterGroup) {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		return
	}
	r.Static("/static", "./static")
	// 如果不是prod环境，则可以访问代码生成接口
	if config.ApplicationConfig.Mode != "prod" {
		r.Static("/form-generator", "./static/form-generator")
	}
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	// r.GET("/swagger/smart/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("smart")))
	r.GET("/swagger/smart/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("smart")))

}

func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	wss := r.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		wss.GET("/ws/:id/:channel", ws.WebsocketManager.WsClient)
		wss.GET("/wslogout/:id/:channel", ws.WebsocketManager.UnWsClient)
	}

	// 设置/api/v1组，在这个组内添加路由不用重复指定/api/v1
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)

		// Refresh time can be longer than token timeout
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	registerBaseRouter(v1, authMiddleware)
}
func InitLDAPRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	ldapRouter := r.Group("/api/v1")
	{
		ldapRouter.POST("/ldap/login", authMiddleware.LoginHandler)
	}
}

// 注册基础的路由，在服务启动时进行
func registerBaseRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysMenu{}
	api2 := apis.SysDept{}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		v1auth.GET("/roleMenuTreeselect/:roleId", api.GetMenuTreeSelect)
		//v1.GET("/menuTreeselect", api.GetMenuTreeSelect)
		v1auth.GET("/roleDeptTreeselect/:roleId", api2.GetDeptTreeRoleSelect)
		v1auth.POST("/logout", handler.LogOut)
	}
}
