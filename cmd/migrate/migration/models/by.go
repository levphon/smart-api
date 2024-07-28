package models

import (
	"go-admin/common/models"
	"gorm.io/gorm"
)

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt models.JSONTime `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt models.JSONTime `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt  `json:"-" gorm:"index;comment:删除时间"`
}
