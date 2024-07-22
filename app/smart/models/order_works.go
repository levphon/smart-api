// @Author sunwenbo
// 2024/7/13 21:14
package models

import (
	"go-admin/common/models"
)

type OrderWorks struct {
	ID             int      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title          string   `gorm:"column:title;type:varchar(100)" json:"title"`                  // 工单标题
	Creator        string   `gorm:"creator:des;type:varchar(20)" json:"creator"`                  // 创建人
	Regenerator    string   `gorm:"regenerator:des;type:varchar(20)" json:"regenerator"`          // 更新人
	CurrentNode    string   `gorm:"column:currentNode;type:varchar(50)" json:"currentNode"`       // 当前节点
	CurrentHandler string   `gorm:"column:currentHandler;type:varchar(20)" json:"currentHandler"` // 当前处理人
	Process        string   `gorm:"column:process;type:varchar(50)" json:"process"`               // 流程
	Priority       string   `gorm:"column:priority;type:varchar(20)" json:"priority"`             // 优先级
	Status         string   `gorm:"column:status;type:varchar(20)" json:"status"`                 // 状态
	Department     string   `gorm:"column:department;type:varchar(50)" json:"department"`         // 部门
	Description    string   `gorm:"description:des;type:varchar(512)" json:"description"`
	FormData       FormData `gorm:"type:json" json:"formData"`
	models.ControlBy
	models.ModelTime
}

func (*OrderWorks) TableName() string {
	return "order_works"
}

func (e *OrderWorks) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrderWorks) GetId() interface{} {
	return e.ID
}
