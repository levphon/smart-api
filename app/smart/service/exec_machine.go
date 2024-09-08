// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"fmt"
	"go-admin/app/smart/models"
	common "go-admin/common/models"
	"go-admin/common/utils"
	"go-admin/config"
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
	connTest := utils.ConnectionTest{}
	err = connTest.TestConnection(data.AuthType, data.Ip, data.Port, data.UserName, data.PassWord, data.PrivateKey)
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

	// 加密密码 key
	cfg := config.ExtConfig.AesSecrets

	encryptedPassword, err := utils.Encrypt(data.PassWord, cfg.Key)

	// 加密密码
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.PassWord), bcrypt.DefaultCost)
	if err != nil {
		e.Log.Errorf("password encryption failed: %v", err)
		return fmt.Errorf("password encryption failed: %v", err)
	}
	data.PassWord = encryptedPassword // 使用加密后的密码

	data.Heartbeat = common.JSONTime(time.Now())
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

	// 构建一个更新数据的 map
	updates := map[string]interface{}{}

	// 手动检测每个字段的变化
	if model.Ip != c.Ip {
		updates["ip"] = c.Ip
	}
	if model.HostName != c.HostName {
		updates["hostname"] = c.HostName
	}
	if model.UserName != c.UserName {
		updates["username"] = c.UserName
	}
	if model.Port != c.Port {
		updates["port"] = c.Port
	}
	if model.Status != c.Status {
		updates["status"] = c.Status
	}
	if model.AuthType != c.AuthType {
		updates["auth_type"] = c.AuthType
	}
	if model.PrivateKey != c.PrivateKey {
		updates["private_key"] = c.PrivateKey
	}
	if model.Description != c.Description {
		updates["description"] = c.Description
	}

	if model.PassWord != c.PassWord {
		// 密码检测和加密
		cfg := config.ExtConfig.AesSecrets
		encryptedPassword, err := utils.Encrypt(c.PassWord, cfg.Key)

		if err != nil {
			e.Log.Errorf("password encryption failed: %v", err)
			return fmt.Errorf("password encryption failed: %v", err)
		}
		updates["password"] = encryptedPassword
	}

	// 如果没有字段发生变化
	if len(updates) == 0 {
		e.Log.Infof("No changes detected for machine with hostname '%v'", c.HostName)
		return fmt.Errorf("no changes for machine with hostname '%v'", c.HostName)
	}

	// 更新人更新
	updates["regenerator"] = c.Regenerator

	// 进行更新
	if err = tx.Model(&model).Updates(updates).Error; err != nil {
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

// 测试连接
func (e *ExecMachine) TestConn(d *dto.ExecMachineGetReq, model *models.ExecMachine) error {
	var machine models.ExecMachine

	// 根据 ID 查找机器信息
	err := e.Orm.First(&machine, d.GetId()).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("machine with ID '%v' not found: %s", d.GetId(), err)
			return fmt.Errorf("machine with ID '%v' not found", d.GetId())
		}
		e.Log.Errorf("error retrieving machine with ID '%v': %s", d.GetId(), err)
		return err
	}
	cfg := config.ExtConfig.AesSecrets
	password, err := utils.Decrypt(machine.PassWord, cfg.Key)
	if err != nil {
		e.Log.Errorf("password decryption failed: %v", err)
		return fmt.Errorf("password decryption failed: %v", err)
	}

	// 复用之前的连接测试代码
	connTest := utils.ConnectionTest{}

	go func() {
		err = connTest.TestConnection(machine.AuthType, machine.Ip, machine.Port, machine.UserName, password, machine.PrivateKey)
	}()

	if err != nil {
		e.Log.Errorf("connection test failed: %v", err)
		return err
	}

	// 更新数据库中的心跳时间
	err = e.Orm.Model(&machine).UpdateColumn("heartbeat", common.JSONTime(time.Now())).Error
	if err != nil {
		e.Log.Errorf("failed to update heartbeat for machine ID '%v': %v", d.GetId(), err)
		return fmt.Errorf("failed to update heartbeat for machine ID '%v': %v", d.GetId(), err)
	}

	return nil
}
