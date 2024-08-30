package dto

import (
	"go-admin/app/smart/models"
	common "go-admin/common/models"
)

type ExecMachineInsertReq struct {
	ID          int             `uri:"id" comment:"编码"`             // 编码
	Ip          string          `json:"ip" comment:"ip地址"`          // IP地址
	HostName    string          `json:"hostName" comment:"主机名"`     // 主机名
	UserName    string          `json:"userName" comment:"用户名"`     // 用户名
	PassWord    string          `json:"passWord" comment:"用户密码"`    // 密码
	Port        int             `json:"port" comment:"端口号"`         // 端口号
	Heartbeat   common.JSONTime `json:"heartbeat" comment:"心跳检查时间"` // 最近一次心跳时间
	Status      int             `json:"status" comment:"状态"`        // 状态
	AuthType    int             `json:"authType" comment:"认证类型"`    // 认证方式：1=用户名密码，2=公私钥
	KeyPath     string          `json:"keyPath" comment:"公私钥路径"`    // 公私钥路径
	Creator     string          `json:"creator" comment:"创建者"`      // 创建者
	Description string          `json:"description" comment:"描述"`   // 描述信息
	common.ControlBy
}

func (s *ExecMachineInsertReq) Generate(model *models.ExecMachine) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Ip = s.Ip
	model.HostName = s.HostName
	model.UserName = s.UserName
	model.PassWord = s.PassWord
	model.Port = s.Port
	model.Status = s.Status
	model.AuthType = s.AuthType
	model.KeyPath = s.KeyPath
	model.Creator = s.Creator
	model.Description = s.Description
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *ExecMachineInsertReq) GetId() interface{} {
	return s.ID
}

type ExecMachineUpdateReq struct {
	ID          int             `uri:"id" comment:"编码"`             // 编码
	Ip          string          `json:"ip" comment:"ip地址"`          // IP地址
	HostName    string          `json:"hostName" comment:"主机名"`     // 主机名
	UserName    string          `json:"userName" comment:"用户名"`     // 用户名
	PassWord    string          ` son:"passWord" comment:"用户密码"`    // 密码
	Port        int             `json:"port" comment:"端口号"`         // 端口号
	Heartbeat   common.JSONTime `json:"heartbeat" comment:"心跳检查时间"` // 最近一次心跳时间
	Status      int             `json:"status" comment:"状态"`        // 状态
	AuthType    int             `json:"authType" comment:"认证类型"`    // 认证方式：1=用户名密码，2=公私钥
	KeyPath     string          `json:"keyPath" comment:"公私钥路径"`    // 公私钥路径
	Regenerator string          `json:"regenerator" comment:"更新人"`  // 更新人
	Description string          `json:"description" comment:"描述"`   // 描述信息
	common.ControlBy
}

// Generate 结构体数据转化
func (s *ExecMachineUpdateReq) Generate(model *models.ExecMachine) {
	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Ip = s.Ip
	model.HostName = s.HostName
	model.UserName = s.UserName
	model.PassWord = s.PassWord
	model.Port = s.Port
	model.Heartbeat = s.Heartbeat
	model.Status = s.Status
	model.AuthType = s.AuthType
	model.KeyPath = s.KeyPath
	model.Regenerator = s.Regenerator // 确保映射
	model.Description = s.Description
	model.ControlBy = s.ControlBy
}

// GetId 获取数据对应的ID
func (s *ExecMachineUpdateReq) GetId() interface{} {
	return s.ID
}

type ExecMachineGetReq struct {
	Id int `uri:"id"`
}

func (s *ExecMachineGetReq) GetId() interface{} {
	return s.Id
}

type ExecMachineDeleteReq struct {
	Id int `json:"id"`
}

func (s *ExecMachineDeleteReq) GetId() interface{} {
	return s.Id
}
