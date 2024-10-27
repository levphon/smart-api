// @Author sunwenbo
// 2024/8/17 19:48
package models

import (
	"smart-api/cmd/migrate/migration/models"
	models2 "smart-api/common/models"
)

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

type ExecutionHistoryTask struct {
	ID            int              `gorm:"primaryKey;autoIncrement" json:"id"`
	TaskID        int              `gorm:"column:task_id;" json:"taskID"`                        // 对应执行任务的ID
	TaskUUID      string           `gorm:"column:uuid; type:varchar(15)" json:"taskUUID"`        // 任务uuid
	TaskName      string           `gorm:"column:task_name; type:varchar(255)" json:"taskName"`  // 任务名称
	MachineID     int              `gorm:"column:machine_id;" json:"machineID"`                  // 执行机器的ID
	HostName      string           `gorm:"column:hostname; type:varchar(45)" json:"hostName"`    // 机器主机名
	Ip            string           `gorm:"column:ip; type:varchar(15)" json:"ip"`                // 机器IP地址
	ExecutionTime int64            `gorm:"column:execution_time;" json:"executionTime"`          // 执行时长
	Status        int              `gorm:"column:status;" json:"status" form:"status"`           // 执行状态 0 为成功 1 为失败 2为未知
	Stdout        string           `gorm:"column:stdout; type:longtext" json:"stdout"`           // 标准输出
	Stderr        string           `gorm:"column:stderr; type:longtext" json:"stderr"`           // 标准错误输出
	ExecutedAt    models2.JSONTime `gorm:"column:executed_at; type:timestamp" json:"executedAt"` // 执行时间
	Creator       string           `gorm:"column:creator; type:varchar(45)" json:"creator"`      // 执行发起人
	models.ControlBy
	models.ModelTime
}

func (*ExecutionHistoryTask) TableName() string {
	return "exec_machine_history"
}
