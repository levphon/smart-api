// @Author sunwenbo
// 2024/7/19 17:58
package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/smart/models"
	"go-admin/app/smart/service"
	"go-admin/app/smart/service/dto"
	"go-admin/common/global"
)

type WorksNotify struct {
	api.Api
}

func (n WorksNotify) SendNotify(c *gin.Context) {
	s := service.WorksNotify{}

	req := dto.WorksNotifyReq{}
	err := n.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		n.Logger.Error(err)
		n.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		n.Error(500, err, fmt.Sprintf("创建失败 err:%v", err))
		return
	}
	n.OK(req.GetId(), fmt.Sprintf("%v 创建成功", req.Title))
}

func (n WorksNotify) GetNotify(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	s := service.WorksNotify{}
	// 创建数据库连接和绑定请求
	err := n.MakeContext(c).
		MakeOrm().
		Bind(binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		n.Logger.Error(err)
		n.Error(500, err, err.Error())
		return
	}

	// 定义一个存储所有工单类别数据的切片
	var objects []models.WorksNotify

	userId := user.GetUserId(c)

	err = s.GetNotify(pageNum, limit, &objects, userId)
	if err != nil {
		n.Error(500, err, "查询失败")
		return
	}

	n.OK(objects, "查询成功")

}

func (n WorksNotify) UpdateNotify(c *gin.Context) {
	s := service.WorksNotify{}
	req := dto.WorksNotifyUpdateReq{}
	err := n.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors

	if err != nil {
		n.Logger.Error(err)
		n.Error(500, err, fmt.Sprintf("更新失败 err:%v", err))
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.UpdateNotify(&req)
	if err != nil {
		n.Error(500, err, err.Error())
		return
	}
	n.OK(req.GetId(), "更新成功")
}
