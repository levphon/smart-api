package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/smart/service"
	"go-admin/app/smart/service/dto"
)

type OrderCommentAPI struct {
	api.Api
}

// Get 获取工单留言
// @Summary 获取工单留言
// @Description 根据工单ID获取所有留言
// @Tags 工单留言
// @Param orderId path int true "工单ID"
// @Success 200 {object} response.Response{data=[]dto.OrderCommentGetReq} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-comment/{orderId} [get]
// @Security Bearer
func (e OrderCommentAPI) Get(c *gin.Context) {
	s := service.OrderComment{}
	req := dto.OrderCommentGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 获取工单的所有留言
	comments, err := s.GetCommentsByOrderID(req.OrderID)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(comments, "查询成功")
}

// Insert 添加留言
// @Summary 添加留言
// @Description 为指定工单添加留言
// @Tags 工单留言
// @Accept application/json
// @Product application/json
// @Param data body dto.OrderCommentInsertReq true "留言数据"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/order-comment [post]
// @Security Bearer
func (e OrderCommentAPI) Insert(c *gin.Context) {
	s := service.OrderComment{}
	req := dto.OrderCommentInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "绑定请求失败")
		return
	}
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, "添加失败")
		return
	}
	e.OK(req.GetId(), "添加成功")
}

// Update 更新留言
// @Summary 更新留言
// @Description 根据留言ID更新留言内容
// @Tags 工单留言
// @Accept application/json
// @Product application/json
// @Param id path int true "留言ID"
// @Param data body dto.OrderCommentUpdateReq true "更新数据"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/order-comment/{id} [put]
// @Security Bearer
func (e OrderCommentAPI) Update(c *gin.Context) {
	s := service.OrderComment{}
	req := dto.OrderCommentUpdateReq{}

	// 绑定请求
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "绑定请求失败")
		return
	}

	// 更新留言
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, "更新失败")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 删除留言
// @Summary 删除留言
// @Description 根据留言ID删除留言
// @Tags 工单留言
// @Param data body dto.OrderCommentDeleteReq true "删除请求"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/order-comment [delete]
// @Security Bearer
func (e OrderCommentAPI) Delete(c *gin.Context) {
	s := service.OrderComment{}
	req := dto.OrderCommentDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "绑定请求失败")
		return
	}
	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}
