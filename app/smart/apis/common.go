// @Author sunwenbo
// 2024/7/19 17:58
package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"smart-api/app/smart/models"
	"smart-api/app/smart/service"
	"smart-api/app/smart/service/dto"
	"smart-api/common/global"
)

type WorksNotify struct {
	api.Api
}

// CreateNotify  创建OperationHistory
// @Summary 创建OperationHistory
// @Description 创建OperationHistory
// @Tags 操作历史
// @Accept application/json
// @Product application/json
// @Param data body dto.WorksNotifyReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/order/works/notify [post]
// @Security Bearer
func (n WorksNotify) CreateNotify(c *gin.Context) {
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

	go func() {
		err := s.Insert(&req)
		if err != nil {

		}
	}()
	if err != nil {
		n.Error(500, err, fmt.Sprintf("创建失败 err:%v", err))
		return
	}
	n.OK(req.GetId(), fmt.Sprintf("%v 创建成功", req.Title))
}

// GetNotify  获取OperationHistory
// @Summary 获取OperationHistory
// @Description 获取OperationHistory
// @Tags 操作历史
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OperationHistory} "{"code": 200, "data": [...]}"
// @Router /api/v1/order/works/notify [get]
// @Security Bearer
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

	userid := user.GetUserId(c)
	err = s.GetNotify(pageNum, limit, &objects, userid)
	if err != nil {
		n.Error(500, err, "查询失败")
		return
	}

	n.OK(objects, "查询成功")

}

// UpdateNotify  修改OperationHistory
// @Summary 修改OperationHistory
// @Description 修改OperationHistory
// @Tags 操作历史
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.WorksNotifyUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/order/works/notify [put]
// @Security Bearer
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

// GetOperationHistory  获取OperationHistory列表
// @Summary 获取OperationHistory列表
// @Description 获取OperationHistory列表
// @Tags 操作历史
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OperationHistory}} "{"code": 200, "data": [...]}"
// @Router /api/v1/order/works/history [get]
// @Security Bearer
func (h OperationHistory) GetOperationHistory(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	title := c.Query("title")

	s := service.OperationHistory{}
	// 创建数据库连接和绑定请求
	err := h.MakeContext(c).
		MakeOrm().
		Bind(binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		h.Logger.Error(err)
		h.Error(500, err, err.Error())
		return
	}
	// 定义一个存储所有工单类别数据的切片
	var objects []models.OperationHistory

	err = s.GetHistory(pageNum, limit, &objects, title)
	if err != nil {
		h.Error(500, err, "查询失败")
		return
	}

	h.OK(objects, "查询成功")

}
