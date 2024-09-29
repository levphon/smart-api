// @Author sunwenbo
// 2024/9/29
package service

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/smart/models"
	"go-admin/app/smart/service/dto"
	"gorm.io/gorm"
)

type UserFavorites struct {
	service.Service
}

// AddFavorite 添加收藏
func (e *UserFavorites) AddFavorite(userID int, req *dto.UserFavoriteInsertReq) error {

	var err error
	var data models.UserFavorites
	req.Generate(userID, &data) // 将请求数据转换为模型数据

	// 开始数据库事务
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback() // 如果有错误则回滚
		} else {
			tx.Commit() // 否则提交事务
		}
	}()

	// 检查该用户是否已收藏此工单
	var existingFavorite models.UserFavorites
	err = tx.Where("user_id = ? AND order_item_id = ?", userID, data.OrderItemID).First(&existingFavorite).Error
	if err == nil {
		// 如果查询到记录，说明已收藏
		e.Log.Errorf("User %d already favorited order item %d", userID, data.OrderItemID)
		return errors.New("您已经将此工单设置为收藏状态")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果查询出错，返回错误
		e.Log.Errorf("db error: %s", err)
		return err
	}

	// 插入新的收藏记录
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

// RemoveFavorite 取消收藏
func (e *UserFavorites) RemoveFavorite(userID int, req *dto.UserFavoriteDeleteReq) error {
	var err error

	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 取消收藏，删除满足 userID 和 orderItemID 的记录
	err = tx.Where("user_id = ? AND order_item_id = ?", userID, req.OrderItemID).Delete(&models.UserFavorites{}).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

// GetUserFavoritesByUserID 获取用户收藏的工单项
func (e *UserFavorites) GetUserFavoritesByUserID(userID int, models *[]models.UserFavorites) error {
	var err error

	// 查询收藏的工单项
	err = e.Orm.Where("user_id = ?", userID).Find(models).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		e.Log.Errorf("db error:%s", err)
		return fmt.Errorf("未找到收藏的工单")
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}
