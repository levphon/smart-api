// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/smart/models"
	"go-admin/app/smart/service/dto"
	"gorm.io/gorm"
)

type FlowTemplates struct {
	service.Service
}

// 分页获取FlowTemplates 所有的数据
func (e *FlowTemplates) GetFlowTemplatesPage(pageNum, limit int, objects *[]models.FlowTemplates, total *int64) error {
	// 计算偏移量
	offset := (pageNum - 1) * limit
	// 查询总数，不应用分页限制
	err := e.Orm.Model(&models.FlowTemplates{}).Count(total).Error
	if err != nil {
		e.Log.Errorf("查询总数失败: %s", err)
		return fmt.Errorf("查询总数失败: %s", err)
	}

	// 查询并分页获取订单项数据
	db := e.Orm.Limit(limit).Offset(offset).Find(objects)
	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询工单类别失败: %s", err)
		return fmt.Errorf("分页查询工单类别失败: %s", err)
	}
	return nil
}

// Get 获取FlowTemplates对象
func (e *FlowTemplates) Get(d *dto.FlowTemplatesGetReq, model *models.FlowTemplates) error {
	var err error
	var data models.FlowTemplates

	db := e.Orm.Model(&data).First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return fmt.Errorf("db error:%s", err)
	}
	return nil
}

// Insert 创建FlowTemplates对象
func (e *FlowTemplates) Insert(c *dto.FlowTemplatesInsertReq) error {
	var err error
	var data models.FlowTemplates
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var existingFlowTemplates models.FlowTemplates
	if err = tx.Where("name = ?", data.Name).First(&existingFlowTemplates).Error; err == nil {
		// 如果存在相同标题的订单项，返回相应的错误消息
		return fmt.Errorf(fmt.Sprintf("flow templates with name '%v' already exists", existingFlowTemplates.Name))
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果查询发生其他错误，返回错误
		return err
	}
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return fmt.Errorf("db error:%s", err)
	}
	return nil
}

// Update FlowTemplates
func (e *FlowTemplates) Update(c *dto.FlowTemplatesUpdateReq) error {
	var err error
	var model = models.FlowTemplates{}

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
			e.Log.Errorf("flow templates with ID '%v' not exists", c.GetId())
			return fmt.Errorf("flow templates with ID '%v' not exists", c.GetId())
		}
		e.Log.Errorf("Error querying flow templates with ID '%v': %s", c.GetId(), err)
		return fmt.Errorf("error querying flow templates with ID '%v': %s", c.GetId(), err)
	}

	// 检查 name 字段是否改变
	oldName := model.Name
	c.Generate(&model)

	// 执行更新操作
	db := tx.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateFlowTemplates error: %s", err)
		return fmt.Errorf("UpdateFlowTemplates error: %s", err)
	}

	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	// 检查模板名称是否有变化
	if oldName != model.Name {
		// 如果绑定模板的名称发生变化，更新 order_items 表中的 bindTempLate 和 Link 字段
		if err = updateOrderItemsBindTempLate(tx, oldName, model.Name); err != nil {
			e.Log.Errorf("Failed to update order_items after flow template update: %s", err)
			return fmt.Errorf("failed to update order_items after flow template update: %s", err)
		}
	}

	return nil
}

// Remove 删除FlowTemplates
func (e *FlowTemplates) Remove(d *dto.FlowTemplatesDeleteReq) error {
	var err error
	var data models.FlowTemplates
	// 查询要删除的模板是否存在
	if err = e.Orm.Model(&data).First(&data, d.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("templates with ID '%v' not found", d.GetId())
		}
		e.Log.Errorf("Error querying templates with ID '%v': %s", d.GetId(), err)
		return fmt.Errorf("error querying templates with ID '%v': %s", d.GetId(), err)
	}

	if data.BindCount != 0 {
		return fmt.Errorf("工单类别 %s 已经有工单在使用，请先删除相关工单", data.Name)
	}

	// 执行删除操作
	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return fmt.Errorf("Delete error: %s", err)
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// 更新 order_items 表中的 bindTempLate 和 Link 字段
func updateOrderItemsBindTempLate(tx *gorm.DB, oldName, newName string) error {
	// 更新 order_items 表中绑定了旧模板名称的记录
	err := tx.Model(&models.OrderItems{}).
		Where("bindTempLate = ?", oldName).
		Updates(map[string]interface{}{
			"bindTempLate": newName,
			"Link":         fmt.Sprintf("/order/classify/%s", newName),
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update order_items: %v", err)
	}
	return nil
}
