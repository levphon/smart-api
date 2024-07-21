// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"fmt"
	"go-admin/app/smart/models"
	models2 "go-admin/common/models"
	"time"

	"gorm.io/gorm"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/smart/service/dto"
)

type OrderItems struct {
	service.Service
}

// 分页获取OrderItems 所有的数据
func (e *OrderItems) GetOrderItemsPage(pageNum, limit int, objects *[]models.OrderItems) error {
	// 计算偏移量
	offset := (pageNum - 1) * limit

	// 查询并分页获取订单项数据
	db := e.Orm.Limit(limit).Offset(offset).Find(objects)
	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询订单项失败: %s", err)
		return fmt.Errorf("分页查询订单项失败: %s", err)
	}
	return nil
}

// Get 获取OrderItems对象
func (e *OrderItems) Get(d *dto.OrderItemsGetReq, model *models.OrderItems) error {
	var err error
	var data models.OrderItems

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

// Insert 创建OrderItems对象
func (e *OrderItems) Insert(c *dto.OrderItemsInsertReq) error {
	var err error
	var data models.OrderItems
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	data.CreatedAt = models2.JSONTime(time.Now())
	data.UpdatedAt = models2.JSONTime(time.Now())
	var existingOrderItems models.OrderItems

	if err = tx.Where("title = ?", data.Title).First(&existingOrderItems).Error; err == nil {
		// 如果存在相同标题的订单项，返回相应的错误消息
		return fmt.Errorf(fmt.Sprintf("Order with title '%v' already exists", existingOrderItems.Title))
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果查询发生其他错误，返回错误
		return err
	}
	// 检查绑定模板字段是否有改变
	if existingOrderItems.BindTempLate != data.BindTempLate {
		// 如果绑定模板字段发生了变化，则更新 FlowTemplate 数据库中的 bindCount 字段
		// 这里假设你有一个名为 updateBindCount 的函数来更新 bindCount
		if err = updateBindCount(existingOrderItems.BindTempLate, data.BindTempLate, tx); err != nil {
			// 处理更新 bindCount 错误
			e.Log.Errorf("db error:%s", err)
			return fmt.Errorf("db error:%s", err)
		}
		data.Link = fmt.Sprintf("/order/classify/%s", data.BindTempLate)
	}

	// 自动拼接链接
	data.Link = fmt.Sprintf("/order/classify/%s", data.BindTempLate)

	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Update OrderItems
func (e *OrderItems) Update(c *dto.OrderItemsUpdateReq) error {
	var err error
	var model = models.OrderItems{}
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
			e.Log.Errorf("Order with ID '%v' not exists", c.GetId())
			return fmt.Errorf("order with ID '%v' not exists", c.GetId())
		}
		e.Log.Errorf("Error querying order with ID '%v': %s", c.GetId(), err)
		return fmt.Errorf("error querying order with ID '%v': %s", c.GetId(), err)
	}

	// 检查绑定模板字段是否有改变
	if c.BindTempLate == "" {
		// 解绑当前模板，模板的bindCount字段减1
		if err = updateBindCount(model.BindTempLate, "", tx); err != nil {
			e.Log.Errorf("Failed to update bindCount when unbinding template: %v", err)
			return fmt.Errorf("failed to update bindCount when unbinding template: %v", err)
		}
		c.Link = ""
	} else if model.BindTempLate != c.BindTempLate {
		// 更新 FlowTemplate 数据库中的 bindCount 字段
		if err = updateBindCount(model.BindTempLate, c.BindTempLate, tx); err != nil {
			e.Log.Errorf("Failed to update bindCount: %v", err)
			return fmt.Errorf("failed to update bindCount: %v", err)
		}
		c.Link = fmt.Sprintf("/order/classify/%s", c.BindTempLate)
	}

	c.Generate(&model)

	// 执行更新操作
	db := tx.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateOrderItems error: %s", err)
		return err
	}

	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysOrderItems
func (e *OrderItems) Remove(d *dto.OrderItemsDeleteReq) error {
	var err error
	var data models.OrderItems
	// 查询要删除的订单项
	if err = e.Orm.Model(&data).First(&data, d.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order with ID '%v' not found", d.GetId())
		}
		e.Log.Errorf("Error querying order with ID '%v': %s", d.GetId(), err)
		return err
	}

	// 检查是否绑定了模板
	if data.BindTempLate != "" {
		return errors.New("该订单项已绑定模板，请解绑后再删除")
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

func updateBindCount(oldBind, newBind string, db *gorm.DB) error {
	//// 更新 oldBind 对应的 FlowTemplate 的 bindCount 字段减 1，确保 bindCount 不会变成负数
	//if err := db.Model(&models.FlowTemplates{}).Where("name = ?", oldBind).
	//	Update("bindCount", gorm.Expr("CASE WHEN bindCount > 0 THEN bindCount - 1 ELSE 0 END")).Error; err != nil {
	//	return err
	//}
	//
	//// 更新 newBind 对应的 FlowTemplate 的 bindCount 字段加 1
	//if err := db.Model(models.FlowTemplates{}).Where("name = ?", newBind).Update("bindCount", gorm.Expr("bindCount + ?", 1)).Error; err != nil {
	//	return err
	//}
	// 如果 oldBind 不为空，更新对应的 FlowTemplate 的 bindCount 字段减 1
	if oldBind != "" {
		if err := db.Model(&models.FlowTemplates{}).Where("name = ?", oldBind).
			Update("bindCount", gorm.Expr("CASE WHEN bindCount > 0 THEN bindCount - 1 ELSE 0 END")).Error; err != nil {
			return err
		}
	}

	// 如果 newBind 不为空，更新对应的 FlowTemplate 的 bindCount 字段加 1
	if newBind != "" {
		if err := db.Model(&models.FlowTemplates{}).Where("name = ?", newBind).
			Update("bindCount", gorm.Expr("bindCount + ?", 1)).Error; err != nil {
			return err
		}
	}

	return nil
}
