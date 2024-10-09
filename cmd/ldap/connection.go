package ldap

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"smart-api/config"
	"time"
)

func LdapConnection1() (*ldap.Conn, error) {
	cfg := config.Ldap{}

	ldapURL := fmt.Sprintf("ldap://%v:%v", cfg.Host, cfg.Port)
	var (
		l   *ldap.Conn
		err error
	)

	if cfg.Tls {
		ldapURL = fmt.Sprintf("ldaps://%v:%v", cfg.Host, cfg.Port)
		tlsConf := &tls.Config{
			InsecureSkipVerify: true,
		}
		l, err = ldap.DialURL(ldapURL, ldap.DialWithTLSConfig(tlsConf))
	} else {
		l, err = ldap.DialURL(ldapURL)
	}

	// 连接LDAP服务器
	if err != nil {
		e.Log.Errorf(fmt.Sprintf("Failed to connect to LDAP server: %v", err))
	}

	// 绑定LDAP管理员账号
	err = l.Bind(cfg.BindDN, cfg.BindPassword)

	if err != nil {
		e.Log.Errorf("LDAP bind failed: %v", err)
	}
	l.SetTimeout(5 * time.Second)

	return l, err
}
