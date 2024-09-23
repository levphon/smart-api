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
// Update 更新留言
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
