// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"gorm.io/gorm"
	"smart-api/app/smart/models"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"smart-api/app/smart/service/dto"
)

type OrderComment struct {
	service.Service
}

// GetCommentsByOrderID 查询工单的所有留言
func (e *OrderComment) GetCommentsByOrderID(orderID int) ([]models.OrderComment, error) {
	var comments []models.OrderComment
	err := e.Orm.Where("order_id = ?", orderID).
		Order("created_at ASC"). // 按创建时间升序排列
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// Insert 创建留言
func (e *OrderComment) Insert(req *dto.OrderCommentInsertReq) error {
	var err error
	var data models.OrderComment

	// 生成数据模型
	req.Generate(&data)

	// 开始数据库事务
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback() // 回滚事务
		} else {
			tx.Commit() // 提交事务
		}
	}()

	// 检查 OrderID 是否存在
	var orderWorks models.OrderWorks
	err = tx.Model(&models.OrderWorks{}).Where("id = ?", data.OrderID).First(&orderWorks).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("Order ID %d 不存在", data.OrderID)
			return errors.New("工单不存在")
		}
		e.Log.Errorf("检查 Order ID 时发生错误: %s", err)
		return err
	}

	// 插入留言
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("插入留言时发生数据库错误: %s", err)
		return err
	}

	return nil
}

// 更新留言
func (e *OrderComment) Update(req *dto.OrderCommentUpdateReq) error {
	var err error

	// 开始数据库事务
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback() // 回滚事务
		} else {
			tx.Commit() // 提交事务
		}
	}()

	// 1. 检查工单是否存在
	var orderWork models.OrderWorks
	err = tx.Where("id = ?", req.OrderID).First(&orderWork).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("工单 ID %d 不存在", req.OrderID)
			return errors.New("工单不存在")
		}
		e.Log.Errorf("查询工单时发生错误: %s", err)
		return err
	}

	// 2. 检查留言是否存在
	var existingComment models.OrderComment
	err = tx.Where("id = ?", req.ID).First(&existingComment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("留言 ID %d 不存在", req.ID)
			return errors.New("留言不存在")
		}
		e.Log.Errorf("查询留言时发生错误: %s", err)
		return err
	}

	// 3. 更新留言字段
	existingComment.Comments = req.Comments
	existingComment.ParentID = req.ParentID // 只在需要更新父留言时设置

	// 4. 更新数据库记录
	err = tx.Save(&existingComment).Error
	if err != nil {
		e.Log.Errorf("更新留言时发生数据库错误: %s", err)
		return err
	}

	return nil
}

// Remove 删除留言
func (e *OrderComment) Remove(req *dto.OrderCommentDeleteReq) error {
	var err error

	// 开始数据库事务
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback() // 回滚事务
		} else {
			tx.Commit() // 提交事务
		}
	}()

	// 1. 检查留言是否存在
	var existingComment models.OrderComment
	err = tx.Where("id = ?", req.ID).First(&existingComment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("留言 ID %d 不存在", req.ID)
			return errors.New("留言不存在")
		}
		e.Log.Errorf("查询留言时发生错误: %s", err)
		return err
	}

	// 2. 删除留言
	err = tx.Delete(&existingComment).Error
	if err != nil {
		e.Log.Errorf("删除留言时发生数据库错误: %s", err)
		return err
	}

	return nil
}
