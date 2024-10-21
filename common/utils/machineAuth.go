// @Author sunwenbo
// 2024/8/28 15:34
package utils

import (
	"bytes"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

type MachineConn struct{}

func (c *MachineConn) TestConnection(authType string, ip string, port int, username, password, privateKey string) error {
	var err error
	switch authType {
	case "1":
		// 通过用户名和密码进行SSH连接测试
		err = c.testPasswordAuth(ip, port, username, password)
	case "2":
		// 通过公私钥进行SSH连接测试
		err = c.testKeyAuth(ip, port, username, privateKey)
	default:
		err = fmt.Errorf("invalid authentication type")
	}

	// 返回上层处理的错误
	if err != nil {
		return err
	}

	return nil
}

func (c *MachineConn) testPasswordAuth(ip string, port int, username, password string) error {
	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second, // 设置超时
	}

	// 拼接IP和端口
	host := fmt.Sprintf("%s:%d", ip, port)

	// 尝试连接SSH
	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		// 打印错误并返回
		fmt.Printf("Failed to dial SSH: %v\n", err)
		return err
	}

	// 只有当 client 不为 nil 时，才会进入 defer 语句
	defer client.Close()
	log.Info("SSH connection established")
	// 连接成功，返回 nil 表示没有错误
	return nil
}

// 使用私钥进行登陆时，首先需要将公钥放在内容放到 .ssh/authorized_keys 文件中，
// 然后需要将私钥提供给客户端，客户端就可以拿着私钥进行认证登陆 ssh -i id_rsa  root@10.50.183.112 -p 52829
func (c *MachineConn) testKeyAuth(ip string, port int, username, privateKey string) error {

	// 解析私钥
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// 拼接ip和端口
	host := fmt.Sprintf("%s:%d", ip, port)
	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	defer client.Close()

	return nil
}

func (c *MachineConn) testTCPPort(ip string, port int) error {
	address := fmt.Sprintf("%s:%d", ip, port)
	timeout := 3 * time.Second // 设置超时时间

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", address, err)
	}
	defer conn.Close()

	return nil
}

// ExecuteCommandWithParams 执行命令并传递参数
func (c *MachineConn) ExecuteCommandWithParams(client *ssh.Client, orderTitle, command, formDataJSON string) (string, string, error) {
	// 创建一个新的 SSH 会话
	session, err := client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	var stdoutBuf, stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf

	// 根据 orderTitle 和当前时间生成唯一的临时脚本文件名
	timeNow := time.Now().Format("20060102150405") // 格式化为 年月日时分秒
	tempFileName := fmt.Sprintf("/tmp/%s_%s.sh", orderTitle, timeNow)

	// 创建临时脚本内容
	tempScript := fmt.Sprintf(`
%s
# 动态生成的脚本，接收 $1 作为 JSON 参数
jsonData=$1
echo "Executing task with JSON data: $jsonData"
`, command)

	// 将脚本写入远程机器的临时文件
	if err := c.writeRemoteFile(client, tempFileName, tempScript); err != nil {
		return "", "", fmt.Errorf("failed to write remote script: %v", err)
	}

	// 创建一个新的 SSH 会话给脚本添加执行权限
	session, err = client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create session for chmod: %v", err)
	}
	defer session.Close()

	// 给临时脚本添加执行权限
	chmodCmd := fmt.Sprintf("chmod +x %s", tempFileName)
	if err := session.Run(chmodCmd); err != nil {
		return "", "", fmt.Errorf("failed to chmod script: %v", err)
	}

	// 创建一个新的 SSH 会话执行脚本
	session, err = client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create session for executing script: %v", err)
	}
	defer session.Close()

	// 执行脚本，并传递 formDataJSON 作为第一个参数 $1
	execCmd := fmt.Sprintf("%s '%s'", tempFileName, formDataJSON)
	fmt.Println("Executing command:", execCmd)
	if err := session.Run(execCmd); err != nil {
		return stdoutBuf.String(), stderrBuf.String(), fmt.Errorf("failed to execute command: %v", err)
	}

	// 创建一个新的 SSH 会话删除临时脚本
	session, err = client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create session for removing script: %v", err)
	}
	defer session.Close()

	// 删除临时脚本
	rmCmd := fmt.Sprintf("rm -f %s", tempFileName)
	if err := session.Run(rmCmd); err != nil {
		return stdoutBuf.String(), stderrBuf.String(), fmt.Errorf("failed to remove temp script: %v", err)
	}

	return stdoutBuf.String(), stderrBuf.String(), nil
}

// writeRemoteFile 在远程机器上写入文件
func (c *MachineConn) writeRemoteFile(client *ssh.Client, path string, content string) error {
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		return fmt.Errorf("failed to create SFTP client: %v", err)
	}
	defer sftpClient.Close()

	// 在远程创建文件
	file, err := sftpClient.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create remote file: %v", err)
	}
	defer file.Close()

	// 写入文件内容
	if _, err := file.Write([]byte(content)); err != nil {
		return fmt.Errorf("failed to write content to file: %v", err)
	}

	return nil
}
