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

type OrderWorks struct {
	api.Api
}

type OperationHistory struct {
	api.Api
}

func (e OrderWorks) GetPage(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	s := service.OrderWorks{}
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
	var objects []models.OrderWorks
	var total int64

	err = s.GetOrderWorksPage(pageNum, limit, &objects, &total)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(objects, int(total), pageNum, limit, "查询成功")
}
func (e OrderWorks) Get(c *gin.Context) {
	s := service.OrderWorks{}
	// 请求结构体 id
	req := dto.OrderWorksGetReq{}
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
	var object models.OrderWorks

	// &object 因为是传递的地址，object 地址对应的内存值会被修改，所以就不需要接收返回值了
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}
func (e OrderWorks) MyBacklogGet(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	s := service.OrderWorks{}
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
	var objects []models.OrderWorks

	userId := user.GetUserId(c)

	err = s.MyBackGet(pageNum, limit, &objects, userId)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(objects, "查询成功")

}
func (e OrderWorks) CreatedByMe(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	s := service.OrderWorks{}
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
	var objects []models.OrderWorks

	userId := user.GetUserId(c)

	err = s.CreatedByMe(pageNum, limit, &objects, userId)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(objects, "查询成功")

}
func (e OrderWorks) RelatedToMe(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	s := service.OrderWorks{}
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
	var objects []models.OrderWorks

	userId := user.GetUserId(c)

	err = s.RelatedToMe(pageNum, limit, &objects, userId)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(objects, "查询成功")

}
func (e OrderWorks) Insert(c *gin.Context) {
	s := service.OrderWorks{}

	req := dto.OrderWorksInsertReq{}
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

// 更新OrderWorks
func (e OrderWorks) Update(c *gin.Context) {
	s := service.OrderWorks{}
	req := dto.OrderWorksUpdateReq{}
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

// 工单处理
func (e OrderWorks) Handle(c *gin.Context) {
	s := service.OrderWorks{}
	req := dto.OrderWorksHandleReq{}
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
	handle := user.GetUserId(c)
	go func() {
		err := s.Handle(&req, handle)
		if err != nil {

		}
	}()

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "操作成功")
}

// 删除OrderWorks
func (e OrderWorks) Delete(c *gin.Context) {
	s := service.OrderWorks{}

	req := dto.OrderWorksDeleteReq{}
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

// 获取工单历史
func (h OperationHistory) GetOperationHistory(c *gin.Context) {
	// 分页查询返回pageNum和limit
	pageNum, limit := global.PagingQuery(c)
	title := c.Query("title")

	s := service.OperationHistory{}
	// 创建数据库连接和绑定请求
	err := h.MakeContext(c).
		MakeOrm().
		Bind(binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		h.Logger.Error(err)
		h.Error(500, err, err.Error())
		return
	}
	// 定义一个存储所有工单类别数据的切片
	var objects []models.OperationHistory

	err = s.GetHistory(pageNum, limit, &objects, title)
	if err != nil {
		h.Error(500, err, "查询失败")
		return
	}

	h.OK(objects, "查询成功")

}
