// @Author sunwenbo
// 2024/10/22 16:41
package utils

import (
	"fmt"
	"os"
	"time"
)

// TaskLogger 用于记录任务执行的日志
type TaskLogger struct {
	logFile *os.File
}

// NewTaskLogger 创建一个新的 TaskLogger 实例
func NewTaskLogger(orderTitle string, hisTaskUUId string) (*TaskLogger, error) {
	// 指定日志文件路径
	logDir := "./logs/task-log"

	// 获取当前日期并格式化为字符串
	currentDateTime := time.Now().Format("200601021504") // YYYYMMDDHHMM 格式

	logFilePath := fmt.Sprintf("%s/%s-%s_%s.log", logDir, orderTitle, currentDateTime, hisTaskUUId)

	// 确保日志目录存在，如果不存在则创建
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	// 打开日志文件，确保以追加模式打开
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	return &TaskLogger{logFile: logFile}, nil
}

// Log 写入日志信息
func (t *TaskLogger) Log(message string) {
	// 使用更高精度的时间戳
	_, err := t.logFile.WriteString(fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), message))
	if err != nil {
		fmt.Printf("Failed to write log: %v\n", err)
	}
}

// Close 关闭日志文件
func (t *TaskLogger) Close() {
	if err := t.logFile.Close(); err != nil {
		fmt.Printf("Failed to close log file: %v\n", err)
	}
}
