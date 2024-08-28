// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"fmt"
	"go-admin/app/smart/models"
	"gorm.io/gorm"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/smart/service/dto"
)

type ExecMachine struct {
	service.Service
}

// 分页获取Ordermachine 所有的数据
func (e *ExecMachine) GetExecMachinePage(pageNum, limit int, objects *[]models.ExecMachine, total *int64) error {
	// 计算偏移量
	offset := (pageNum - 1) * limit

	// 查询总数，不应用分页限制
	err := e.Orm.Model(&models.ExecMachine{}).Count(total).Error
	if err != nil {
		e.Log.Errorf("查询总数失败: %s", err)
		return fmt.Errorf("查询总数失败: %s", err)
	}

	// 查询并分页获取订单项数据
	db := e.Orm.Limit(limit).Offset(offset).Find(objects)

	if err := db.Error; err != nil {
		e.Log.Errorf("执行节点查询失败: %s", err)
		return fmt.Errorf("执行节点查询失败: %s", err)
	}
	return nil
}

// 根据ID获取机器信息
func (e *ExecMachine) Get(d *dto.ExecMachineGetReq, model *models.ExecMachine) error {
	var err error
	var data models.ExecMachine

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

// Insert 创建ExecMachine对象
func (e *ExecMachine) Insert(c *dto.ExecMachineInsertReq) error {
	var err error
	var data models.ExecMachine
	c.Generate(&data)

	// 连接测试
	connTest := ConnectionTest{}
	err = connTest.TestConnection(data.AuthType, data.Ip, data.Port, data.UserName, data.PassWord, data.KeyPath)
	if err != nil {
		e.Log.Errorf("connection test failed: %v", err)
		return fmt.Errorf("connection test failed: %v", err)
	}

	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var existingMachine models.ExecMachine
	if err = tx.Where("ip = ?", data.Ip).First(&existingMachine).Error; err == nil {
		e.Log.Errorf("machine with name '%v' already exists", existingMachine.HostName)
		return fmt.Errorf(fmt.Sprintf("machine with name '%v' already exists", existingMachine.HostName))
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	data.Heartbeat = time.Now()
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Update machine
func (e *ExecMachine) Update(c *dto.ExecMachineUpdateReq) error {
	var err error
	var model = models.ExecMachine{}

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
			e.Log.Errorf("machine with ID '%v' not exists", c.GetId())
			return fmt.Errorf("machine with ID '%v' not exists", c.GetId())
		}
		e.Log.Errorf("Error querying machine with ID '%v': %s", c.GetId(), err)
		return fmt.Errorf("error querying  machine with ID '%v': %s", c.GetId(), err)
	}

	// 判断是否需要执行更新操作
	if model.HostName == c.HostName && model.Ip == c.Ip && model.UserName == c.UserName && model.PassWord == c.PassWord && model.Port == c.Port &&
		model.Status == c.Status && model.AuthType == c.AuthType && model.KeyPath == c.KeyPath && model.Description == c.Description {
		e.Log.Errorf(fmt.Sprintf("No changes for machine with '%v'", model.HostName))
		return fmt.Errorf("no changes for machine with '%v'", model.HostName)
	}

	// 手动映射要更新的字段
	model.HostName = c.HostName
	model.Ip = c.Ip
	model.UserName = c.UserName
	model.PassWord = c.PassWord
	model.Port = c.Port
	model.Status = c.Status
	model.AuthType = c.AuthType
	model.KeyPath = c.KeyPath
	model.Description = c.Description

	// 更新所有字段（包括零值）
	if err = tx.Model(&model).Updates(model).Error; err != nil {
		e.Log.Errorf("Failed to update machine with ID '%v': %s", c.GetId(), err)
		return fmt.Errorf("failed to update machine with ID '%v': %s", c.GetId(), err)
	}

	return nil
}

// Remove 删除Ordermachine
func (e *ExecMachine) Remove(d *dto.ExecMachineDeleteReq) error {
	var err error

	var data models.ExecMachine
	// 查询要删除的任务
	if err = e.Orm.Model(&data).First(&data, d.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("machine with ID '%v' not found", d.GetId())
		}
		e.Log.Errorf("Error querying machine with ID '%v': %s", d.GetId(), err)
		return fmt.Errorf("error querying machine with ID '%v': %s", d.GetId(), err)
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
