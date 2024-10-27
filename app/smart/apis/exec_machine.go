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

type ExecMachine struct {
	api.Api
}

type ExecutionHistory struct {
	api.Api
}

// GetPage 获取ExecMachine列表
// @Summary 获取ExecMachine列表
// @Description 获取ExecMachine列表
// @Tags 执行节点管理
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ExecMachine}} "{"code": 200, "data": [...]}"
// @Router /api/v1/exec-machine [get]
// @Security Bearer
func (e ExecMachine) GetPage(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	s := service.ExecMachine{}
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

	// 定义一个存储所有执行节点数据的切片
	var objects []models.ExecMachine
	var total int64

	err = s.GetExecMachinePage(pageNum, limit, &objects, &total)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查询失败: %v", err))
		return
	}

	e.PageOK(objects, int(total), pageNum, limit, "查询成功")
}

// Get 获取ExecMachine
// @Summary 获取ExecMachine
// @Description 获取ExecMachine
// @Tags 执行节点管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ExecMachine} "{"code": 200, "data": [...]}"
// @Router /api/v1/exec-machine/{id} [get]
// @Security Bearer
func (e ExecMachine) Get(c *gin.Context) {
	s := service.ExecMachine{}
	// 请求结构体 id
	req := dto.ExecMachineGetReq{}
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
	var object models.ExecMachine

	// &object 因为是传递的地址，object 地址对应的内存值会被修改，所以就不需要接收返回值了
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查询失败: %v", err))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建ExecMachine
// @Summary 创建ExecMachine
// @Description 创建ExecMachine
// @Tags 执行节点管理
// @Accept application/json
// @Product application/json
// @Param data body dto.ExecMachineInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/exec-machine [post]
// @Security Bearer
func (e ExecMachine) Insert(c *gin.Context) {
	s := service.ExecMachine{}

	req := dto.ExecMachineInsertReq{}
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
	e.OK(req.GetId(), fmt.Sprintf("%v 创建成功", req.HostName))
}

// Update 修改ExecMachine
// @Summary 修改ExecMachine
// @Description 修改ExecMachine
// @Tags 执行节点管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ExecMachineUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/exec-machine/{id} [put]
// @Security Bearer
func (e ExecMachine) Update(c *gin.Context) {
	s := service.ExecMachine{}
	req := dto.ExecMachineUpdateReq{}
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

// Delete 删除ExecMachine
// @Summary 删除ExecMachine
// @Description 删除ExecMachine
// @Tags 执行节点管理
// @Param data body dto.ExecMachineDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/exec-machine [delete]
// @Security Bearer
func (e ExecMachine) Delete(c *gin.Context) {
	s := service.ExecMachine{}

	req := dto.ExecMachineDeleteReq{}
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

// TestConnection ExecMachine
// @Summary 测试连接ExecMachine
// @Description 测试连接ExecMachine
// @Tags 历史任务
// @Param data body dto.ExecMachineGetReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "连接成功"}"
// @Router /api/v1/exec-machine/testConnection [post]
// @Security Bearer
func (e *ExecMachine) TestConnection(c *gin.Context) {
	s := service.ExecMachine{}
	req := dto.ExecMachineGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// &object 因为是传递的地址，object 地址对应的内存值会被修改，所以就不需要接收返回值了
	err = s.TestConn(&req)
	// 根据测试结果返回状态信息
	if err != nil {
		e.Error(500, err, fmt.Sprintf("连接失败 err:%v", err.Error()))
	} else {
		e.OK(req.GetId(), fmt.Sprintf("测试连接成功"))
	}

}
