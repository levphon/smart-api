package utils

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/cronjob"
	"go-admin/app/smart/models"
	"go-admin/config"
	"gorm.io/gorm"
	"time"

	"github.com/robfig/cron/v3"
)

type MachineService struct {
	Orm *gorm.DB
}

// 测试机器连接并更新心跳字段
func (s *MachineService) TestMachine(machine *models.ExecMachine) error {
	var err error
	// 获取加密配置
	cfg := config.ExtConfig.AesSecrets

	decryptPwd, err := Decrypt(machine.PassWord, cfg.Key)
	if err != nil {
		return fmt.Errorf("password decryption failed: %v", err)
	}

	// 复用之前的连接测试代码
	connTest := ConnectionTest{}
	err = connTest.TestConnection(machine.AuthType, machine.Ip, machine.Port, machine.UserName, decryptPwd, machine.PrivateKey)
	if err != nil {
		// 连接失败，更新机器状态为 2
		updateErr := s.Orm.Model(machine).UpdateColumns(map[string]interface{}{
			"status": 2, // 2 表示连接失败
		}).Error
		if updateErr != nil {
			return fmt.Errorf("failed to update status for machine ID '%v' after connection failure: %v", machine.ID, updateErr)
		}
		// 返回连接失败错误
		return fmt.Errorf("connection test failed: %v", err)
	}

	// 连接成功，更新心跳时间和状态为正常
	err = s.Orm.Model(machine).UpdateColumns(map[string]interface{}{
		"heartbeat": time.Now(), // 更新心跳时间
		"status":    1,          // 1 表示连接正常
	}).Error
	if err != nil {
		return fmt.Errorf("failed to update heartbeat for machine ID '%v': %v", machine.ID, err)
	}

	return nil
}

// 检测所有机器状态
func (s *MachineService) CheckAllMachines() {
	var machines []models.ExecMachine
	// 从数据库中获取所有机器
	err := s.Orm.Find(&machines).Error
	if err != nil {
		fmt.Printf("failed to retrieve machines: %v\n", err)
		return
	}

	// 逐台机器进行状态检查
	for _, machine := range machines {
		err = s.TestMachine(&machine)
		if err != nil {
			fmt.Printf("failed to test connection for machine ID '%v': %v\n", machine.ID, err)
		}
	}
}

func StartCronJob(dbs map[string]*gorm.DB) {
	for k, _ := range dbs {
		sdk.Runtime.SetCrontab(k, cronjob.NewWithSeconds())
	}

	// 获取默认数据库实例
	orm, ok := dbs[""]
	if !ok {
		panic("failed to get default DB from map")
	}

	service := MachineService{
		Orm: orm,
	}

	// 创建一个 cron 调度器
	c := cron.New()

	// 每 5 分钟检测一次所有机器状态
	_, err := c.AddFunc("*/5 * * * *", func() {
		service.CheckAllMachines()
	})
	if err != nil {
		fmt.Printf("failed to schedule task: %v\n", err)
		return
	}

	// 启动 cron 调度器
	c.Start()

	// 阻止程序退出，以确保定时任务可以持续运行
	select {}
}
