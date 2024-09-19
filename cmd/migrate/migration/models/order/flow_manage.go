// @Author sunwenbo
// 2024/7/16 21:14
package models

import (
	"go-admin/common/models"
)

type StrucTure map[string]interface{}
type FlowSliceInt []int

type FlowManage struct {
	ID          int          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string       `gorm:"column:name;type:varchar(100)" json:"name"` // 流程名称
	CategoryID  uint         `gorm:"column:categoryId" json:"categoryId"`
	Notice      FlowSliceInt `gorm:"column:notice;type:json" json:"notice"` // 变为json切片
	Comments    bool         `gorm:"column:comments;default:false" json:"comments"`
	Ratings     bool         `gorm:"column:ratings;default:false" json:"ratings"`
	Description string       `gorm:"description:des;type:varchar(512)" json:"description"`
	Creator     string       `gorm:"creator:des;type:varchar(20)" json:"creator"`         // 创建人
	Regenerator string       `gorm:"regenerator:des;type:varchar(20)" json:"regenerator"` // 更新人
	StrucTure   StrucTure    `gorm:"column:structure;type:json" json:"structure"`
	models.ControlBy
	models.ModelTime
}

func (*FlowManage) TableName() string {
	return "flow_manage"
}
