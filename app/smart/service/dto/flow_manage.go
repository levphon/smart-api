package dto

import (
	"smart-api/app/smart/models"
	common "smart-api/common/models"
)

type FlowManageInsertReq struct {
	ID          int                `uri:"id" comment:"编码"`      // 编码
	Name        string             `json:"name" comment:"流程名称"` // 流程名称
	CategoryID  uint               `json:"categoryId" comment:"类别"`
	Notice      models.NoticeSlice `json:"notice" comment:"通知方式"`
	Comments    bool               `json:"comments" comment:"是否开启留言"`
	Ratings     bool               `json:"ratings" comment:"是否开启评分"`
	Description string             `json:"description" comment:"流程描述"`
	Creator     string             `json:"creator" comment:"创建人"` // 创建人
	StrucTure   models.StrucTure   `json:"structure" comment:"流程数据"`
	common.ControlBy
	common.ModelTime
}

func (s *FlowManageInsertReq) Generate(model *models.FlowManage) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Name = s.Name
	model.CategoryID = s.CategoryID
	model.Comments = s.Comments
	model.Ratings = s.Ratings
	model.Notice = s.Notice
	model.Creator = s.Creator
	model.Description = s.Description
	model.StrucTure = s.StrucTure
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *FlowManageInsertReq) GetId() interface{} {
	return s.ID
}

type FlowManageUpdateReq struct {
	ID          int                `uri:"id" comment:"编码"`      // 编码
	Name        string             `json:"name" comment:"流程名称"` // 流程名称
	CategoryID  uint               `json:"categoryId" comment:"类别"`
	Notice      models.NoticeSlice `json:"notice" comment:"通知方式"`
	Comments    bool               `json:"comments" comment:"是否开启留言"`
	Ratings     bool               `json:"ratings" comment:"是否开启评分"`
	Description string             `json:"description" comment:"流程描述"`
	Regenerator string             `json:"regenerator" comment:"更新人"` // 更新人
	StrucTure   models.StrucTure   `json:"structure" comment:"流程数据"`
	common.ControlBy
	common.ModelTime
}

// Generate 结构体数据转化 从 SysDeptControl 至 SysDept 对应的模型
func (s *FlowManageUpdateReq) Generate(model *models.FlowManage) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Name = s.Name
	model.CategoryID = s.CategoryID
	model.Comments = s.Comments
	model.Ratings = s.Ratings
	model.Notice = s.Notice
	model.Description = s.Description
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
