package dto

import (
	"go-admin/app/smart/models"
	common "go-admin/common/models"
)

type OrderItemsInsertReq struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Title        string `gorm:"column:title;type:varchar(50)" json:"title"`               // 订单项标题
	BindTempLate string `gorm:"column:bindTempLate;type:varchar(50)" json:"bindTempLate"` // 绑定模板
	Description  string `gorm:"column:description;type:varchar(256)" json:"description"`  // 订单项描述
	Favorite     bool   `gorm:"default:false" json:"favorite"`                            // 是否为收藏项
	Icon         string `gorm:"column:icon;type:varchar(50)" json:"icon"`                 // 订单项图标
	CategoryID   uint   `gorm:"column:categoryId" json:"categoryId"`
	Link         string `gorm:"column:link;type:varchar(256)" json:"link"` // 连接到网页
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
	model.Favorite = s.Favorite
	model.Icon = s.Icon
	model.CategoryID = s.CategoryID
	model.Link = s.Link
}

// GetId 获取数据对应的ID
func (s *OrderItemsInsertReq) GetId() interface{} {
	return s.ID
}

type OrderItemsUpdateReq struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Title        string `gorm:"column:title;type:varchar(50)" json:"title"`               // 订单项标题
	BindTempLate string `gorm:"column:bindTempLate;type:varchar(50)" json:"bindTempLate"` // 绑定模板
	Description  string `gorm:"column:description;type:varchar(256)" json:"description"`  // 订单项描述
	Favorite     bool   `gorm:"default:false" json:"favorite"`                            // 是否为收藏项
	Icon         string `gorm:"column:icon;type:varchar(50)" json:"icon"`                 // 订单项图标
	CategoryID   uint   `gorm:"column:categoryId" json:"categoryId"`
	Link         string `gorm:"column:link;type:varchar(256)" json:"link"` // 连接到网页
	common.ControlBy
	common.ModelTime
}

// Generate 结构体数据转化 从 SysDeptControl 至 SysDept 对应的模型
func (s *OrderItemsUpdateReq) Generate(model *models.OrderItems) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Title = s.Title
	model.BindTempLate = s.BindTempLate
	model.Description = s.Description
	model.Favorite = s.Favorite
	model.Icon = s.Icon
	model.CategoryID = s.CategoryID
	model.Link = s.Link
}

// GetId 获取数据对应的ID
func (s *OrderItemsUpdateReq) GetId() interface{} {
	return s.ID
}

type OrderitemsGetReq struct {
	Id uint `uri:"id"`
}

func (s *OrderitemsGetReq) GetId() interface{} {
	return s.Id
}

type OrderItemsDeleteReq struct {
	Id uint `json:"id"`
}

func (s *OrderItemsDeleteReq) GetId() interface{} {
	return s.Id
}

type OrderitemsLabel struct {
	Id       uint              `gorm:"-" json:"id"`
	Label    string            `gorm:"-" json:"label"`
	Children []OrderitemsLabel `gorm:"-" json:"children"`
}
