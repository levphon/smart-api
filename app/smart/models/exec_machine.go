// @Author sunwenbo
// 2024/8/27 21:14
package models

import (
	"smart-api/common/models"
)

type ExecMachine struct {
	ID          int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Ip          string          `gorm:"column:ip; type: varchar(15)" json:"ip" form:"ip"`                               // IP地址
	HostName    string          `gorm:"column:hostname; type: varchar(45)" json:"hostName" form:"hostName"`             // 主机名
	UserName    string          `gorm:"column:username;type: varchar(45)" json:"userName" form:"username"`              // 用户名
	PassWord    string          `gorm:"column:password;type: varchar(100)" json:"passWord" form:"password"`             // 密码
	Port        int             `gorm:"column:port;" json:"port" form:"port"`                                           // 端口
	Status      int             `gorm:"column:status;" json:"status" form:"status"`                                     // 状态  1 为在线 2为离线
	AuthType    string          `gorm:"column:auth_type;type:varchar(10)" json:"authType" form:"authType"`              // 认证方式：1=用户名密码，2=公私钥
	PrivateKey  string          `gorm:"column:private_key;type:varchar(4096)" json:"privateKey" form:"privateKey"`      // 私钥内容
	Creator     string          `gorm:"column:creator; type: varchar(45)" json:"creator" form:"creator"`                // 创建者
	Heartbeat   models.JSONTime `gorm:"column:heartbeat;type:timestamp;default:NULL" json:"heartbeat" form:"heartbeat"` // 最近一次心跳时间
	Regenerator string          `gorm:"regenerator:des;type:varchar(20)" json:"regenerator"`                            // 更新人
	Description string          `gorm:"column:description; type: longtext" json:"description" form:"description"`       // 描述信息
	models.ControlBy
	models.ModelTime
}

func (*ExecMachine) TableName() string {
	return "exec_machine"
}

func (e *ExecMachine) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ExecMachine) GetId() interface{} {
	return e.ID
}

type ExecutionHistory struct {
	ID            int             `gorm:"primaryKey;autoIncrement" json:"id"`
	TaskID        int             `gorm:"column:task_id;" json:"taskID"`                        // 对应执行任务的ID
	TaskName      string          `gorm:"column:task_name; type:varchar(255)" json:"taskName"`  // 任务名称
	MachineID     int             `gorm:"column:machine_id;" json:"machineID"`                  // 执行机器的ID
	HostName      string          `gorm:"column:hostname; type:varchar(45)" json:"hostName"`    // 机器主机名
	ExecutionTime int64           `gorm:"column:execution_time;" json:"executionTime"`          // 执行时长
	Ip            string          `gorm:"column:ip; type:varchar(15)" json:"ip"`                // 机器IP地址
	Status        int             `gorm:"column:status;" json:"status" form:"status"`           // 执行状态 0 为成功 1 为失败 2为未知
	Stdout        string          `gorm:"column:stdout; type:longtext" json:"stdout"`           // 标准输出
	Stderr        string          `gorm:"column:stderr; type:longtext" json:"stderr"`           // 标准错误输出
	ExecutedAt    models.JSONTime `gorm:"column:executed_at; type:timestamp" json:"executedAt"` // 执行时间
	Creator       string          `gorm:"column:creator; type:varchar(45)" json:"creator"`      // 执行发起人
	models.ControlBy
	models.ModelTime
}

func (*ExecutionHistory) TableName() string {
	return "exec_machine_history"
}

func (e *ExecutionHistory) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ExecutionHistory) GetId() interface{} {
	return e.ID
}
