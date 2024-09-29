// @Author sunwenbo
// 2024/7/13 21:14
package models

import (
	"go-admin/common/models"
)

type OrderItems struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Title        string `gorm:"column:title;type:varchar(50)" json:"title"`               // 订单项标题
	BindTempLate string `gorm:"column:bindTempLate;type:varchar(50)" json:"bindTempLate"` // 绑定模板
	Description  string `gorm:"column:description;type:varchar(256)" json:"description"`  // 订单项描述
	Icon         string `gorm:"column:icon;type:varchar(50)" json:"icon"`                 // 订单项图标
	CategoryID   uint   `gorm:"column:categoryId" json:"categoryId"`
	Link         string `gorm:"column:link;type:varchar(256)" json:"link"`           // 连接到网页
	Creator      string `gorm:"creator:des;type:varchar(20)" json:"creator"`         // 创建人
	Regenerator  string `gorm:"regenerator:des;type:varchar(20)" json:"regenerator"` // 更新人
	models.ControlBy
	models.ModelTime
}

func (*OrderItems) TableName() string {
	return "order_items"
}
