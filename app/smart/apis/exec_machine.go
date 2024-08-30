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

type ExecMachine struct {
	api.Api
}

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

// TestConnection 测试机器连接是否成功
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
	var object models.ExecMachine

	// &object 因为是传递的地址，object 地址对应的内存值会被修改，所以就不需要接收返回值了
	err = s.TestConn(&req, &object)
	// 根据测试结果返回状态信息
	if err != nil {
		e.Error(500, err, fmt.Sprintf("连接失败 err:%v", err.Error()))
	} else {
		e.OK(req.GetId(), fmt.Sprintf("测试连接成功"))
	}

}
