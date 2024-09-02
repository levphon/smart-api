// @Author sunwenbo
// 2024/8/27 21:14
package models

import (
	"go-admin/common/models"
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
	KeyPath     string          `gorm:"column:key_path;type:varchar(255)" json:"keyPath" form:"keyPath"`                // 公私钥路径
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
