package dto

import (
	"go-admin/app/smart/models"
	common "go-admin/common/models"
)

// OrderItemsGetPageReq 列表或者搜索使用结构体
type OrderItemsGetPageReq struct {
	ID           uint   `form:"id" search:"type:exact;column:id;table:order_items" comment:"id"`
	Title        string `form:"title" search:"type:exact;column:title;table:order_items" comment:"标题"`                 // 订单项标题
	BindTempLate string `form:"bindTempLate" search:"type:exact;column:bindTempLate;table:order_items" comment:"绑定模板"` // 绑定模板
	Description  string `form:"description" search:"type:exact;column:description;table:order_items" comment:"描述"`     // 订单项描述
	Favorite     bool   `form:"favorite" search:"type:exact;column:favorite;table:order_items" comment:"是否为收藏项"`       // 是否为收藏项
	Creator      string `form:"creator" search:"type:exact;column:creator;table:order_items" comment:"创建人"`            // 创建人
	Icon         string `form:"icon" search:"type:exact;column:icon;table:order_items" comment:"图标"`                   // 图标
	CategoryID   uint   `form:"categoryId" search:"type:exact;column:categoryId;table:order_items" comment:"类别"`
	Link         string `form:"link" search:"type:exact;column:link;table:order_items" comment:"连接到网页"` // 连接到网页
	common.ModelTime
}

func (m *OrderItemsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrderItemsInsertReq struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Title        string `gorm:"column:title;type:varchar(50)" json:"title"`               // 订单项标题
	BindTempLate string `gorm:"column:bindTempLate;type:varchar(50)" json:"bindTempLate"` // 绑定模板
	Description  string `gorm:"column:description;type:varchar(256)" json:"description"`  // 订单项描述
	Favorite     bool   `gorm:"default:false" json:"favorite"`                            // 是否为收藏项
	Creator      string `gorm:"column:creator;type:varchar(50)" json:"creator"`           // 创建人
	Icon         string `gorm:"column:icon;type:varchar(50)" json:"icon"`                 // 订单项图标
	CategoryID   uint   `gorm:"column:categoryId" json:"categoryId"`
	Link         string `gorm:"column:link;type:varchar(256)" json:"link"` // 连接到网页
	//CreatedAt    time.Time `gorm:"column:createAt;type:timestamp;default:CURRENT_TIMESTAMP" json:"createdAt"`
	//UpdatedAt    time.Time `gorm:"column:updateAt;type:timestamp;default:CURRENT_TIMESTAMP" json:"updatedAt"`
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
	model.Creator = s.Creator
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
	Creator      string `gorm:"column:creator;type:varchar(50)" json:"creator"`           // 创建人
	Icon         string `gorm:"column:icon;type:varchar(50)" json:"icon"`                 // 订单项图标
	CategoryID   uint   `gorm:"column:categoryId" json:"categoryId"`
	Link         string `gorm:"column:link;type:varchar(256)" json:"link"` // 连接到网页
	//CreatedAt    time.Time `gorm:"column:createAt;type:timestamp;default:CURRENT_TIMESTAMP" json:"createdAt"`
	//UpdatedAt    time.Time `gorm:"column:updateAt;type:timestamp;default:CURRENT_TIMESTAMP" json:"updatedAt"`
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
	model.Creator = s.Creator
	model.Icon = s.Icon
	model.CategoryID = s.CategoryID
	model.Link = s.Link
}

// GetId 获取数据对应的ID
func (s *OrderItemsUpdateReq) GetId() interface{} {
	return s.ID
}

type OrderitemsGetReq struct {
	Id int `uri:"id"`
}

func (s *OrderitemsGetReq) GetId() interface{} {
	return s.Id
}

type OrderItemsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrderItemsDeleteReq) GetId() interface{} {
	return s.Ids
}

type OrderitemsLabel struct {
	Id       int               `gorm:"-" json:"id"`
	Label    string            `gorm:"-" json:"label"`
	Children []OrderitemsLabel `gorm:"-" json:"children"`
}
