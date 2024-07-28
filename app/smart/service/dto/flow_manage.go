package dto

import (
	"go-admin/app/smart/models"
	common "go-admin/common/models"
)

type FlowManageInsertReq struct {
	ID          int                    `uri:"id" comment:"编码"`      // 编码
	Name        string                 `json:"name" comment:"流程名称"` // 流程名称
	Icon        string                 `json:"icon" comment:"图标"`   // 订单项图标
	CategoryID  uint                   `json:"categoryId" comment:"类别"`
	Template    models.FlowSliceString `json:"template" comment:"关联模板"`
	Task        models.FlowSliceString `json:"task" comment:"任务"`
	Notice      models.FlowSliceInt    `json:"notice" comment:"通知方式"`
	Comments    bool                   `json:"comments" comment:"是否开启留言"`
	Ratings     bool                   `json:"ratings" comment:"是否开启评分"`
	Description string                 `json:"description" comment:"流程描述"`
	Creator     string                 `json:"creator" comment:"创建人"`     // 创建人
	Regenerator string                 `json:"regenerator" comment:"更新人"` // 更新人
	StrucTure   models.StrucTure       `json:"structure" comment:"流程数据"`
	common.ControlBy
	common.ModelTime
}

func (s *FlowManageInsertReq) Generate(model *models.FlowManage) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Name = s.Name
	model.Icon = s.Icon
	model.CategoryID = s.CategoryID
	model.Template = s.Template
	model.Comments = s.Comments
	model.Ratings = s.Ratings
	model.Task = s.Task
	model.Notice = s.Notice
	model.Description = s.Description
	model.Creator = s.Creator
	model.Regenerator = s.Regenerator
	model.StrucTure = s.StrucTure
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *FlowManageInsertReq) GetId() interface{} {
	return s.ID
}

type FlowManageUpdateReq struct {
	ID          int                    `uri:"id" comment:"编码"`      // 编码
	Name        string                 `json:"name" comment:"流程名称"` // 流程名称
	Icon        string                 `json:"icon" comment:"图标"`   // 订单项图标
	CategoryID  uint                   `json:"categoryId" comment:"类别"`
	Template    models.FlowSliceString `json:"template" comment:"关联模板"`
	Task        models.FlowSliceString `json:"task" comment:"任务"`
	Notice      models.FlowSliceInt    `json:"notice" comment:"通知方式"`
	Comments    bool                   `json:"comments" comment:"是否开启留言"`
	Ratings     bool                   `json:"ratings" comment:"是否开启评分"`
	Description string                 `json:"description" comment:"流程描述"`
	Creator     string                 `json:"creator" comment:"创建人"`     // 创建人
	Regenerator string                 `json:"regenerator" comment:"更新人"` // 更新人
	StrucTure   models.StrucTure       `json:"structure" comment:"流程数据"`
	common.ControlBy
	common.ModelTime
}

// Generate 结构体数据转化 从 SysDeptControl 至 SysDept 对应的模型
func (s *FlowManageUpdateReq) Generate(model *models.FlowManage) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Name = s.Name
	model.Icon = s.Icon
	model.CategoryID = s.CategoryID
	model.Template = s.Template
	model.Comments = s.Comments
	model.Ratings = s.Ratings
	model.Task = s.Task
	model.Notice = s.Notice
	model.Description = s.Description
	model.Creator = s.Creator
	model.Regenerator = s.Regenerator
	model.StrucTure = s.StrucTure
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *FlowManageUpdateReq) GetId() interface{} {
	return s.ID
}

type FlowManageGetReq struct {
	Id int `uri:"id"`
}

func (s *FlowManageGetReq) GetId() interface{} {
	return s.Id
}

type FlowManageDeleteReq struct {
	Id int `json:"id"`
}

func (s *FlowManageDeleteReq) GetId() interface{} {
	return s.Id
}
