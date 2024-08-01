package dto

import (
	"go-admin/app/smart/models"
	common "go-admin/common/models"
)

type OrderWorksInsertReq struct {
	ID               int             `uri:"id" comment:"编码"`     // 编码
	Title            string          `json:"title" comment:"标题"` // 工单标题
	Creator          string          `json:"creator" comment:"创建人"`
	CurrentNode      string          `json:"currentNode" comment:"当前节点"`         // 当前节点
	CurrentHandler   string          `json:"currentHandler" comment:"当前处理人"`     // 当前处理人
	CurrentHandlerId int             `json:"currentHandlerId" comment:"当前处理人ID"` // 当前处理人
	Process          string          `json:"process" comment:"流程"`               // 流程
	Priority         string          `json:"priority" comment:"优先级"`             // 优先级
	Template         string          `json:"template" comment:"模板"`              // 模板
	Status           string          `json:"status" comment:"状态"`                // 状态
	Department       string          `json:"department" comment:"部门"`            // 部门
	Description      string          `json:"description" comment:"描述"`           // 描述
	FormData         models.FormData `json:"formData" comment:"工单表单数据"`          // 工单表单数据
	common.ControlBy
}

func (s *OrderWorksInsertReq) Generate(model *models.OrderWorks) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Title = s.Title
	model.Creator = s.Creator
	model.CurrentNode = s.CurrentNode
	model.CurrentHandler = s.CurrentHandler
	model.CurrentHandlerID = s.CurrentHandlerId
	model.Process = s.Process
	model.Priority = s.Priority
	model.Status = s.Status
	model.Template = s.Template
	model.Department = s.Department
	model.Description = s.Description
	model.FormData = s.FormData
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *OrderWorksInsertReq) GetId() interface{} {
	return s.ID
}

type OrderWorksUpdateReq struct {
	ID               int             `uri:"id" comment:"编码"`                     // 编码
	Title            string          `json:"title" comment:"标题"`                 // 工单标题
	Regenerator      string          `json:"regenerator" comment:"更新人"`          // 更新人
	CurrentNode      string          `json:"currentNode" comment:"当前节点"`         // 当前节点
	CurrentHandler   string          `json:"currentHandler" comment:"当前处理人"`     // 当前处理人
	CurrentHandlerId int             `json:"currentHandlerId" comment:"当前处理人ID"` // 当前处理人
	Process          string          `json:"process" comment:"流程"`               // 流程
	Priority         string          `json:"priority" comment:"优先级"`             // 优先级
	Template         string          `json:"template" comment:"模板"`              // 模板
	Status           string          `json:"status" comment:"状态"`                // 状态
	Department       string          `json:"department" comment:"部门"`            // 部门
	Description      string          `json:"description" comment:"描述"`           // 描述
	FormData         models.FormData `json:"formData" comment:"工单表单数据"`          // 工单表单数据
	common.ControlBy                 // 包含创建人和更新人
	common.ModelTime
}

// Generate 结构体数据转化
func (s *OrderWorksUpdateReq) Generate(model *models.OrderWorks) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Title = s.Title
	model.Regenerator = s.Regenerator
	model.CurrentNode = s.CurrentNode
	model.CurrentHandler = s.CurrentHandler
	model.CurrentHandlerID = s.CurrentHandlerId
	model.Process = s.Process
	model.Template = s.Template
	model.Priority = s.Priority
	model.Status = s.Status
	model.Department = s.Department
	model.Description = s.Description
	model.FormData = s.FormData
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *OrderWorksUpdateReq) GetId() interface{} {
	return s.ID
}

type OrderWorksGetReq struct {
	Id int `uri:"id"`
}

func (s *OrderWorksGetReq) GetId() interface{} {
	return s.Id
}

type OrderWorksDeleteReq struct {
	Id int `json:"id"`
}

func (s *OrderWorksDeleteReq) GetId() interface{} {
	return s.Id
}

type OperationHistoryGetReq struct {
	Id int `uri:"id"`
}

func (h *OperationHistoryGetReq) GetId() interface{} {
	return h.Id
}
