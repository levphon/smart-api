package dto

import (
	"go-admin/app/smart/models"
	common "go-admin/common/models"
)

type OrderItemsInsertReq struct {
	ID           uint   `uri:"id" comment:"编码"`              // 编码
	Title        string `json:"title" comment:"标题"`          // 标题
	BindTempLate string `json:"bindTempLate" comment:"绑定模板"` // 绑定模板
	Description  string `json:"description" comment:"描述信息"`  // 订单项描述
	Icon         string `json:"icon" comment:"图标"`           // 订单项图标
	CategoryID   uint   `json:"categoryId" comment:"类别ID"`   //类别ID
	Link         string `json:"link" comment:"链接"`           // 连接到网页
	Creator      string `json:"creator" comment:"创建人"`
	common.ControlBy
	common.ModelTime
}

func (s *OrderItemsInsertReq) Generate(model *models.OrderItems) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Title = s.Title
	model.BindTempLate = s.BindTempLate
	model.Description = s.Description
	model.Icon = s.Icon
	model.CategoryID = s.CategoryID
	model.Creator = s.Creator
	model.Link = s.Link
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *OrderItemsInsertReq) GetId() interface{} {
	return s.ID
}

type OrderItemsUpdateReq struct {
	ID           uint   `uri:"id" comment:"编码"`              // 编码
	Title        string `json:"title" comment:"标题"`          // 标题
	BindTempLate string `json:"bindTempLate" comment:"绑定模板"` // 绑定模板
	Description  string `json:"description" comment:"描述信息"`  // 订单项描述
	Icon         string `json:"icon" comment:"图标"`           // 订单项图标
	CategoryID   uint   `json:"categoryId" comment:"类别ID"`   //类别ID
	Link         string `json:"link" comment:"链接"`           // 连接到网页
	Regenerator  string `json:"regenerator" comment:"更新人"`   // 更新人
	common.ControlBy
	common.ModelTime
}

// 修改 Generate 方法以接受指针接收器
func (s *OrderItemsUpdateReq) Generate(model *models.OrderItems) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Title = s.Title
	model.BindTempLate = s.BindTempLate
	model.Description = s.Description
	model.Icon = s.Icon
	model.CategoryID = s.CategoryID
	model.Regenerator = s.Regenerator
	if s.Link != "" {
		model.Link = s.Link
	}
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *OrderItemsUpdateReq) GetId() interface{} {
	return s.ID
}

type OrderItemsGetReq struct {
	Id uint `uri:"id"`
}

func (s *OrderItemsGetReq) GetId() interface{} {
	return s.Id
}

type OrderItemsDeleteReq struct {
	Id uint `json:"id"`
}

func (s *OrderItemsDeleteReq) GetId() interface{} {
	return s.Id
}
