package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
	"smart-ops/common"
)

func LdapDelUser(l *ldap.Conn, cfg *common.Config, para *common.User) (result common.PostResponse) {

	// 管理员用户密码认证
	if para.AdminUser != "admin" || para.AdminPwd != cfg.Ldap.BindPassword {
		result.Code = 50001
		result.Data = fmt.Sprintf("LDAP permission management requires the authentication administrator user and password. 当前管理员用户为: %v", para.AdminUser)
		logrus.Error(fmt.Sprintf("LDAP permission management requires the authentication administrator user and password.当前管理员用户为: %v", para.AdminUser))
		return result
	}

	var userDN string
	if para.UserSecondOu == "sre" || para.UserSecondOu == "qa" || para.UserSecondOu == "ipo" {
		// 用户属性信息
		userDN = fmt.Sprintf("cn=%v,ou=%v,ou=users,%v", para.UserName, para.UserSecondOu, cfg.Ldap.SearchDomain) // 用户的DN
	} else {
		// 用户属性信息
		userDN = fmt.Sprintf("cn=%v,ou=%v,ou=%v,ou=users,%v", para.UserName, para.UserSecondOu, para.UserFirstOu, cfg.Ldap.SearchDomain) // 用户的DN
	}

	defer l.Close()

	// 删除用户
	delRequest := ldap.NewDelRequest(userDN, nil)
	err := l.Del(delRequest)

	if err != nil {
		result.Code = 50001
		result.Data = fmt.Sprintf("Failed to delete user:%v. reason: %v", para.UserName, err)
		logrus.Info(fmt.Sprintf("Failed to delete user:%v. reason: %v", para.UserName, err))
		return result
	}
	result.Code = 20000
	result.Data = fmt.Sprintf("delete user: %v successfully!", para.UserName)
	logrus.Info(fmt.Sprintf("delete user: %v successfully!\n", para.UserName))

	return result
}
