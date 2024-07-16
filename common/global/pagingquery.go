// @Author sunwenbo
// 2024/7/16 12:05
package global

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func PagingQuery(c *gin.Context) (pageNum, limit int) {
	// 从请求中获取分页参数
	page := c.Query("page")
	pageSize := c.Query("pageSize")

	// 默认分页参数，你也可以根据需求自定义默认值
	defaultPage := 1
	defaultPageSize := 999

	// 将字符串参数转换为整数，处理可能的错误
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		pageNum = defaultPage
	}
	limit, err = strconv.Atoi(pageSize)
	if err != nil || limit < 1 {
		limit = defaultPageSize
	}
	return
}
