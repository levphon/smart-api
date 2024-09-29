// @Author sunwenbo
// 2024/7/13 21:12
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

type OrderItems struct {
	api.Api
}

// GetPage 获取OrderItems列表
// @Summary 获取OrderItems列表
// @Description 获取OrderItems列表
// @Tags 工单项目
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OrderItems}} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-items [get]
// @Security Bearer
func (e OrderItems) GetPage(c *gin.Context) {

	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)

	s := service.OrderItems{}
	// 创建数据库连接和绑定请求
	err := e.MakeContext(c).
		MakeOrm().
		Bind(binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 定义一个存储所有订单项数据的切片
	var objects []models.OrderItems
	var total int64

	err = s.GetOrderItemsPage(pageNum, limit, &objects, &total)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(objects, int(total), pageNum, limit, "查询成功")
}

// Get 获取OrderItems
// @Summary 获取OrderItems
// @Description 获取OrderItems
// @Tags 工单项目
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrderItems} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-items/{id} [get]
// @Security Bearer
func (e OrderItems) Get(c *gin.Context) {
	s := service.OrderItems{}
	req := dto.OrderItemsGetReq{}

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
	var object models.OrderItems

	// 调用修改后的 Get 方法
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建OrderItems
// @Summary 创建OrderItems
// @Description 创建OrderItems
// @Tags 工单项目
// @Accept application/json
// @Product application/json
// @Param data body dto.OrderItemsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/order-items [post]
// @Security Bearer
func (e OrderItems) Insert(c *gin.Context) {
	s := service.OrderItems{}

	req := dto.OrderItemsInsertReq{}
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
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	req.Creator = user.GetUserName(c)
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建失败 err:%v", err))
		return
	}
	e.OK(req.GetId(), fmt.Sprintf("%v 创建成功", req.Title))
}

// Update 修改OrderItems
// @Summary 修改OrderItems
// @Description 修改OrderItems
// @Tags 工单项目
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrderItemsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/order-items/{id} [put]
// @Security Bearer
func (e OrderItems) Update(c *gin.Context) {
	s := service.OrderItems{}
	req := dto.OrderItemsUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf("更新失败 err:%v", err))
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	req.Regenerator = user.GetUserName(c)
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 删除OrderItems
// @Summary 删除OrderItems
// @Description 删除OrderItems
// @Tags 工单项目
// @Param data body dto.OrderItemsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/order-items [delete]
// @Security Bearer
func (e OrderItems) Delete(c *gin.Context) {
	s := service.OrderItems{}
	req := dto.OrderItemsDeleteReq{}
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

	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除失败 err:%v", err))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
