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
	"strconv"
)

type OrderItems struct {
	api.Api
}

func (e OrderItems) GetPage(c *gin.Context) {
	// 从请求中获取分页参数
	page := c.Query("page")
	pageSize := c.Query("pageSize")

	// 默认分页参数，你也可以根据需求自定义默认值
	defaultPage := 1
	defaultPageSize := 999

	// 将字符串参数转换为整数，处理可能的错误
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		pageNum = defaultPage
	}
	limit, err := strconv.Atoi(pageSize)
	if err != nil || limit < 1 {
		limit = defaultPageSize
	}

	s := service.OrderItems{}
	// 创建数据库连接和绑定请求
	err = e.MakeContext(c).
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

	err = s.GetOrderItemsPage(pageNum, limit, &objects)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(objects, "查询成功")
}

func (e OrderItems) Get(c *gin.Context) {
	s := service.OrderItems{}
	req := dto.OrderitemsGetReq{}
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

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

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
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建失败 err:%v", err))
		return
	}
	e.OK(req.GetId(), fmt.Sprintf("%v 创建成功", req.Title))
}

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
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

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
	fmt.Println(err)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除失败 err:%v", err))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
