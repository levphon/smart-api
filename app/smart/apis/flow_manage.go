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

type FlowManage struct {
	api.Api
}

// GetPage 获取FlowManage列表
// @Summary 获取FlowManage列表
// @Description 获取FlowManage列表
// @Tags 工单流程管理
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FlowManage}} "{"code": 200, "data": [...]}"
// @Router /api/v1/flow-manage [get]
// @Security Bearer
func (e FlowManage) GetPage(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	s := service.FlowManage{}
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
	var objects []models.FlowManage
	var total int64
	err = s.GetFlowManagePage(pageNum, limit, &objects, &total)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(objects, int(total), pageNum, limit, "查询成功")
}

// Get 获取FlowManage
// @Summary 获取FlowManage
// @Description 获取FlowManage
// @Tags 工单流程管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FlowManage} "{"code": 200, "data": [...]}"
// @Router /api/v1/flow-manage/{id} [get]
// @Security Bearer
func (e FlowManage) Get(c *gin.Context) {
	s := service.FlowManage{}
	// 请求结构体 id
	req := dto.FlowManageGetReq{}
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
	var object models.FlowManage

	// &object 因为是传递的地址，object 地址对应的内存值会被修改，所以就不需要接收返回值了
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建FlowManage
// @Summary 创建FlowManage
// @Description 创建FlowManage
// @Tags 工单流程管理
// @Accept application/json
// @Product application/json
// @Param data body dto.FlowManageInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/flow-manage [post]
// @Security Bearer
func (e FlowManage) Insert(c *gin.Context) {
	s := service.FlowManage{}

	req := dto.FlowManageInsertReq{}
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

// Clone 克隆FlowManage
// @Summary 克隆FlowManage
// @Description 克隆FlowManage
// @Tags 工单流程管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FlowManageUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "克隆成功"}"
// @Router /api/v1/flow-manage/clone/{id} [post]
// @Security Bearer
func (e FlowManage) Clone(c *gin.Context) {
	s := service.FlowManage{}
	req := dto.FlowManageInsertReq{}

	// 获取要克隆的流程ID
	id := c.Param("id")
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
	err = s.Clone(id, c)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("克隆失败 err:%v", err))
		return
	}
	e.OK(id, "克隆成功")
}

// Update 修改FlowManage
// @Summary 修改FlowManage
// @Description 修改FlowManage
// @Tags 工单流程管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FlowManageUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/flow-manage/{id} [put]
// @Security Bearer
func (e FlowManage) Update(c *gin.Context) {
	s := service.FlowManage{}
	req := dto.FlowManageUpdateReq{}
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

// Delete 删除FlowManage
// @Summary 删除FlowManage
// @Description 删除FlowManage
// @Tags 工单流程管理
// @Param data body dto.FlowManageDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/flow-manage [delete]
// @Security Bearer
func (e FlowManage) Delete(c *gin.Context) {
	s := service.FlowManage{}

	req := dto.FlowManageDeleteReq{}
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
