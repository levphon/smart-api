package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type Test struct {
	service.Service
}

// HelloWorld 方法，返回 "Hello World"
func (s Test) HelloWorld() string {
	return "Hello World"
}
