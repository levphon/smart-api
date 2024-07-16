package dto

import (
	"go-admin/app/smart/models"
	common "go-admin/common/models"
)

type OrderCategoryInsertReq struct {
	ID          int    `uri:"id" comment:"编码"`             // 编码
	Name        string `json:"name" comment:"标题"`          // 标题
	ChineseName string `json:"chineseName" comment:"中文名称"` //中文名称
	common.ControlBy
}

func (s *OrderCategoryInsertReq) Generate(model *models.OrderCategory) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Name = s.Name
	model.ChineseName = s.ChineseName
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *OrderCategoryInsertReq) GetId() interface{} {
	return s.ID
}

type OrderCategoryUpdateReq struct {
	ID          int    `uri:"id" comment:"编码"`             // 编码
	Name        string `json:"name" comment:"标题"`          // 标题
	ChineseName string `json:"chineseName" comment:"中文名称"` //中文名称
	common.ControlBy
}

// Generate 结构体数据转化 从 SysDeptControl 至 SysDept 对应的模型
func (s *OrderCategoryUpdateReq) Generate(model *models.OrderCategory) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Name = s.Name
	model.ChineseName = s.ChineseName
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *OrderCategoryUpdateReq) GetId() interface{} {
	return s.ID
}

type OrderCategoryGetReq struct {
	Id int `uri:"id"`
}

func (s *OrderCategoryGetReq) GetId() interface{} {
	return s.Id
}

type OrderCategoryDeleteReq struct {
	Id int `json:"id"`
}

func (s *OrderCategoryDeleteReq) GetId() interface{} {
	return s.Id
}
