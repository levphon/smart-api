package jobs

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/common/utils"
	"time"
)

// InitJob
// 需要将定义的struct 添加到字典中；
// 字典 key 可以配置到 自动任务 调用目标 中；
func InitJob() {
	host := ""
	db := sdk.Runtime.GetDbByKey(host)

	jobList = map[string]JobExec{
		"ExamplesOne": ExamplesOne{},
		"CheckAllMachines": CheckAllMachinesJob{
			MachineService: &utils.MachineService{
				Orm: db, // 使用指定的 DB 实例
			},
		},
	}
}

// ExamplesOne
// 新添加的job 必须按照以下格式定义，并实现Exec函数
type ExamplesOne struct {
}

// CheckAllMachinesJob 用于检查所有机器状态的定时任务
type CheckAllMachinesJob struct {
	MachineService *utils.MachineService // 传入你定义的服务
}

func (t CheckAllMachinesJob) Exec(arg interface{}) error {
	// 调用 CheckAllMachines 方法
	t.MachineService.CheckAllMachines()
	return nil
}

func (t ExamplesOne) Exec(arg interface{}) error {
	str := time.Now().Format(timeFormat) + " [INFO] JobCore ExamplesOne exec success"
	// TODO: 这里需要注意 Examples 传入参数是 string 所以 arg.(string)；请根据对应的类型进行转化；
	switch arg.(type) {

	case string:
		if arg.(string) != "" {
			fmt.Println("string", arg.(string))
			fmt.Println(str, arg.(string))
		} else {
			fmt.Println("arg is nil")
			fmt.Println(str, "arg is nil")
		}
		break
	}

	return nil
}
