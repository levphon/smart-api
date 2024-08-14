package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
	"smart-ops/common"
)

func LdapAddUserToGroup(l *ldap.Conn, cfg *common.Config, para *common.UserToGroup) (result common.PostResponse) {
	// 管理员用户密码认证
	if para.AdminUser != "admin" || para.AdminPwd != cfg.Ldap.BindPassword {
		result.Code = 50001
		result.Data = fmt.Sprintf("LDAP permission management requires the authentication administrator user and password. 当前管理员用户为: %v", para.AdminUser)
		logrus.Error(fmt.Sprintf("LDAP permission management requires the authentication administrator user and password.当前管理员用户为: %v", para.AdminUser))
		return result
	}

	defer l.Close()

	// 组DN
	groupDN := fmt.Sprintf("cn=%v,ou=%v,ou=Application,%v", para.GroupCn, para.GroupOu, cfg.Ldap.SearchDomain)
	groupMemberAttr := "memberUid" // 组成员属性名
	// 获取现有组的属性
	searchRequest := ldap.NewSearchRequest(
		groupDN,
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectClass=*)",
		[]string{groupMemberAttr},
		nil,
	)

	searchResult, err := l.Search(searchRequest)
	if err != nil {
		logrus.Error("Failed to search for group: %v", err)
	}
	if len(searchResult.Entries) != 1 {
		logrus.Error("Group not found or multiple groups found")
	}

	// 添加用户到组的成员列表
	modifyRequest := ldap.NewModifyRequest(groupDN, nil)
	modifyRequest.Add(groupMemberAttr, []string{para.UserName})
	err = l.Modify(modifyRequest)

	if err != nil {
		result.Code = 50001
		result.Data = fmt.Sprintf("Failed to add user:%v to group:%v  %v", para.UserName, para.GroupCn, err)
		logrus.Error(fmt.Sprintf("Failed to add user:%v to group:%v  %v", para.UserName, para.GroupCn, err))
		return result
	}
	result.Code = 20000
	result.Data = fmt.Sprintf("User add user:%v to the %v successfully!", para.UserName, para.GroupCn)
	logrus.Info(fmt.Sprintf("User add user:%v to the %v successfully!", para.UserName, para.GroupCn))

	return result
}
