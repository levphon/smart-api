// @Author sunwenbo
// 2024/7/13 21:14
package models

import (
	"go-admin/common/models"
)

type OrderCategory struct {
	ID            int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string          `gorm:"column:name;type:varchar(50);unique" json:"name"`
	ChineseName   string          `gorm:"column:chineseName;type:varchar(50);unique" json:"chineseName"`
	FlowTemplates []FlowTemplates `gorm:"foreignKey:CategoryID" json:"flowTemplates"`
	models.ControlBy
	models.ModelTime
}

func (*OrderCategory) TableName() string {
	return "order_category"
}

func (e *OrderCategory) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrderCategory) GetId() interface{} {
	return e.ID
}
