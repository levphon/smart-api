// @Author sunwenbo
// 2024/7/16 21:14
package models

import (
	models2 "go-admin/cmd/migrate/migration/models"
)

type OrderCategory struct {
	ID            int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string          `gorm:"column:name;type:varchar(50)" json:"name"`
	ChineseName   string          `gorm:"column:chineseName;type:varchar(50)" json:"chineseName"`
	FlowTemplates []FlowTemplates `gorm:"foreignKey:CategoryID" json:"flowTemplates"`
	models2.ControlBy
	models2.ModelTime
}

func (OrderCategory) TableName() string {
	return "order_category"
}
