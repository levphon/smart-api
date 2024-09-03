// @Author sunwenbo
// 2024/8/28 15:34
package utils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
)

type ConnectionTest struct{}

func (c *ConnectionTest) TestConnection(authType string, ip string, port int, username, password, privateKey string) error {
	switch authType {
	case "1":
		// 通过用户名和密码进行SSH连接测试
		return c.testPasswordAuth(ip, port, username, password)
	case "2":
		// 通过公私钥进行SSH连接测试
		return c.testKeyAuth(ip, port, username, privateKey)
	default:
		return fmt.Errorf("invalid authentication type")
	}
}

func (c *ConnectionTest) testPasswordAuth(ip string, port int, username, password string) error {
	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// 拼接ip和端口
	host := fmt.Sprintf("%s:%d", ip, port)
	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return err
	}
	defer client.Close()

	return nil
}

// 使用私钥进行登陆时，首先需要将公钥放在内容放到 .ssh/authorized_keys 文件中，
// 然后需要将私钥提供给客户端，客户端就可以拿着私钥进行认证登陆 ssh -i id_rsa  root@10.50.183.112 -p 52829
func (c *ConnectionTest) testKeyAuth(ip string, port int, username, privateKey string) error {

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
