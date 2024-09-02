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

// GetPage 获取OrderWorks列表
// @Summary 获取OrderWorks列表
// @Description 获取OrderWorks列表
// @Tags 工单处理相关
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OrderWorks}} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-works [get]
// @Security Bearer
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

// Get 获取OrderWorks
// @Summary 获取OrderWorks
// @Description 获取OrderWorks
// @Tags 工单处理相关
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrderWorks} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-works/{id} [get]
// @Security Bearer
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

// MyBacklogGet 获取我的待办
// @Summary 获取我的待办
// @Description 获取我的待办
// @Tags 工单处理相关
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrderWorks} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-works/myBacklog [get]
// @Security Bearer
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

// CreatedByMe 获取我创建的
// @Summary 获取我创建的
// @Description 获取我创建的
// @Tags 工单处理相关
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrderWorks} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-works/createdByMe [get]
// @Security Bearer
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

// RelatedToMe 获取关于我的工单
// @Summary 获取关于我的工单
// @Description 获取关于我的工单
// @Tags 工单处理相关
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrderWorks} "{"code": 200, "data": [...]}"
// @Router /api/v1/order-works/relatedToMe [get]
// @Security Bearer
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

// Insert 创建OrderWorks
// @Summary 创建OrderWorks
// @Description 创建OrderWorks
// @Tags 工单处理相关
// @Accept application/json
// @Product application/json
// @Param data body dto.OrderWorksInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/order-works [post]
// @Security Bearer
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

// Update 修改OrderWorks
// @Summary 修改OrderWorks
// @Description 修改OrderWorks
// @Tags 工单处理相关
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrderWorksUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/order-works/{id} [put]
// @Security Bearer
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

// Handle 处理OrderWorks
// @Summary 处理OrderWorks
// @Description 处理OrderWorks
// @Tags 工单处理相关
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrderWorksUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/order-works/handle [put]
// @Security Bearer
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

// Delete 删除OrderWorks
// @Summary 删除OrderWorks
// @Description 删除OrderWorks
// @Tags 工单处理相关
// @Param data body dto.OrderWorksDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/order-works [delete]
// @Security Bearer
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
