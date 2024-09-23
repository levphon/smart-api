package dto

import (
	"go-admin/app/smart/models"
	common "go-admin/common/models"
)

type OrderCommentInsertReq struct {
	OrderID  int    `json:"orderID" comment:"工单ID"`   // 工单ID
	ParentID *int   `json:"parentID" comment:"父留言ID"` // 父留言ID
	Comments string `json:"comments" comment:"留言"`    // 留言
	common.ControlBy
}

func (s *OrderCommentInsertReq) Generate(model *models.OrderComment) {
	model.OrderID = s.OrderID
	model.ParentID = s.ParentID
	model.Comments = s.Comments
	model.ControlBy = s.ControlBy
}

func (s *OrderCommentInsertReq) GetId() interface{} {
	return s.OrderID
}

type OrderCommentUpdateReq struct {
	ID       int    `uri:"id" comment:"ID"`         // 留言ID
	OrderID  int    `json:"orderID" comment:"工单ID"` // 工单ID
	ParentID *int   `json:"parentID" comment:"父留言ID"`
	Comments string `json:"comments" comment:"留言"` // 留言内容
	common.ControlBy
	common.ModelTime
}

func (s *OrderCommentUpdateReq) Generate(model *models.OrderComment) {
	model.ID = s.ID
	model.OrderID = s.OrderID
	model.ParentID = s.ParentID
	model.Comments = s.Comments
	model.ControlBy = s.ControlBy
}

func (s *OrderCommentUpdateReq) GetId() interface{} {
	return s.ID
}

type OrderCommentGetReq struct {
	OrderID int `uri:"orderID" comment:"工单ID"` // 工单ID
}

// GetId 获取留言的 ID
func (s *OrderCommentGetReq) GetId() interface{} {
	return s.OrderID
}

type OrderCommentDeleteReq struct {
	ID int `json:"id" comment:"留言ID"` // 留言ID
}

// GetId 获取留言的 ID
func (s *OrderCommentDeleteReq) GetId() interface{} {
	return s.ID
}
