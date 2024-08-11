package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
	"log"
	"smart-ops/common"
	"sort"
	"strconv"
)

func LdapAddGroup(l *ldap.Conn, cfg *common.Config, para *common.AddGroup) (result common.LdapAuthResponse) {

	defer l.Close()

	// 设置要创建的组的DN
	groupDN := fmt.Sprintf("cn=%v,ou=%v,ou=Application,%v", para.GroupCn, para.GroupOu, cfg.Ldap.SearchDomain) // 组的DN

	baseDN := "ou=Application," + cfg.Ldap.SearchDomain
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectClass=posixGroup)",
		[]string{"gidNumber"},
		nil,
	)
	searchResult, err := l.Search(searchRequest)
	if err != nil {
		log.Fatalf("LDAP搜索失败: %v", err)
	}

	// 遍历搜索结果并提取组ID
	gidIntArr := []int{}
	for _, entry := range searchResult.Entries {
		groupId := entry.GetAttributeValue("gidNumber")
		gid, _ := strconv.Atoi(groupId)
		gidIntArr = append(gidIntArr, gid)
	}
	sort.Ints(gidIntArr)
	newGid := gidIntArr[len(gidIntArr)-1] + 1
	addRequest := ldap.NewAddRequest(groupDN, nil)
	addRequest.Attribute("objectClass", []string{"posixGroup", "top"})
	addRequest.Attribute("cn", []string{para.GroupCn})
	addRequest.Attribute("gidNumber", []string{strconv.Itoa(newGid)})

	err = l.Add(addRequest)

	if err != nil {
		result.Code = 50001
		result.Data = fmt.Sprintf("Failed to create group:%v. reason: %v", para.GroupCn, err)
		logrus.Error(fmt.Sprintf("Failed to create group:%v. reason: %v", para.GroupCn, err))
		return result
	}
	result.Code = 20000
	result.Data = fmt.Sprintf("created group: %v successfully!", para.GroupCn)
	logrus.Info(fmt.Sprintf("created group: %v successfully!\n", para.GroupCn))

	return result
}
