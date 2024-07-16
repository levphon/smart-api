package dto

import (
	"go-admin/app/smart/models"
	common "go-admin/common/models"
)

type FormData map[string]interface{}

type FlowTemplatesInsertReq struct {
	ID          uint     `uri:"id" comment:"编码"`    // 编码
	Name        string   `json:"name" comment:"标题"` // 标题
	Description string   `json:"description" comment:"描述信息"`
	BindCount   int      `json:"bindCount" comment:"绑定次数"`
	FormData    FormData `json:"formData" comment:"表单数据"`
	CategoryID  uint     `json:"categoryId" comment:"类别ID"`
	common.ControlBy
}

func (s *FlowTemplatesInsertReq) Generate(model *models.FlowTemplates) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Name = s.Name
	model.Description = s.Description
	model.BindCount = s.BindCount
	model.FormData = models.FormData(s.FormData)
	model.CategoryID = s.CategoryID
}

// GetId 获取数据对应的ID
func (s *FlowTemplatesInsertReq) GetId() interface{} {
	return s.ID
}

type FlowTemplatesUpdateReq struct {
	ID          uint     `uri:"id" comment:"编码"`    // 编码
	Name        string   `json:"name" comment:"标题"` // 标题
	Description string   `json:"description" comment:"描述信息"`
	BindCount   int      `json:"bindCount" comment:"绑定次数"`
	FormData    FormData `json:"formData" comment:"表单数据"`
	CategoryID  uint     `json:"categoryId" comment:"类别ID"`
	common.ControlBy
}

// Generate 结构体数据转化 从 SysDeptControl 至 SysDept 对应的模型
func (s *FlowTemplatesUpdateReq) Generate(model *models.FlowTemplates) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Name = s.Name
	model.Description = s.Description
	model.BindCount = s.BindCount
	model.FormData = models.FormData(s.FormData)
	model.CategoryID = s.CategoryID
}

// GetId 获取数据对应的ID
func (s *FlowTemplatesUpdateReq) GetId() interface{} {
	return s.ID
}

type FlowTemplatesGetReq struct {
	Id int `uri:"id"`
}

func (s *FlowTemplatesGetReq) GetId() interface{} {
	return s.Id
}

type FlowTemplatesDeleteReq struct {
	Id int `json:"id"`
}

func (s *FlowTemplatesDeleteReq) GetId() interface{} {
	return s.Id
}
