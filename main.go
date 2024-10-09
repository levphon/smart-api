package main

import (
	"smart-api/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6 --instanceName smart -o ./docs/smart

// @title smart-api API
// @version 1.0.0
// @description 基于Gin + Vue + Element UI的前后端分离运维工单系统的接口文档
// @description 添加qq群: xxxx 进入技术交流群 请先star，谢谢！
// @license.name MIT
// @contact.email swb956721830@163.com

// @license.name Apache 2.0
// @license.url https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
	cmd.Execute()
}
