// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"fmt"
	"go-admin/app/smart/models"
	"gorm.io/gorm"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/smart/service/dto"
)

type OrderTask struct {
	service.Service
}

// 分页获取OrderTask 所有的数据
func (e *OrderTask) GetOrderTaskPage(pageNum, limit int, objects *[]models.OrderTask, total *int64) error {

	// 计算偏移量
	offset := (pageNum - 1) * limit

	// 查询总数，不应用分页限制
	err := e.Orm.Model(&models.OrderTask{}).Count(total).Error
	if err != nil {
		e.Log.Errorf("查询总数失败: %s", err)
		return fmt.Errorf("查询总数失败: %s", err)
	}

	// 查询并分页获取订单项数据
	db := e.Orm.Limit(limit).Offset(offset).Find(objects)

	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询工单失败: %s", err)
		return fmt.Errorf("分页查询工单失败: %s", err)
	}
	return nil
}

// 根据ID获取任务
func (e *OrderTask) Get(d *dto.OrderTaskGetReq, model *models.OrderTask) error {
	var err error
	var data models.OrderTask

	db := e.Orm.Model(&data).First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		e.Log.Errorf("db error:%s", err)
		return fmt.Errorf("查看对象不存在或无权查看")
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建OrderTask对象
func (e *OrderTask) Insert(c *dto.OrderTaskInsertReq) error {
	var err error
	var data models.OrderTask
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var existingTask models.OrderTask
	if err = tx.Where("name = ?", data.Name).First(&existingTask).Error; err == nil {
		// 如果存在相同标题的订单项，返回相应的错误消息
		return fmt.Errorf(fmt.Sprintf("Task with name '%v' already exists", existingTask.Name))
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

// Update OrderTask
func (e *OrderTask) Update(c *dto.OrderTaskUpdateReq) error {
	var err error
	var model = models.OrderTask{}

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
			e.Log.Errorf("task task with ID '%v' not exists", c.GetId())
			return fmt.Errorf("task task with ID '%v' not exists", c.GetId())
		}
		e.Log.Errorf("Error querying task with ID '%v': %s", c.GetId(), err)
		return fmt.Errorf("error querying  task with ID '%v': %s", c.GetId(), err)
	}

	// 判断是否需要执行更新操作
	if model.Name == c.Name && model.TaskType == c.TaskType && model.Interpreter == c.Interpreter && model.Content == c.Content && model.Description == c.Description {
		e.Log.Errorf(fmt.Sprintf("No changes for task with '%v'", model.Name))
		return fmt.Errorf("no changes for task with '%v'", model.Name)
	}

	// 更新订单项字段
	if err = tx.Model(&model).Updates(c).Error; err != nil {
		// 处理更新错误
		e.Log.Errorf("Failed to update task with title '%v': %s", c.Name, err)
		return fmt.Errorf("failed to update task with title '%v': %s", c.Name, err)
	}

	return nil
}

// Remove 删除OrderTask
func (e *OrderTask) Remove(d *dto.OrderTaskDeleteReq) error {
	var err error

	var data models.OrderTask
	// 查询要删除的任务
	if err = e.Orm.Model(&data).First(&data, d.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("task with ID '%v' not found", d.GetId())
		}
		e.Log.Errorf("Error querying task with ID '%v': %s", d.GetId(), err)
		return fmt.Errorf("error querying task with ID '%v': %s", d.GetId(), err)
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
