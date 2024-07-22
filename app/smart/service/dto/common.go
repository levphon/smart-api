package dto

import (
	"go-admin/app/smart/models"
	common "go-admin/common/models"
	"time"
)

type WorksNotifyReq struct {
	ID             uint   `uri:"id" comment:"编码"`     // 编码
	Title          string `json:"title" comment:"标题"` // 工单标题
	Department     string `json:"department" comment:"部门"`
	Priority       string `json:"priority" comment:"优先级"`
	Status         string `json:"status" comment:"状态"`
	CurrentHandler string `json:"currentHandler" comment:"当前处理人"`
	Message        string `json:"message" comment:"消息内容"`    // 通知的详细消息内容。
	ReadStatus     int    `json:"readStatus" comment:"是否已读"` // 通知的阅读状态，例如是否已读。
	OrderID        uint   `json:"orderId"  comment:"关联工单ID"` // 关联的工单 ID，方便关联具体的工单。
	common.ModelTime
	common.ControlBy
}

func (n *WorksNotifyReq) Generate(model *models.WorksNotify) {
	if n.ID != 0 {
		model.ID = n.ID
	}
	model.Title = n.Title
	model.Department = n.Department
	model.Priority = n.Priority
	model.Status = n.Status
	model.CurrentHandler = n.CurrentHandler
	model.Message = n.Message
	model.ReadStatus = n.ReadStatus
	model.OrderID = n.OrderID
	model.ModelTime = n.ModelTime
	model.ControlBy = n.ControlBy
}

// GetId 获取数据对应的ID
func (n *WorksNotifyReq) GetId() interface{} {
	return n.ID
}

type WorksNotifyUpdateReq struct {
	ID         uint `uri:"id" comment:"编码"`            // 编码
	ReadStatus int  `json:"readStatus" comment:"是否已读"` // 通知的阅读状态，例如是否已读。
	common.ModelTime
	common.ControlBy
}

func (n *WorksNotifyUpdateReq) Generate(model *models.WorksNotify) {
	if n.ID != 0 {
		model.ID = n.ID
	}
	model.ReadStatus = n.ReadStatus
	model.UpdatedAt = common.JSONTime(time.Now())
	model.ControlBy = n.ControlBy
}

// GetId 获取数据对应的ID
func (n *WorksNotifyUpdateReq) GetId() interface{} {
	return n.ID
}

type WorksNotifyGetReq struct {
	Id uint `uri:"id"`
}

func (s *WorksNotifyGetReq) GetId() interface{} {
	return s.Id
}

//
//type WorksNotifyDeleteReq struct {
//	Id uint `json:"id"`
//}
//
//func (s *WorksNotifyDeleteReq) GetId() interface{} {
//	return s.Id
//}
