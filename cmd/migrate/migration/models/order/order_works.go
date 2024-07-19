// @Author sunwenbo
// 2024/7/17 21:14
package models

import (
	models2 "go-admin/cmd/migrate/migration/models"
)

type OrderWorks struct {
	ID             int      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title          string   `gorm:"column:title;type:varchar(100)" json:"title"`            // 工单标题
	CurrentNode    string   `gorm:"column:currentNode;type:varchar(50)" json:"currentNode"` // 当前节点
	CurrentHandler int      `gorm:"column:currentHandler" json:"currentHandler"`            // 当前处理人
	Process        string   `gorm:"column:process;type:varchar(50)" json:"process"`         // 流程
	Priority       string   `gorm:"column:priority;type:varchar(20)" json:"priority"`       // 优先级
	Status         string   `gorm:"column:status;type:varchar(20)" json:"status"`           // 状态
	Creator        string   `gorm:"column:creator;type:varchar(50)" json:"creator"`         // 创建人
	Department     string   `gorm:"column:department;type:varchar(50)" json:"department"`   // 部门
	Description    string   `gorm:"description:des;type:varchar(512)" json:"description"`
	FormData       FormData `gorm:"type:json" json:"formData"`
	models2.ControlBy
	models2.ModelTime
}

func (OrderWorks) TableName() string {
	return "order_works"
}
