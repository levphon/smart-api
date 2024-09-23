// @Author sunwenbo
// 2024/8/17 21:14
package models

import (
	"go-admin/common/models"
)

type OrderComment struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID  int    `gorm:"column:order_id" json:"orderID"`                   // 与工单ID关联
	ParentID *int   `gorm:"column:parent_id" json:"parentID"`                 // 父留言ID，如果为空表示是顶级留言
	Comments string `gorm:"column:comment;type:varchar(255)" json:"comments"` // 留言内容
	models.ControlBy
	models.ModelTime
}

func (*OrderComment) TableName() string {
	return "order_comment"
}

func (e *OrderComment) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrderComment) GetId() interface{} {
	return e.ID
}
