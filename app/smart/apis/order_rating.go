package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/smart/service"
	"go-admin/app/smart/service/dto"
)

type OrderRatingAPI struct {
	api.Api
}

// Insert 添加评分
func (e OrderRatingAPI) Insert(c *gin.Context) {
	s := service.OrderRating{}
	req := dto.OrderRatingInsertReq{}
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
		e.Error(500, err, "评分失败")
		return
	}
	e.OK(req.GetId(), "评分成功")
}

// Update 更新评分
func (e OrderRatingAPI) Update(c *gin.Context) {
	s := service.OrderRating{}
	req := dto.OrderRatingUpdateReq{}
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
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, "更新失败")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 删除评分
func (e OrderRatingAPI) Delete(c *gin.Context) {
	s := service.OrderRating{}
	req := dto.OrderRatingDeleteReq{}
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
