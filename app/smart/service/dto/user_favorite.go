package dto

import (
	"smart-api/app/smart/models"
	common "smart-api/common/models"
)

type UserFavoriteInsertReq struct {
	OrderItemID uint `json:"orderItemId" binding:"required"` // 工单项ID
	UserID      int  `json:"userID"`                         // 用户id
	common.ControlBy
}

// Generate 生成 OrderRating 模型实例
func (s *UserFavoriteInsertReq) Generate(userID int, model *models.UserFavorites) {
	model.OrderItemID = s.OrderItemID
	model.UserID = userID
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的 ID
func (s *UserFavoriteInsertReq) GetId() interface{} {
	return s.OrderItemID
}

type UserFavoriteDeleteReq struct {
	OrderItemID uint `json:"orderItemId" binding:"required"` // 工单项ID
}

// GetId 删除数据对应的 ID
func (s *UserFavoriteDeleteReq) GetId() interface{} {
	return s.OrderItemID
}
