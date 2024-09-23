// @Author sunwenbo
// 2024/8/17 21:14
package models

import (
	"go-admin/common/models"
)

type OrderRating struct {
	ID      int `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID int `gorm:"column:order_id" json:"orderID"`        // 与工单ID关联
	Ratings int `gorm:"column:rating;type:int" json:"ratings"` // 评分
	models.ControlBy
	models.ModelTime
}

func (*OrderRating) TableName() string {
	return "order_rating"
}
