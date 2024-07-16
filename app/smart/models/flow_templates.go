// @Author sunwenbo
// 2024/7/13 21:14
package models

import (
	"go-admin/common/models"
)

type FormData map[string]interface{}

// vform模板管理表
type FlowTemplates struct {
	ID          uint     `gorm:"primaryKey;autoIncrement"`
	Name        string   `gorm:"column:name;type:varchar(100)" json:"name"`
	Description string   `gorm:"column:description;type:varchar(200)" json:"description"`
	BindCount   int      `gorm:"column:bindCount" json:"bindCount"`
	FormData    FormData `gorm:"type:json" json:"formData"`
	CategoryID  uint     `gorm:"column:categoryId" json:"categoryId"`
	models.ControlBy
	models.ModelTime
}

func (*FlowTemplates) TableName() string {
	return "flow_templates"
}

func (e *FlowTemplates) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FlowTemplates) GetId() interface{} {
	return e.ID
}
