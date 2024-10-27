// @Author sunwenbo
// 2024/7/13 21:12
package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"net/http"
	"smart-api/app/smart/models"
	"smart-api/app/smart/service"
	"smart-api/app/smart/service/dto"
	"smart-api/common/global"
)

type OrderTask struct {
	api.Api
}

// GetPage 获取OrderTask列表
// @Summary 获取OrderTask列表
// @Description 获取OrderTask列表
// @Tags 工单任务
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OrderTask}} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-task [get]
// @Security Bearer
func (e OrderTask) GetPage(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	s := service.OrderTask{}
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
	var objects []models.OrderTask
	var total int64

	err = s.GetOrderTaskPage(pageNum, limit, &objects, &total)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(objects, int(total), pageNum, limit, "查询成功")
}

// Get 获取OrderTask
// @Summary 获取OrderTask
// @Description 获取OrderTask
// @Tags 工单任务
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrderTask} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-task/{id} [get]
// @Security Bearer
func (e OrderTask) Get(c *gin.Context) {
	s := service.OrderTask{}
	// 请求结构体 id
	req := dto.OrderTaskGetReq{}
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
	var object models.OrderTask

	// &object 因为是传递的地址，object 地址对应的内存值会被修改，所以就不需要接收返回值了
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建OrderTask
// @Summary 创建OrderTask
// @Description 创建OrderTask
// @Tags 工单任务
// @Accept application/json
// @Product application/json
// @Param data body dto.OrderTaskInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/order-task [post]
// @Security Bearer
func (e OrderTask) Insert(c *gin.Context) {
	s := service.OrderTask{}

	req := dto.OrderTaskInsertReq{}
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

// Update 修改OrderTask
// @Summary 修改OrderTask
// @Description 修改OrderTask
// @Tags 工单任务
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrderTaskUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/order-task/{id} [put]
// @Security Bearer
func (e OrderTask) Update(c *gin.Context) {
	s := service.OrderTask{}
	req := dto.OrderTaskUpdateReq{}
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

// Delete 删除OrderTask
// @Summary 删除OrderTask
// @Description 删除OrderTask
// @Tags 工单任务
// @Param data body dto.OrderTaskDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/order-task [delete]
// @Security Bearer
func (e OrderTask) Delete(c *gin.Context) {
	s := service.OrderTask{}

	req := dto.OrderTaskDeleteReq{}
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

// GetHistoryTaskPage
// @Summary 获取历史任务列表
// @Description 获取ExecutionTaskHistory列表
// @Tags 执行节点管理
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ExecutionHistoryTask}} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-task/history [get]
// @Security Bearer
func (e OrderTask) GetHistoryTaskPage(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	s := service.OrderTask{}
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

	// 定义一个存储所有历史任务数据的切片
	var objects []models.ExecutionHistoryTask
	var total int64

	err = s.GetHistoryTask(pageNum, limit, &objects, &total)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查询失败: %v", err))
		return
	}

	e.PageOK(objects, int(total), pageNum, limit, "查询成功")
}

// DeleteHistoryTask 删除历史执行任务
// @Summary 删除历史执行任务
// @Description 删除历史执行任务
// @Tags 历史任务
// @Param data body dto.OrderHistoryTaskDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/order-task/history [delete]
// @Security Bearer
func (e OrderTask) DeleteHistoryTask(c *gin.Context) {
	s := service.OrderTask{}

	req := dto.OrderHistoryTaskDeleteReq{}
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
	err = s.DeleteHistoryTask(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除失败 err:%v", err))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// GetTaskLogsByID 获取历史任务日志数据
// @Summary 获取历史任务日志数据
// @Description 根据任务 ID 获取对应的历史任务日志内容
// @Tags 历史任务
// @Param id path string true "任务 ID"
// @Success 200 {object} response.Response "{"code": 200, "message": "获取任务日志成功", "data": "日志内容"}"
// @Failure 500 {object} response.Response "{"code": 500, "message": "获取任务日志失败"}"
// @Router /api/v1/order-task/history/logs/{id} [get]
// @Security Bearer
func (e *OrderTask) GetTaskLogsByID(c *gin.Context) {
	// 从 URL 中获取任务 ID
	hisTaskID := c.Param("id")

	// 调用 service 层获取日志文件内容
	s := service.OrderTask{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, "初始化服务失败")
		return
	}

	// 获取日志文件内容
	logContent, err := s.GetTaskLogFileContent(hisTaskID)
	if err != nil {
		e.Error(http.StatusInternalServerError, err, "获取任务日志失败")
		return
	}

	// 返回日志文件内容
	e.OK(logContent, "获取任务日志成功")
}
