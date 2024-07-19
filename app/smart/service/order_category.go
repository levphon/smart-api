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

type OrderCategory struct {
	service.Service
}

// 分页获取OrderCategory 所有的数据
func (e *OrderCategory) GetOrderCategoryPage(pageNum, limit int, objects *[]models.OrderCategory) error {
	// 计算偏移量
	offset := (pageNum - 1) * limit

	// 查询并分页获取订单项数据
	db := e.Orm.Limit(limit).Offset(offset).Find(objects)
	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询工单类别失败: %s", err)
		return fmt.Errorf("分页查询工单类别失败: %s", err)
	}
	return nil
}

// Get 获取OrderCategory对象
func (e *OrderCategory) Get(d *dto.OrderCategoryGetReq, model *models.OrderCategory) error {
	var err error
	var data models.OrderCategory

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		e.Log.Errorf("db error:%s", err)
		return fmt.Errorf("查看对象不存在或无权查看")
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return fmt.Errorf("db error:%s", err)
	}
	return nil
}

// Insert 创建OrderCategory对象
func (e *OrderCategory) Insert(c *dto.OrderCategoryInsertReq) error {
	var err error
	var data models.OrderCategory
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
	var existingCategory models.OrderCategory
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

// Update OrderCategory
func (e *OrderCategory) Update(c *dto.OrderCategoryUpdateReq) error {
	var err error
	var model = models.OrderCategory{}

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
		return fmt.Errorf(fmt.Sprintf("Error querying order category with ID '%v': %s", c.GetId(), err))
	}
	// 检查 ChineseName 是否变化
	if model.ChineseName == c.ChineseName {
		e.Log.Infof("order category with ID '%v' has no changes in ChineseName", c.GetId())
		return errors.New(fmt.Sprintf("order category with ID '%v' has no changes in ChineseName", c.GetId()))
	}

	c.Generate(&model)

	// 执行更新操作
	db := tx.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateOrderCategory error: %s", err)
		return err
	}

	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysOrderCategory
func (e *OrderCategory) Remove(d *dto.OrderCategoryDeleteReq) error {
	var err error

	var data models.OrderCategory
	// 查询要删除的订单项
	if err = e.Orm.Model(&data).First(&data, d.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order with ID '%v' not found", d.GetId())
		}
		e.Log.Errorf("Error querying order with ID '%v': %s", d.GetId(), err)
		return fmt.Errorf("error querying order with ID '%v': %s", d.GetId(), err)
	}

	// 检查是否有工单在使用该类别
	var orderCount int64
	if err = e.Orm.Model(&models.OrderItems{}).Where("categoryId = ?", d.GetId()).Count(&orderCount).Error; err != nil {
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
		return fmt.Errorf("delete error: %s", err)
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}
