package dto

import (
	"smart-api/app/smart/models"
	common "smart-api/common/models"
)

type OrderRatingInsertReq struct {
	OrderID int `json:"orderID" comment:"工单ID"` // 工单ID
	Ratings int `json:"ratings" comment:"评分"`   // 评分
	common.ControlBy
}

// Generate 生成 OrderRating 模型实例
func (s *OrderRatingInsertReq) Generate(model *models.OrderRating) {
	model.OrderID = s.OrderID
	model.Ratings = s.Ratings
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的 ID
func (s *OrderRatingInsertReq) GetId() interface{} {
	return s.OrderID
}

type OrderRatingUpdateReq struct {
	ID      int `uri:"id" comment:"ID"`         // 评分ID
	OrderID int `json:"orderID" comment:"工单ID"` // 工单ID
	Ratings int `json:"ratings" comment:"评分"`   // 评分
	common.ControlBy
	common.ModelTime
}

// Generate 生成 OrderRating 模型实例
func (s *OrderRatingUpdateReq) Generate(model *models.OrderRating) {
	model.ID = s.ID
	model.OrderID = s.OrderID
	model.Ratings = s.Ratings
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的 ID
func (s *OrderRatingUpdateReq) GetId() interface{} {
	return s.ID
}

type OrderRatingGetReq struct {
	OrderId int `uri:"orderID" comment:"工单ID"` // 评分ID
}

// GetId 获取数据对应的 ID
func (s *OrderRatingGetReq) GetId() interface{} {
	return s.OrderId
}

type OrderRatingDeleteReq struct {
	ID int `json:"id" comment:"ID"` // 评分ID
}

// GetId 获取数据对应的 ID
func (s *OrderRatingDeleteReq) GetId() interface{} {
	return s.ID
}
