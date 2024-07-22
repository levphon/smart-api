// @Author sunwenbo
// 2024/7/13 21:14
package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"go-admin/common/models"
)

type FormData map[string]interface{}

// vform模板管理表
type FlowTemplates struct {
	ID          uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string   `gorm:"column:name;type:varchar(100)" json:"name"`
	Description string   `gorm:"column:description;type:varchar(200)" json:"description"`
	BindCount   int      `gorm:"column:bindCount" json:"bindCount"`
	FormData    FormData `gorm:"type:json" json:"formData"`
	CategoryID  uint     `gorm:"column:categoryId" json:"categoryId"`
	Creator     string   `gorm:"creator:des;type:varchar(20)" json:"creator"`         // 创建人
	Regenerator string   `gorm:"regenerator:des;type:varchar(20)" json:"regenerator"` // 更新人
	models.ControlBy
	models.ModelTime
}

func (*FlowTemplates) TableName() string {
	return "flow_templates"
}

func (e *FlowTemplates) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FlowTemplates) GetId() interface{} {
	return e.ID
}

// FormData 的 Scan 和 Value 方法
func (e *FormData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, &e)
}

func (e FormData) Value() (driver.Value, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}
