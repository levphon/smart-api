// @Author sunwenbo
// 2024/8/17 19:48
package models

import "smart-api/cmd/migrate/migration/models"

type OrderTask struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"column:name; type: varchar(256)" json:"name" form:"name"`                      // 任务名称
	TaskType    string `gorm:"column:taskType; type: varchar(45)" json:"taskType" form:"task_type"`          // 任务类型
	Interpreter string `gorm:"column:interpreter; type: varchar(100)" json:"interpreter" form:"interpreter"` // 解释器
	Content     string `gorm:"column:content; type: longtext" json:"content" form:"content"`                 // 任务内容
	Creator     string `gorm:"column:creator; type: varchar(45)" json:"creator" form:"creator"`              // 创建者
	Regenerator string `gorm:"regenerator:des;type:varchar(20)" json:"regenerator"`                          // 更新人
	Description string `gorm:"column:description; type: longtext" json:"description" form:"description"`     // 描述信息
	models.ControlBy
	models.ModelTime
}

func (*OrderTask) TableName() string {
	return "order_task"
}
