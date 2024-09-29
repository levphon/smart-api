// @Author sunwenbo
// 2024/7/13 21:14
package models

import (
	"go-admin/common/models"
)

type UserFavorites struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	UserID      int  `gorm:"column:user_id" json:"userId"`
	OrderItemID uint `gorm:"column:order_item_id" json:"orderItemId"` // 确保类型一致
	models.ControlBy
	models.ModelTime
}

func (*UserFavorites) TableName() string {
	return "user_favorites"
}
func (e *UserFavorites) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *UserFavorites) GetId() interface{} {
	return e.ID
}
