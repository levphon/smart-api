// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"fmt"
	"go-admin/app/smart/models"
	"time"

	"gorm.io/gorm"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/smart/service/dto"
)

type FlowTemplates struct {
	service.Service
}

// 分页获取FlowTemplates 所有的数据
func (e *FlowTemplates) GetFlowTemplatesPage(pageNum, limit int, objects *[]models.FlowTemplates) error {
	// 计算偏移量
	offset := (pageNum - 1) * limit

	// 查询并分页获取订单项数据
	db := e.Orm.Limit(limit).Offset(offset).Find(objects)
	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询工单类别失败: %s", err)
		return err
	}
	return nil
}

// Get 获取FlowTemplates对象
func (e *FlowTemplates) Get(d *dto.FlowTemplatesGetReq, model *models.FlowTemplates) error {
	var err error
	var data models.FlowTemplates

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
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
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	var existingCategory models.FlowTemplates
	if err = tx.Where("name = ?", data.Name).First(&existingCategory).Error; err == nil {
		// 如果存在相同标题的订单项，返回相应的错误消息
		return fmt.Errorf(fmt.Sprintf("order category with name '%v' already exists", existingCategory.Name))
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果查询发生其他错误，返回错误
		return err
	}

	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
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
			e.Log.Errorf("order category with ID '%v' not exists", c.GetId())
			return fmt.Errorf("order category with ID '%v' not exists", c.GetId())
		}
		e.Log.Errorf("Error querying order category with ID '%v': %s", c.GetId(), err)
		return err
	}

	c.Generate(&model)

	// 执行更新操作
	db := tx.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateFlowTemplates error: %s", err)
		return err
	}

	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysFlowTemplates
func (e *FlowTemplates) Remove(d *dto.FlowTemplatesDeleteReq) error {
	var err error
	var data models.FlowTemplates
	// 查询要删除的订单项
	if err = e.Orm.Model(&data).First(&data, d.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order with ID '%v' not found", d.GetId())
		}
		e.Log.Errorf("Error querying order with ID '%v': %s", d.GetId(), err)
		return err
	}

	// 检查是否有工单在使用该类别
	var orderCount int64
	if err = e.Orm.Model(&models.FlowTemplates{}).Where("categoryId = ?", d.GetId()).Count(&orderCount).Error; err != nil {
		e.Log.Errorf("Error checking order items: %v", err)
		return fmt.Errorf("error checking order items: %v", err)
	}

	if orderCount > 0 {
		return fmt.Errorf("工单类别 %s 已经有工单在使用，请先删除相关工单", data.Name)
	}

	// 执行删除操作
	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}
