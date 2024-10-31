// @Author sunwenbo
// 2024/7/13 21:14
package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"smart-api/common/models"
)

type StrucTure map[string]interface{}

type NoticeSlice []interface{}

type FlowManage struct {
	ID          int         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string      `gorm:"column:name;type:varchar(100)" json:"name"` // 流程名称
	CategoryID  uint        `gorm:"column:categoryId" json:"categoryId"`
	Notice      NoticeSlice `gorm:"column:notice;type:json" json:"notice"` // 变为json切片
	Comments    bool        `gorm:"column:comments;default:false" json:"comments"`
	Ratings     bool        `gorm:"column:ratings;default:false" json:"ratings"`
	Description string      `gorm:"description:des;type:varchar(512)" json:"description"`
	Creator     string      `gorm:"creator:des;type:varchar(20)" json:"creator"`         // 创建人
	Regenerator string      `gorm:"regenerator:des;type:varchar(20)" json:"regenerator"` // 更新人
	StrucTure   StrucTure   `gorm:"column:structure;type:json" json:"structure"`
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
func (e *NoticeSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("scan source is not []byte")
	}
	return json.Unmarshal(bytes, &e)
}

func (e NoticeSlice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}
