// @Author sunwenbo
// 2024/8/28 15:34
package service

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"time"
)

type ConnectionTest struct{}

func (c *ConnectionTest) TestConnection(authType int, ip string, port int, username, password, keyPath string) error {
	switch authType {
	case 1:
		// 通过用户名和密码进行SSH连接测试
		return c.testPasswordAuth(ip, port, username, password)
	case 2:
		// 通过公私钥进行SSH连接测试
		return c.testKeyAuth(ip, port, username, keyPath)
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

func (c *ConnectionTest) testKeyAuth(ip string, port int, username, keyPath string) error {
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return err
	}

	// 解析私钥
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
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
		return err
	}
	defer client.Close()

	return nil
}
