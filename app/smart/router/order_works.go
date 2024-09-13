// @Author sunwenbo
// 2024/7/12 20:28
package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/smart/apis"
	"go-admin/common/middleware"
	"go-admin/common/utils"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerOrderWorksAuthRouter)
}

// 注册工单类别路由
func registerOrderWorksAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.OrderWorksAPI{}
	hisApi := apis.OperationHistory{}
	notifyApi := apis.WorksNotify{}

	r := v1.Group("/order-works").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())

	wss := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())

	{
		// 查询所有的工单
		r.GET("", api.GetPage)
		// 根据ID 查询
		r.GET("/:id", api.Get)
		// 我的待办
		r.GET("/myBacklog", api.MyBacklogGet)
		// 我创建的
		r.GET("/createdByMe", api.CreatedByMe)
		// 与我相关的
		r.GET("/relatedToMe", api.RelatedToMe)
		// 创建工单
		r.POST("", api.Insert)
		// 更新工单信息、生成更新记录
		r.PUT("", api.Update)
		// 工单处理,更新工单当前节点、当前处理人
		r.PUT("/handle", api.Handle)
		// 删除工单
		r.DELETE("", api.Delete)
	}
	{
		// 获取历史变更信息
		r.GET("/history", hisApi.GetOperationHistory)
	}
	{
		// 获取通知消息
		r.GET("/notify", notifyApi.GetNotify)

		// 发送通知消息
		r.POST("/notify", notifyApi.CreateNotify)

		// 处理通知消息
		r.PUT("/notify", notifyApi.UpdateNotify)

	}

	{
		// 自定义WebSocket处理程序
		wss.GET("/ws/:id/task", utils.WsHandler)
	}

}
