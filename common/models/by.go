package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type JSONTime time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (jt JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", time.Time(jt).UTC().Format(timeFormat))
	return []byte(formatted), nil
}

// UnmarshalJSON parses JSON into JSONTime
func (jt *JSONTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(timeFormat, s)
	if err != nil {
		return err
	}
	*jt = JSONTime(t)
	return nil
}

// Scan implements the sql.Scanner interface
func (jt *JSONTime) Scan(value interface{}) error {
	if value == nil {
		*jt = JSONTime{}
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("cannot scan %T into JSONTime", value)
	}
	*jt = JSONTime(t)
	return nil
}

// Value implements the driver.Valuer interface
func (jt JSONTime) Value() (driver.Value, error) {
	return time.Time(jt).Format(timeFormat), nil
}

type ModelTime struct {
	CreatedAt JSONTime       `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt JSONTime       `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

// Convert JSONTime to time.Time
func (jt JSONTime) ToTime() time.Time {
	return time.Time(jt)
}
