// @Author sunwenbo
// 2024/7/13 21:14
package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"go-admin/common/models"
)

//type BindFlowData map[string]interface{}

type OrderWorks struct {
	ID               int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Title            string     `gorm:"column:title;type:varchar(100)" json:"title"`                  // 工单标题
	Creator          string     `gorm:"creator:des;type:varchar(20)" json:"creator"`                  // 创建人
	Regenerator      string     `gorm:"regenerator:des;type:varchar(20)" json:"regenerator"`          // 更新人
	CurrentNode      string     `gorm:"column:currentNode;type:varchar(50)" json:"currentNode"`       // 当前节点
	CurrentHandler   string     `gorm:"column:currentHandler;type:varchar(50)" json:"currentHandler"` // 当前处理人
	CurrentHandlerID int        `gorm:"column:currentHandlerId" json:"currentHandlerID"`              // 当前处理人ID
	Process          string     `gorm:"column:process;type:varchar(50)" json:"process"`               // 流程
	Template         string     `gorm:"column:template;type:varchar(50)" json:"template"`             // 模板
	Priority         string     `gorm:"column:priority;type:varchar(20)" json:"priority"`             // 优先级
	Status           string     `gorm:"column:status;type:varchar(20)" json:"status"`                 // 状态
	Department       string     `gorm:"column:department;type:varchar(50)" json:"department"`         // 部门
	Description      string     `gorm:"description:des;type:varchar(512)" json:"description"`
	FormData         FormData   `gorm:"type:json" json:"formData"`
	BindFlowData     FlowManage `gorm:"column:bindFlowData;type:json" json:"bindFlowData"`
	models.ControlBy
	models.ModelTime
}

func (*OrderWorks) TableName() string {
	return "order_works"
}

func (e *OrderWorks) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrderWorks) GetId() interface{} {
	return e.ID
}

// BindFlowData 的 Scan 和 Value 方法
func (e *FlowManage) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, &e)
}

func (e FlowManage) Value() (driver.Value, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}
