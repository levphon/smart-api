package utils

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"smart-api/app/smart/models"
	"time"
)

type MachineService struct {
	Orm *gorm.DB
}

// 测试机器连接并更新心跳字段
func (s *MachineService) TestMachine(machine *models.ExecMachine) error {
	var err error

	// 使用 TCP 端口探测
	connTest := MachineConn{}
	err = connTest.testTCPPort(machine.Ip, machine.Port)
	if err != nil {
		// 连接失败，更新机器状态为 2
		updateErr := s.Orm.Model(machine).UpdateColumns(map[string]interface{}{
			"status": 2, // 2 表示连接失败
		}).Error
		if updateErr != nil {
			return fmt.Errorf("failed to update status for machine ID '%v' after connection failure: %v", machine.ID, updateErr)
		}
		// 返回连接失败错误
		return fmt.Errorf("TCP connection test failed: %v", err)
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
	// 暂时禁用日志
	s.Orm = s.Orm.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Silent),
	})

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
	// 如果需要恢复日志输出，可重新设置
	// s.Orm = s.Orm.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Info)})

}
