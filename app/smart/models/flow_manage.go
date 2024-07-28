// @Author sunwenbo
// 2024/7/13 21:14
package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"go-admin/common/models"
)

type StrucTure map[string]interface{}

type FlowSliceString []string
type FlowSliceInt []int

type FlowManage struct {
	ID          int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string          `gorm:"column:name;type:varchar(100)" json:"name"` // 流程名称
	Icon        string          `gorm:"column:icon;type:varchar(50)" json:"icon"`  // 订单项图标
	CategoryID  uint            `gorm:"column:categoryId" json:"categoryId"`
	Template    FlowSliceString `gorm:"column:template;type:json" json:"template"` // 变为json切片
	Task        FlowSliceString `gorm:"column:task;type:json" json:"task"`         // 变为json切片
	Notice      FlowSliceInt    `gorm:"column:notice;type:json" json:"notice"`     // 变为json切片
	Comments    bool            `gorm:"column:comments;default:false" json:"comments"`
	Ratings     bool            `gorm:"column:ratings;default:false" json:"ratings"`
	Description string          `gorm:"description:des;type:varchar(512)" json:"description"`
	Creator     string          `gorm:"creator:des;type:varchar(20)" json:"creator"`         // 创建人
	Regenerator string          `gorm:"regenerator:des;type:varchar(20)" json:"regenerator"` // 更新人
	//CurrentNode    string   `gorm:"column:currentNode;type:varchar(50)" json:"currentNode"`       // 当前节点
	//CurrentHandler string   `gorm:"column:currentHandler;type:varchar(20)" json:"currentHandler"` // 当前处理人
	StrucTure StrucTure `gorm:"column:structure;type:json" json:"structure"`
	models.ControlBy
	models.ModelTime
}

func (*FlowManage) TableName() string {
	return "flow_manage"
}

func (e *FlowManage) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FlowManage) GetId() interface{} {
	return e.ID
}

// FormData 的 Scan 和 Value 方法
func (e *StrucTure) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, &e)
}

func (e StrucTure) Value() (driver.Value, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

// FormData 的 Scan 和 Value 方法
func (e *FlowSliceString) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("scan source is not []byte")
	}
	return json.Unmarshal(bytes, &e)
}

func (e FlowSliceString) Value() (driver.Value, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

// FormData 的 Scan 和 Value 方法
func (e *FlowSliceInt) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("scan source is not []byte")
	}
	return json.Unmarshal(bytes, &e)
}

func (e FlowSliceInt) Value() (driver.Value, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}
