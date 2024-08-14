package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
	"regexp"
	"smart-ops/common"
)

func LdapAddUser(l *ldap.Conn, cfg *common.Config, para *common.User) (result common.PostResponse) {
	//defer l.Close()
	var userDN string
	regex := regexp.MustCompile("^[a-zA-Z0-9-]+$")
	if !regex.MatchString(para.UserName) {
		result.Code = 50001
		result.Data = fmt.Sprintf("Current user name %s,The user name does not meet the requirements. It must contain letters or numbers or '-'.", para.UserName)
		logrus.Error(fmt.Sprintf("Current user name %s,The user name does not meet the requirements. It must contain letters or numbers or '-'.", para.UserName))
		return result
	}

	if para.UserSecondOu == "sre" || para.UserSecondOu == "qa" || para.UserSecondOu == "ipo" {
		// 用户属性信息
		userDN = fmt.Sprintf("cn=%v,ou=%v,ou=users,%v", para.UserName, para.UserSecondOu, cfg.Ldap.SearchDomain) // 用户的DN
	} else {
		// 用户属性信息
		userDN = fmt.Sprintf("cn=%v,ou=%v,ou=%v,ou=users,%v", para.UserName, para.UserSecondOu, para.UserFirstOu, cfg.Ldap.SearchDomain) // 用户的DN
	}

	if para.UserTwoPasswd != para.UserOnePasswd {
		result.Code = 50001
		result.Data = fmt.Sprintf("Failed Sign In user:%v. Reason: The two passwords are different", para.UserName)
		logrus.Error(fmt.Sprintf("Failed Sign In user:%v. Reason: The two passwords are different", para.UserName))
		return result
	}

	userAttrs := []ldap.Attribute{
		ldap.Attribute{"objectClass", []string{"inetOrgPerson", "top"}},
		ldap.Attribute{"cn", []string{para.UserName}},
		ldap.Attribute{"sn", []string{para.UserName}},
		ldap.Attribute{"uid", []string{para.UserName}},
		ldap.Attribute{"userPassword", []string{para.UserTwoPasswd}},
	}

	// 创建用户条目
	err := l.Add(&ldap.AddRequest{
		DN:         userDN,
		Attributes: userAttrs,
	})

	if err != nil {
		result.Code = 50001
		result.Data = fmt.Sprintf("Failed to create user:%v. reason: %v", para.UserName, err)
		logrus.Error(fmt.Sprintf("Failed to create user:%v. reason: %v", para.UserName, err))
		return result
	}
	result.Code = 20000
	result.Data = fmt.Sprintf("created user: %v successfully!", para.UserName)
	logrus.Info(fmt.Sprintf("created user: %v successfully!\n", para.UserName))

	return result
}
