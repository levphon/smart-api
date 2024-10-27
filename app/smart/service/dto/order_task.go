package dto

import (
	"smart-api/app/smart/models"
	common "smart-api/common/models"
)

type OrderTaskInsertReq struct {
	ID          int    `uri:"id" comment:"编码"`            // 编码
	Name        string `json:"name" comment:"标题"`         // 任务名称
	TaskType    string `json:"taskType" comment:"创建人"`    // 任务类型
	Interpreter string `json:"interpreter" comment:"解释器"` // 解释器
	Content     string `json:"content" comment:"当前节点"`    // 任务内容
	Creator     string `json:"creator" comment:"创建者"`     // 创建者
	Description string `json:"description" comment:"描述"`  // 描述信息
	common.ControlBy
}

func (s *OrderTaskInsertReq) Generate(model *models.OrderTask) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Name = s.Name
	model.TaskType = s.TaskType
	model.Interpreter = s.Interpreter
	model.Content = s.Content
	model.Creator = s.Creator
	model.Description = s.Description
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *OrderTaskInsertReq) GetId() interface{} {
	return s.ID
}

type OrderTaskUpdateReq struct {
	ID          int    `uri:"id" comment:"编码"`            // 编码
	Name        string `json:"name" comment:"标题"`         // 任务名称
	TaskType    string `json:"taskType" comment:"创建人"`    // 任务类型
	Interpreter string `json:"interpreter" comment:"解释器"` // 解释器
	Content     string `json:"content" comment:"当前节点"`    // 任务内容
	Regenerator string `json:"regenerator" comment:"更新者"` // 更新者
	Description string `json:"description" comment:"描述"`  // 描述信息
	common.ControlBy
}

// Generate 结构体数据转化
func (s *OrderTaskUpdateReq) Generate(model *models.OrderTask) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Name = s.Name
	model.TaskType = s.TaskType
	model.Interpreter = s.Interpreter
	model.Content = s.Content
	model.Regenerator = s.Regenerator
	model.Description = s.Description
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *OrderTaskUpdateReq) GetId() interface{} {
	return s.ID
}

type OrderTaskGetReq struct {
	Id int `uri:"id"`
}

func (s *OrderTaskGetReq) GetId() interface{} {
	return s.Id
}

type OrderTaskDeleteReq struct {
	Id int `json:"id"`
}

func (s *OrderTaskDeleteReq) GetId() interface{} {
	return s.Id
}

type OrderHistoryTaskDeleteReq struct {
	Id int `json:"id"`
}

func (s *OrderHistoryTaskDeleteReq) GetId() interface{} {
	return s.Id
}
