// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"fmt"
	"go-admin/app/smart/models"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/smart/service/dto"
)

// 定义映射英文到中文的优先级全局变量
var PriorityMap = map[string]string{
	"normal":      "一般",
	"urgent":      "紧急",
	"very-urgent": "非常紧急",
}

// 定义映射英文到中文的状态全局变量
var StatusMap = map[string]string{
	"under-way":          "进行中",
	"termination":        "已结束",
	"manual-termination": "手动结束",
	"reject":             "驳回",
}

type OrderWorks struct {
	service.Service
}

type OperationHistory struct {
	service.Service
}

// 分页获取OrderWorksry 所有的数据
func (e *OrderWorks) GetOrderWorksPage(pageNum, limit int, objects *[]models.OrderWorks) error {

	// 计算偏移量
	offset := (pageNum - 1) * limit

	// 查询并分页获取订单项数据
	db := e.Orm.Limit(limit).Offset(offset).Find(objects)

	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询工单失败: %s", err)
		return fmt.Errorf("分页查询工单失败: %s", err)
	}
	return nil
}

// Get 获取我的待办
func (e *OrderWorks) MyBackGet(pageNum, limit int, objects *[]models.OrderWorks, userid int) error {
	// 计算偏移量
	offset := (pageNum - 1) * limit
	// 查询并分页获取订单项数据
	db := e.Orm.Where("currentHandler = ?", userid).Limit(limit).Offset(offset).Find(objects)

	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询工单失败: %s", err)
		return fmt.Errorf("分页查询工单失败: %s", err)
	}
	return nil
}

// Get 获取我创建的
func (e *OrderWorks) CreatedByMe(pageNum, limit int, objects *[]models.OrderWorks, userid int) error {
	// 计算偏移量
	offset := (pageNum - 1) * limit
	// 查询并分页获取订单项数据
	db := e.Orm.Where("create_by = ?", userid).Limit(limit).Offset(offset).Find(objects)

	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询工单失败: %s", err)
		return fmt.Errorf("分页查询工单失败: %s", err)
	}
	return nil
}

// Get 获取与我相关的
func (e *OrderWorks) RelatedToMe(pageNum, limit int, objects *[]models.OrderWorks, userid int) error {
	// 计算偏移量
	offset := (pageNum - 1) * limit
	// 查询并分页获取订单项数据
	db := e.Orm.Where("create_by = ? OR currentHandler = ?", userid, userid).Limit(limit).Offset(offset).Find(objects)

	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询工单失败: %s", err)
		return fmt.Errorf("分页查询工单失败: %s", err)
	}
	return nil
}

