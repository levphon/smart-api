// @Author sunwenbo
// 2024/7/16 21:14
package models

import (
	models2 "go-admin/cmd/migrate/migration/models"
)

type FormData map[string]interface{}

type FlowTemplates struct {
	ID          uint     `gorm:"primaryKey;autoIncrement"`
	Name        string   `gorm:"column:name;type:varchar(100)" json:"name"`
	Description string   `gorm:"column:description;type:varchar(200)" json:"description"`
	BindCount   int      `gorm:"column:bindCount" json:"bindCount"`
	FormData    FormData `gorm:"type:json" json:"formData"`
	CategoryID  uint     `gorm:"column:categoryId" json:"categoryId"`
	models2.ControlBy
	models2.ModelTime
}

func (FlowTemplates) TableName() string {
	return "flow_templates"
}
