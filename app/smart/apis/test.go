package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/smart/service"
	"net/http"
)

type Test struct {
	api.Api
}

func (e Test) Get(c *gin.Context) {
	s := service.Test{}
	message := s.HelloWorld()

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})

}
