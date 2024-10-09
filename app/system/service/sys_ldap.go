package service

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-ldap/ldap/v3"
	"gorm.io/gorm"
	"net"
	"smart-api/app/system/models"
	"smart-api/app/system/service/dto"
	"smart-api/config"
	"time"
)

type SysLdap struct {
	Orm *gorm.DB
	Log *logger.Helper
}

// Ldap用户登陆
func (e SysLdap) LdapAuth(c *dto.LdapUserInsertReq) error {
	// ldap服务器认证以及判断该用户是否登陆过系统，如果没有登陆过会先自动注册
	err := e.checkLdapUser(c)
	if err != nil {
		return err
	}
	return nil
}

func (e SysLdap) checkLdapUser(c *dto.LdapUserInsertReq) error {
	var err error

	cfg := config.ExtConfig.Ldap
	var users models.SysUser
	tx := e.Orm.Debug().Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	l, err := e.ldapConnection(cfg)

	if err != nil {
		return err
	}

	// 构建搜索请求
	searchRequest := ldap.NewSearchRequest(
		cfg.SearchDomain,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(uid=%s)", c.Username),
		[]string{"dn"},
		nil,
	)
	// 发送搜索请求
	searchResult, err := l.Search(searchRequest)
	if err != nil {
		e.Log.Error(fmt.Sprintf("LDAP search failed: %s\n", err))
		return err
	}

	if len(searchResult.Entries) == 0 || len(searchResult.Entries) != 1 {
		e.Log.Error(fmt.Sprintf("LDAP %s User does not exist or too many entries returned", c.Username))
		return errors.New(fmt.Sprintf("LDAP %s User does not exist or too many entries returned", c.Username))
	}

	// 使用用户的 DN 绑定并进行密码验证
	userDN := searchResult.Entries[0].DN

	err = l.Bind(userDN, c.Password)

	if err != nil {
		e.Log.Error(fmt.Sprintf("LDAP %s authentication failed: %s", c.Username, err))
		return err
	}

	// LDAP认证通过后，查询数据库该用户是否存在，
	//如果不存在则走注册逻辑，注册后再走登陆逻辑，
	if err = tx.Where("username = ?", c.Username).First(&users).Error; err != nil {
		// 用户不存在, 执行用户注册的逻辑
		err = e.registerLdapUser(c)
		if err != nil {
			e.Log.Error(fmt.Sprintf("%v 注册失败.请联系管理员", c.Username))
			return fmt.Errorf("%v 注册失败.请联系管理员", c.Username)
		}
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (e SysLdap) registerLdapUser(c *dto.LdapUserInsertReq) error {
	var err error
	var data models.SysUser
	var i int64
	err = e.Orm.Model(&data).Where("username = ?", c.Username).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err = errors.New("用户名已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

func (e SysLdap) ldapConnection(cfg *config.Ldap) (*ldap.Conn, error) {
	ldapURL := fmt.Sprintf("ldap://%v:%v", cfg.Host, cfg.Port)
	var (
		l   *ldap.Conn
		err error
	)

	// 设置超时的 Dialer
	dialer := &net.Dialer{
		Timeout: 5 * time.Second, // 设置超时时间
	}

	if cfg.Tls {
		ldapURL = fmt.Sprintf("ldaps://%v:%v", cfg.Host, cfg.Port)
		tlsConf := &tls.Config{
			InsecureSkipVerify: true,
		}
		// 连接LDAP服务器
		l, err = ldap.DialURL(ldapURL, ldap.DialWithDialer(dialer), ldap.DialWithTLSConfig(tlsConf))
	} else {
		l, err = ldap.DialURL(ldapURL, ldap.DialWithDialer(dialer))
	}

	// 连接LDAP服务器时出现错误
	if err != nil {
		e.Log.Errorf(fmt.Sprintf("Failed to connect to LDAP server: %v", err))
		return nil, fmt.Errorf(fmt.Sprintf("Failed to connect to LDAP server: %v", err))
	}

	// SetTimeout设置发送请求后触发MessageTimeout的时间
	l.SetTimeout(5 * time.Second)

	// 绑定LDAP管理员账号
	err = l.Bind(cfg.BindDN, cfg.BindPassword)

	if err != nil {
		e.Log.Errorf("LDAP bind failed: %v", err)
		err = l.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return l, err
}
