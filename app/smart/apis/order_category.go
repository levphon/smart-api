// @Author sunwenbo
// 2024/7/13 21:12
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

type OrderCategory struct {
	api.Api
}

// GetPage 获取OrderCategory列表
// @Summary 获取OrderCategory列表
// @Description 获取OrderCategory列表
// @Tags 工单分类
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OrderCategory}} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-category/ [get]
// @Security Bearer
func (e OrderCategory) GetPage(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)

	s := service.OrderCategory{}
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

	// 定义一个存储所有工单类别数据的切片
	var objects []models.OrderCategory
	var total int64

	err = s.GetOrderCategoryPage(pageNum, limit, &objects, &total)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(objects, int(total), pageNum, limit, "查询成功")
}

// Get 获取OrderCategory
// @Summary 获取OrderCategory
// @Description 获取OrderCategory
// @Tags 工单分类
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrderCategory} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-category/{id} [get]
// @Security Bearer
func (e OrderCategory) Get(c *gin.Context) {
	s := service.OrderCategory{}
	// 请求结构体 id
	req := dto.OrderCategoryGetReq{}
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
	var object models.OrderCategory

	// &object 因为是传递的地址，object 地址对应的内存值会被修改，所以就不需要接收返回值了
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建OrderCategory
// @Summary 创建OrderCategory
// @Description 创建OrderCategory
// @Tags 工单分类
// @Accept application/json
// @Product application/json
// @Param data body dto.OrderCategoryInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/order-category [post]
// @Security Bearer
func (e OrderCategory) Insert(c *gin.Context) {
	s := service.OrderCategory{}

	req := dto.OrderCategoryInsertReq{}
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
	e.OK(req.GetId(), fmt.Sprintf("%v 创建成功", req.Name))
}

// Update 修改OrderCategory
// @Summary 修改OrderCategory
// @Description 修改OrderCategory
// @Tags 工单分类
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrderCategoryUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/order-category/{id} [put]
// @Security Bearer
func (e OrderCategory) Update(c *gin.Context) {
	s := service.OrderCategory{}
	req := dto.OrderCategoryUpdateReq{}
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

// Delete 删除OrderCategory
// @Summary 删除OrderCategory
// @Description 删除OrderCategory
// @Tags 工单分类
// @Param data body dto.OrderCategoryDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/order-category [delete]
// @Security Bearer
func (e OrderCategory) Delete(c *gin.Context) {
	s := service.OrderCategory{}

	req := dto.OrderCategoryDeleteReq{}
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
