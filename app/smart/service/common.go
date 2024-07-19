// @Author sunwenbo
// 2024/7/19 17:56
package service

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/smart/models"
	"go-admin/app/smart/service/dto"
	"gorm.io/gorm"
)

type WorksNotify struct {
	service.Service
}

func (e *WorksNotify) Insert(c *dto.WorksNotifyReq) error {
	var err error
	var data models.WorksNotify
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	fmt.Println(&data)
	var count int64
	if err := tx.Model(&models.WorksNotify{}).Where("order_id = ?", data.OrderID).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count notifications")
	}

	// 如果同一个工单的通知消息数量超过五条，返回提示消息
	if count >= 3 {
		return fmt.Errorf(fmt.Sprintf("%v 催办上限 大于3 次", data.Title))
	}
	message := fmt.Sprintf("工单标题: %s 需要你的处理", data.Title)
	data.Message = message
	data.ReadStatus = 0

	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf(fmt.Sprintf(" 创建%v通知消息失败 ", data.Title))
		return fmt.Errorf(" 创建%v通知消息失败 ", data.Title)
	}
	return nil
}

func (e *WorksNotify) GetNotify(pageNum, limit int, objects *[]models.WorksNotify, userid int) error {
	// 计算偏移量
	offset := (pageNum - 1) * limit
	// 查询并分页获取订单项数据
	db := e.Orm.Where("currentHandler = ?", userid).Limit(limit).Offset(offset).Find(objects)

	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询通知信息失败: %s", err)
		return fmt.Errorf("分页查询通知信息失败: %s", err)
	}
	// 判断结果数量是否为 0
	if db.RowsAffected == 0 {
		e.Log.Infof("没有通知消息")
		return nil
	}
	return nil
}

func (e *WorksNotify) UpdateNotify(c *dto.WorksNotifyUpdateReq) error {
	var err error
	var model = models.WorksNotify{}

	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	// 根据 ID 查找要更新的记录
	if err = tx.First(&model, c.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("notify with ID '%v' not exists", c.GetId())
			return fmt.Errorf("notify with ID '%v' not exists", c.GetId())
		}
		e.Log.Errorf("Error querying notify with ID '%v': %s", c.GetId(), err)
		return fmt.Errorf("error querying notify with ID '%v': %s", c.GetId(), err)
	}

	fmt.Println("model.ReadStatus == c.ReadStatus", model.ReadStatus, c.ReadStatus)

	// 判断是否需要执行更新操作
	if model.ReadStatus == c.ReadStatus {
		return fmt.Errorf(fmt.Sprintf("No changes for notify with '%v'", model.Title))
	}

	c.Generate(&model)

	// 显式地更新字段
	updates := map[string]interface{}{
		"read_status": c.ReadStatus,
		"updated_at":  model.UpdatedAt, // 更新更新时间
	}

	// 更新通知消息字段
	if err = tx.Model(&model).Updates(updates).Error; err != nil {
		// 处理更新错误
		e.Log.Errorf("Failed to update notify with title '%v' notify ", model.Title)
		return fmt.Errorf("Failed to update notify with title '%v' notify ", model.Title)
	}

	return nil
}
