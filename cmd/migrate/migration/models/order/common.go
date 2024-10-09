// @Author sunwenbo
// 2024/7/18 16:45
package models

import (
	"smart-api/common/models"
	"time"
)

type OperationHistory struct {
	ID             uint      `gorm:"primaryKey" json:"id"`                                    // 主键ID
	Title          string    `gorm:"column:title;type:varchar(100)" json:"title"`             // 工单标题
	NodeName       string    `gorm:"column:node_name;type:varchar(255)" json:"nodeName"`      // 节点名称
	Transfer       string    `gorm:"column:transfer;type:varchar(255)" json:"transfer"`       // 流转
	Remark         string    `gorm:"column:remark;type:text" json:"remark"`                   // 备注
	Status         string    `gorm:"column:status;type:varchar(255)" json:"status"`           // 流转状态 1 同意， 0 拒绝， 2 其他
	HandlerId      int       `gorm:"column:handlerId;" json:"handlerId"`                      // 处理人ID
	HandlerName    string    `gorm:"column:handlerName;type:varchar(125)" json:"handlerName"` // 处理人姓名
	HandleTime     time.Time `gorm:"column:handleTime" json:"handleTime"`                     // 处理时间
	HandleDuration int64     `gorm:"column:handleDuration;type:bigint" json:"handleDuration"` // 处理时长 (以秒为单位)
	models.ModelTime
}

func (*OperationHistory) TableName() string {
	return "operation_history"
}

// 通知消息
type WorksNotify struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title          string `gorm:"column:title;type:varchar(100)" json:"title"`
	Department     string `gorm:"column:department;type:varchar(50)" json:"department"`
	Priority       string `gorm:"column:priority;type:varchar(20)" json:"priority"`
	Status         string `gorm:"column:status;type:varchar(20)" json:"status"`
	CurrentHandler int    `gorm:"column:currentHandler" json:"currentHandler"`
	Message        string `gorm:"column:message;type:text" json:"message"` // 通知的详细消息内容。
	ReadStatus     int    `gorm:"column:read_status" json:"readStatus"`    // 通知的阅读状态，例如是否已读。 0为未读 1为已读
	OrderID        uint   `gorm:"column:order_id" json:"orderId"`          // 关联的工单 ID，方便关联具体的工单。
	models.ModelTime
	models.ControlBy
}

func (*WorksNotify) TableName() string {
	return "works_notify"
}
