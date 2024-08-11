package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
	"smart-ops/common"
)

func LdapUpdateUserPdw(l *ldap.Conn, cfg *common.Config, para *common.UpdateUserPdw) (result common.PostResponse) {
	defer l.Close()
	var userDN string

	if para.UserSecondOu == "sre" || para.UserSecondOu == "qa" || para.UserSecondOu == "ipo" {
		userDN = fmt.Sprintf("cn=%v,ou=%v,ou=users,%v", para.UserName, para.UserSecondOu, cfg.Ldap.SearchDomain) // 用户的DN
	} else {
		userDN = fmt.Sprintf("cn=%v,ou=%v,ou=%v,ou=users,%v", para.UserName, para.UserSecondOu, para.UserFirstOu, cfg.Ldap.SearchDomain) // 用户的DN
	}

	// 用户DN和新密码
	newPassword := para.UserNewPassword

	// 搜索用户并检索其散列密码
	searchRequest := ldap.NewSearchRequest(
		userDN,
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectClass=*)",
		[]string{"userPassword"}, // 根据您的LDAP架构更改此属性
		nil,
	)

	searchResult, err := l.Search(searchRequest)
	if err != nil {
		logrus.Error(err)
	}

	if len(searchResult.Entries) != 1 {
		logrus.Error("用户未找到或找到多个条目")
	}

	if err := l.Bind(userDN, para.UserOldPassword); err == nil {
		logrus.Info(fmt.Sprintf("User:%v Auth Required password  successfully!", para.UserName))
	} else {
		logrus.Error(fmt.Sprintf("Failed Auth Required User:%v. Reason: %v", para.UserName, err))
	}

	if para.UserNewPassword != para.UserNew2Password {
		result.Code = 50001
		result.Data = fmt.Sprintf("Failed Update to user:%v. Reason: The two passwords are different", para.UserName)
		logrus.Error(fmt.Sprintf("Failed Update to user:%v. Reason: The two passwords are different", para.UserName))
		return result
	}

	// 执行的修改操作
	mod := ldap.NewModifyRequest(userDN, nil)
	mod.Replace("userPassword", []string{newPassword})
	err = l.Modify(mod)

	if err != nil {
		result.Code = 50001
		result.Data = fmt.Sprintf("Failed Update to user password:%v. Reason: %v", para.UserName, err)
		logrus.Info(fmt.Sprintf("Failed Update to user password:%v. Reason: %v", para.UserName, err))
		return result
	}

	result.Code = 20000
	result.Data = fmt.Sprintf("Update to user password: %v successfully!", para.UserName)
	logrus.Info(fmt.Sprintf("Update to user password: %v successfully!\n", para.UserName))

	return result
}
