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

// Get 获取工单评分
// @Summary 获取工单评分
// @Description 根据工单ID获取所有评分
// @Tags 工单评分
// @Param orderId path int true "工单ID"
// @Success 200 {object} response.Response{data=[]models.OrderRating} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-rating/{orderId} [get]
// @Security Bearer
func (e OrderRatingAPI) Get(c *gin.Context) {
	s := service.OrderRating{}
	req := dto.OrderRatingGetReq{}
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

	// 获取工单评分
	rating, err := s.GetRatingByOrderID(req.OrderId)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	if rating == nil {
		// 如果没有评分，返回 200 和自定义消息
		e.OK(nil, "该工单没有进行评分或者还没有到结束节点")
		return
	}

	e.OK(rating, "查询成功")
}

// Insert 添加评分
// @Summary 添加评分
// @Description 为指定工单添加评分
// @Tags 工单评分
// @Accept application/json
// @Product application/json
// @Param data body dto.OrderRatingInsertReq true "评分数据"
// @Success 200 {object} response.Response	"{"code": 200, "message": "评分成功"}"
// @Router /api/v1/order-rating [post]
// @Security Bearer
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
// @Summary 更新评分
// @Description 根据评分ID更新评分信息
// @Tags 工单评分
// @Accept application/json
// @Product application/json
// @Param id path int true "评分ID"
// @Param data body dto.OrderRatingUpdateReq true "更新数据"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/order-rating/{id} [put]
// @Security Bearer
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
// @Summary 删除评分
// @Description 根据评分ID删除评分
// @Tags 工单评分
// @Param data body dto.OrderRatingDeleteReq true "删除请求"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/order-rating [delete]
// @Security Bearer
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
