package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"net/http"
)

func DemoEvn() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if config.ApplicationConfig.Mode == "demo" {
			if method == "GET" ||
				method == "OPTIONS" ||
				c.Request.RequestURI == "/api/v1/login" ||
				c.Request.RequestURI == "/api/v1/logout" {
				c.Next()
			} else {
				// 添加跨域头部
				c.Header("Access-Control-Allow-Origin", "*")
				c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
				c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept, user-agent")
				c.Header("Access-Control-Allow-Credentials", "true")
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "谢谢您的参与，但为了大家更好的体验，所以本次提交就算了吧！\U0001F600\U0001F600\U0001F600",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
