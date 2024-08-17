package dto

type SysNotify struct {
	OrderName   string `json:"orderName" binding:"required"`   // 工单名称
	ReceiveName string `json:"receiveName" binding:"required"` // 接收者名称
}