// 根据ID获取工单
func (e *OrderWorks) Get(d *dto.OrderWorksGetReq, model *models.OrderWorks) error {
	var err error
	var data models.OrderWorks

	db := e.Orm.Model(&data).
		First(model, d.GetId())
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

// Insert 创建OrderWorks对象
func (e *OrderWorks) Insert(c *dto.OrderWorksInsertReq) error {
	var err error
	var data models.OrderWorks
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
	var existingOrderWorks models.OrderWorks
	if err = tx.Where("title = ?", data.Title).First(&existingOrderWorks).Error; err == nil {
		// 如果存在相同标题的订单项，返回相应的错误消息
		return fmt.Errorf(fmt.Sprintf("order works with name '%v' already exists", existingOrderWorks.Title))
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

// Update OrderWorksry
func (e *OrderWorks) Update(c *dto.OrderWorksUpdateReq) error {
	var err error
	var model = models.OrderWorks{}

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
			e.Log.Errorf("order works with ID '%v' not exists", c.GetId())
			return fmt.Errorf("order works with ID '%v' not exists", c.GetId())
		}
		e.Log.Errorf("Error querying order works with ID '%v': %s", c.GetId(), err)
		return fmt.Errorf("error querying order works with ID '%v': %s", c.GetId(), err)
	}

	// 判断是否需要执行更新操作
	if model.Priority == c.Priority && model.Status == c.Status && model.CurrentHandler == c.CurrentHandler {
		e.Log.Errorf(fmt.Sprintf("No changes for order with '%v'", model.Title))
		return fmt.Errorf("no changes for order with '%v'", model.Title)
	}
	// 构建更新内容描述字符串
	updateDescriptionStr := buildUpdateDescription(&model, c)

	// 显式地指定要更新的字段
	updates := map[string]interface{}{
		"priority":       c.Priority,
		"status":         c.Status,
		"currentHandler": c.CurrentHandler,
	}

	// 更新订单项字段
	if err = tx.Model(&model).Updates(updates).Error; err != nil {
		// 处理更新错误
		e.Log.Errorf("Failed to update order with title '%v': %s", c.Title, err)
		return fmt.Errorf("failed to update order with title '%v': %s", c.Title, err)
	}

	// 记录操作历史
	if err = recordOperationHistory(tx, c, updateDescriptionStr); err != nil {
		e.Log.Errorf("Failed to record operation history: %s", err)
		return fmt.Errorf("failed to record operation history: %s", err)
	}

	return nil
}

// Remove 删除SysOrderWorksry
func (e *OrderWorks) Remove(d *dto.OrderWorksDeleteReq) error {
	var err error

	var data models.OrderWorks
	// 查询要删除的订单项
	if err = e.Orm.Model(&data).First(&data, d.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order with ID '%v' not found", d.GetId())
		}
		e.Log.Errorf("Error querying order with ID '%v': %s", d.GetId(), err)
		return fmt.Errorf("error querying order with ID '%v': %s", d.GetId(), err)
	}

	// 检查当前工单的状态是否为进行中 under-way，如果为进行中则返回报错，提示不能删除
	if data.Status == "under-way" {
		return errors.New("当前工单进行中，无法删除")
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

func buildUpdateDescription(model *models.OrderWorks, c *dto.OrderWorksUpdateReq) string {
	var updateDescription []string

	if model.Priority != c.Priority {
		oldPriority := PriorityMap[model.Priority]
		newPriority := PriorityMap[c.Priority]
		updateDescription = append(updateDescription, fmt.Sprintf("优先级: '%v' => '%v'", oldPriority, newPriority))
	}
	if model.Status != c.Status {
		oldStatus := StatusMap[model.Status]
		newStatus := StatusMap[c.Status]
		updateDescription = append(updateDescription, fmt.Sprintf("状态: '%v' => '%v'", oldStatus, newStatus))
	}
	if model.CurrentHandler != c.CurrentHandler {
		oldHandler := model.CurrentHandler
		newHandler := c.CurrentHandler
		updateDescription = append(updateDescription, fmt.Sprintf("处理人: '%v' => '%v'", oldHandler, newHandler))
	}

	return strings.Join(updateDescription, ", ")
}

func recordOperationHistory(tx *gorm.DB, c *dto.OrderWorksUpdateReq, updateDescriptionStr string) error {
	var operation = models.OperationHistory{
		Title:          c.Title,
		NodeName:       c.CurrentNode,
		Transfer:       "修改工单信息",
		Remark:         fmt.Sprintf("更新：%v", updateDescriptionStr),
		Handler:        c.ControlBy.UpdateBy,
		HandleTime:     time.Now(),
		HandleDuration: calculateHandleDuration(c.ModelTime.UpdatedAt),
		ControlBy:      c.ControlBy, // 可以根据实际情况更改为具体处理人
		ModelTime:      c.ModelTime,
	}
	return tx.Create(&operation).Error

}

// 计算处理时长
func calculateHandleDuration(updatedAt time.Time) int64 {
	currentTime := time.Now()
	duration := currentTime.Sub(updatedAt)
	return int64(duration.Seconds())
}

// 根据ID获取工单
func (h *OperationHistory) GetHistory(pageNum, limit int, objects *[]models.OperationHistory, title string) error {
	// 计算偏移量
	offset := (pageNum - 1) * limit
	// 查询并分页获取订单项数据
	db := h.Orm.Where("title = ?", title).Limit(limit).Offset(offset).Find(objects)

	if err := db.Error; err != nil {
		h.Log.Errorf("历史操作失败: %s", err)
		return fmt.Errorf("历史操作失败: %s", err)
	}
	return nil
}
