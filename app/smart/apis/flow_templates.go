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

type FlowTemplates struct {
	api.Api
}

// GetPage 获取FlowTemplates列表
// @Summary 获取FlowTemplates列表
// @Description 获取FlowTemplates列表
// @Tags 流程模板管理
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FlowTemplates}} "{"code": 200, "data": [...]}"
// @Router /api/v1/flow-templates [get]
// @Security Bearer
func (e FlowTemplates) GetPage(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)

	s := service.FlowTemplates{}
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
	var objects []models.FlowTemplates
	var total int64

	err = s.GetFlowTemplatesPage(pageNum, limit, &objects, &total)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(objects, int(total), pageNum, limit, "查询成功")
}

// Get 获取FlowTemplates
// @Summary 获取FlowTemplates
// @Description 获取FlowTemplates
// @Tags 流程模板管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FlowTemplates} "{"code": 200, "data": [...]}"
// @Router /api/v1/flow-templates/{id} [get]
// @Security Bearer
func (e FlowTemplates) Get(c *gin.Context) {
	s := service.FlowTemplates{}
	// 请求结构体 id
	req := dto.FlowTemplatesGetReq{}
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
	var object models.FlowTemplates

	// &object 因为是传递的地址，object 地址对应的内存值会被修改，所以就不需要接收返回值了
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建FlowTemplates
// @Summary 创建FlowTemplates
// @Description 创建FlowTemplates
// @Tags 流程模板管理
// @Accept application/json
// @Product application/json
// @Param data body dto.FlowTemplatesInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/flow-templates [post]
// @Security Bearer
func (e FlowTemplates) Insert(c *gin.Context) {
	s := service.FlowTemplates{}

	req := dto.FlowTemplatesInsertReq{}
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

// Update 修改FlowTemplates
// @Summary 修改FlowTemplates
// @Description 修改FlowTemplates
// @Tags 流程模板管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FlowTemplatesUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/flow-templates/{id} [put]
// @Security Bearer
func (e FlowTemplates) Update(c *gin.Context) {
	s := service.FlowTemplates{}
	req := dto.FlowTemplatesUpdateReq{}
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

// Delete 删除FlowTemplates
// @Summary 删除FlowTemplates
// @Description 删除FlowTemplates
// @Tags 流程模板管理
// @Param data body dto.FlowTemplatesDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/flow-templates [delete]
// @Security Bearer
func (e FlowTemplates) Delete(c *gin.Context) {
	s := service.FlowTemplates{}
	req := dto.FlowTemplatesDeleteReq{}
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
