// @Author sunwenbo
// 2024/7/12 20:28
package api

import (
	"fmt"
	"smart-api/app/smart/router"
)

func init() {
	fmt.Println("注册Smart api路由...")
	//注册Smart路由 fixme 其他应用的路由，自定义注册路由
	// AppRouters 在system.go 中定义，将自定义的路由追加到AppRouters切片实现路由注册
	AppRouters = append(AppRouters, router.InitRouter)
}
