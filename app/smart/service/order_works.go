// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"fmt"
	"go-admin/app/smart/models"
	models2 "go-admin/common/models"
	"strconv"
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
	db := e.Orm.Where("currentHandlerId = ?", userid).Limit(limit).Offset(offset).Find(objects)

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
	db := e.Orm.Where("create_by = ? OR currentHandlerId = ?", userid, userid).Limit(limit).Offset(offset).Find(objects)

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

	var existingOrderWorks models.OrderWorks
	if err = tx.Where("title = ?", data.Title).First(&existingOrderWorks).Error; err == nil {
		// 如果存在相同标题的订单项，返回相应的错误消息
		return fmt.Errorf(fmt.Sprintf("order works with name '%v' already exists", existingOrderWorks.Title))
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果查询发生其他错误，返回错误
		return err
	}

	// Step 1: 查询模板表，获取绑定的流程ID
	var template models.FlowTemplates
	if err = tx.Where("name = ?", c.Template).First(&template).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("template with name '%v' not found", c.Template)
		}
		return err
	}

	// 获取绑定的流程 ID
	flowID := template.BindFlow

	// Step 2: 根据流程 ID 查询流程管理表，获取流程数据
	var flowManageData models.FlowManage
	if err = tx.Where("id = ?", flowID).First(&flowManageData).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("flow with ID '%v' not found", flowID)
		}
		return err
	}

	// 将默认流程数据给当前工单流程
	data.BindFlowData = flowManageData
	// 获取当前流程的名称，
	data.Process = flowManageData.Name
	// 获取所有的nodes数据
	nodes, ok := flowManageData.StrucTure["nodes"].([]interface{})
	if !ok {
		return fmt.Errorf("failed to parse nodes from structure")
	}

	// 遍历 nodes 获取第一个 userTask 的当前节点和处理人
	for _, nodeInterface := range nodes {
		node, ok := nodeInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if node["clazz"] == "userTask" {
			// 设置当前节点
			data.CurrentNode = node["label"].(string)

			assignValues, ok := node["assignValue"].([]interface{})
			if !ok || len(assignValues) == 0 {
				return fmt.Errorf("userTask node has no assignValue")
			}

			handlerID, err := strconv.Atoi(fmt.Sprint(assignValues[0]))
			if err != nil {
				return fmt.Errorf("failed to convert handler ID to int: %v", err)
			}

			// 获取姓名并设置到 CurrentHandler
			handlerName, err := GetHandlerNameByID(tx, handlerID)
			if err != nil {
				return fmt.Errorf("failed to get handler name by ID: %v", err)
			}

			data.CurrentHandlerID = handlerID
			data.CurrentHandler = handlerName
			break
		}
	}
	if data.CurrentHandlerID == 0 {
		return fmt.Errorf("userTask node not found in flow structure")
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

	// 显式地指定要更新的字段
	updates := map[string]interface{}{
		"priority":         c.Priority,
		"status":           c.Status,
		"currentHandler":   c.CurrentHandler,
		"currentHandlerId": c.CurrentHandlerId,
	}

	// 更新订单项字段
	if err = tx.Model(&model).Updates(updates).Error; err != nil {
		// 处理更新错误
		e.Log.Errorf("Failed to update order with title '%v': %s", c.Title, err)
		return fmt.Errorf("failed to update order with title '%v': %s", c.Title, err)
	}

	// 构建更新内容描述字符串
	updateDescriptionStr := BuildUpdateDescription(&model, c)

	historyReq := dto.OrderWorksHistReq{
		Title:          model.Title,
		Transfer:       "更新", // 假设流转操作类型保存在 ActionType 中
		Remark:         updateDescriptionStr,
		CurrentNode:    model.CurrentNode,
		Status:         c.Status,
		HandlerId:      c.ControlBy.UpdateBy,
		HandleTime:     models2.JSONTime(time.Now()),
		HandleDuration: calculateHandleDuration(model.ModelTime.UpdatedAt),
	}
	// 记录操作历史
	if err = RecordOperationHistory(tx, &historyReq); err != nil {
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
