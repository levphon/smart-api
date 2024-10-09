// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"
	"smart-api/app/smart/models"
	"smart-api/app/smart/service/dto"
)

// OrderRating 服务结构体
type OrderRating struct {
	service.Service
}

// GetRatingByOrderID 根据工单ID获取评分
func (e *OrderRating) GetRatingByOrderID(orderID int) (*models.OrderRating, error) {
	var rating models.OrderRating
	err := e.Orm.Where("order_id = ?", orderID).First(&rating).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有找到记录，返回 nil 和一个自定义的错误
			return nil, nil // 返回 nil 以指示没有评分
		}
		return nil, err // 返回其他错误
	}
	return &rating, nil // 返回找到的评分
}

// Insert 插入评分
func (e *OrderRating) Insert(req *dto.OrderRatingInsertReq) error {
	var err error
	var data models.OrderRating
	req.Generate(&data) // 将请求数据转换为模型数据

	// 开始数据库事务
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback() // 如果有错误则回滚
		} else {
			tx.Commit() // 否则提交事务
		}
	}()

	// 查询是否已有评分，如果已有评分则返回错误
	var existingRating models.OrderRating
	err = tx.Where("order_id = ?", data.OrderID).First(&existingRating).Error
	if err == nil {
		return errors.New("该工单已存在评分")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 查询 order_works 表获取 bindFlowData
	var orderWork models.OrderWorks // 假设你的 OrderWork 模型已定义
	err = tx.Where("id = ?", data.OrderID).First(&orderWork).Error
	if err != nil {
		return err // 处理查询错误
	}

	// 默认将 Taskhandler 赋为空值
	data.Taskhandler = 0

	// 获取 bindFlowData 中的 nodes 列表
	if nodes, ok := orderWork.BindFlowData.StrucTure["nodes"].([]interface{}); ok {
		// 遍历所有节点
		for _, node := range nodes {
			nodeMap, ok := node.(map[string]interface{})
			if !ok {
				continue
			}

			// 检查节点类型是否为 "receiveTask"
			if nodeMap["clazz"] == "receiveTask" {
				// 提取 assignValue
				if assignValues, ok := nodeMap["assignValue"].([]interface{}); ok && len(assignValues) > 0 {
					if handler, ok := assignValues[0].(float64); ok { // 假设 assignValue 是数字
						data.Taskhandler = int(handler) // 设置 Taskhandler 值
					}
				} else {
					return errors.New("未找到 assignValue 值")
				}
			}
		}
	} else {
		return errors.New("未找到 nodes 列表")
	}

	// 插入新的评分记录
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

// Update 更新评分
func (e *OrderRating) Update(req *dto.OrderRatingUpdateReq) error {
	var err error
	var data models.OrderRating
	req.Generate(&data) // 将请求数据转换为模型数据

	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 根据 orderID 查询评分记录
	var existingRating models.OrderRating
	err = tx.Where("order_id = ?", data.OrderID).First(&existingRating).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评分记录不存在")
		}
		return err
	}

	// 更新评分记录
	err = tx.Model(&existingRating).Updates(map[string]interface{}{
		"rating": data.Ratings,
	}).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

// Remove 删除评分
func (e *OrderRating) Remove(req *dto.OrderRatingDeleteReq) error {
	var err error

	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 删除评分
	err = tx.Delete(&models.OrderRating{}, req.ID).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}
