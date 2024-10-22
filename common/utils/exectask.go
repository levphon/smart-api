// @Author sunwenbo
// 2024/10/22 18:56
package utils

import (
	"bytes"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// ExecuteCommandWithParams 执行命令并传递参数
func (c *MachineConn) ExecuteCommandWithParams(client *ssh.Client, orderTitle, command, formDataJSON string) (string, error) {

	// 创建 TaskLogger 实例
	taskLog, err := NewTaskLogger(orderTitle)
	if err != nil {
		return "", err
	}
	defer taskLog.Close() // 确保在函数结束时关闭日志文件

	// 创建一个新的 SSH 会话
	session, err := client.NewSession()
	if err != nil {
		taskLog.Log(fmt.Sprintf("Failed to create session: %v", err))
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	taskLog.Log("SSH session created successfully.")

	// 缓存 stdout 和 stderr
	var stdoutBuf, stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf

	// 根据 orderTitle 生成唯一的临时脚本文件名
	tempFileName := fmt.Sprintf("/tmp/%s.sh", orderTitle)
	taskLog.Log(fmt.Sprintf("Generated file name: %s", tempFileName))

	// 创建临时脚本内容
	tempScript := command
	taskLog.Log(fmt.Sprintf("Generated script: \n%s", tempScript))

	// 将脚本写入远程机器的临时文件
	if err := c.writeRemoteFile(client, tempFileName, tempScript, taskLog.Log); err != nil {
		taskLog.Log(fmt.Sprintf("Failed to write remote script: %v", err))
		return "", fmt.Errorf("failed to write remote script: %v", err)
	}
	taskLog.Log("Remote script written successfully.")

	// 给临时脚本添加执行权限
	chmodCmd := fmt.Sprintf("chmod +x %s", tempFileName)
	taskLog.Log(fmt.Sprintf("Running chmod command: %s", chmodCmd))
	if err := session.Run(chmodCmd); err != nil {
		taskLog.Log(fmt.Sprintf("Failed to chmod script: %v", err))
		return "", fmt.Errorf("failed to chmod script: %v", err)
	}
	taskLog.Log("Script chmod success.")

	// 重新创建 session 执行脚本，确保之前的会话不受影响
	session, err = client.NewSession()
	if err != nil {
		taskLog.Log(fmt.Sprintf("Failed to create session for script execution: %v", err))
		return "", fmt.Errorf("failed to create session for script execution: %v", err)
	}
	defer session.Close()

	// 绑定 stdout 和 stderr
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf

	// 执行脚本，并传递 formDataJSON 作为参数
	execCmd := fmt.Sprintf("%s '%s'", tempFileName, formDataJSON)
	taskLog.Log(fmt.Sprintf("Executing command: %s", execCmd))

	if err := session.Run(execCmd); err != nil {
		// 记录执行失败时的输出信息
		taskLog.Log(fmt.Sprintf("Failed to execute command: %v", err))
		taskLog.Log(fmt.Sprintf("stderr: %s", stderrBuf.String()))
		return stderrBuf.String(), fmt.Errorf("failed to execute command: %v", err)
	}
	// 获取标准输出
	stdoutStr := stdoutBuf.String()
	stderrStr := stderrBuf.String()

	// 检查标准错误输出
	if stderrStr != "" {
		taskLog.Log(fmt.Sprintf("Command executed with warnings/errors in stderr:\nstderr: %s", stderrStr))
	}

	taskLog.Log(fmt.Sprintf("Command executed successfully.\nstdout: %s", stdoutStr))

	return stdoutStr, nil

}

// writeRemoteFile 在远程机器上写入文件
func (c *MachineConn) writeRemoteFile(client *ssh.Client, path string, content string, logMessage func(string)) error {
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		logMessage(fmt.Sprintf("Failed to create SFTP client: %v", err))
		return fmt.Errorf("failed to create SFTP client: %v", err)
	}
	defer sftpClient.Close()

	// 在远程创建文件
	file, err := sftpClient.Create(path)
	if err != nil {
		logMessage(fmt.Sprintf("Failed to create remote file: %v", err))
		return fmt.Errorf("failed to create remote file: %v", err)
	}
	defer file.Close()

	// 写入文件内容
	if _, err := file.Write([]byte(content)); err != nil {
		logMessage(fmt.Sprintf("Failed to write content to remote file: %v", err))
		return fmt.Errorf("failed to write content to file: %v", err)
	}
	logMessage("File written successfully.")

	return nil
}
