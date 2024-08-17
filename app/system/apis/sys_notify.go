package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/system/service"
	"go-admin/app/system/service/dto"
)

type SysNotify struct {
	api.Api
}

func (e SysNotify) NotifySend(c *gin.Context) {
	s := service.SysNotify{}
	req := dto.SysNotify{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 设置创建
	err = s.SendMessage(&req)
	if err != nil {
		e.Error(500, err, "发送通知失败")
		return
	}
	e.OK(nil, "通知发送成功")
}
