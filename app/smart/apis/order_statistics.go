// @Author sunwenbo
// 2024/9/25 20:28
package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"smart-api/app/smart/service"
	"smart-api/app/smart/service/dto"
)

type OrderStatistics struct {
	api.Api
}

// GetOrderStatistics 获取工单统计数据
// @Summary 获取工单统计数据
// @Description 获取工单统计数据
// @Tags 工单统计
// @Success 200 {object} response.Response{data=dto.OrderStatisticsResponse} "{"code": 200, "data": {...}}"
// @Router /api/v1/statistics [get]
// @Security Bearer
func (e OrderStatistics) GetOrderStatistics(c *gin.Context) {
	s := service.OrderStatistics{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	data, err := s.GetStatistics()
	if err != nil {
		e.Error(500, err, "查询统计数据失败")
		return
	}

	e.OK(data, "查询成功")
}

// GetOrderCountByPeriod 获取工单数量统计数据
// @Summary 获取工单数量统计数据
// @Description 根据周或月统计工单数量
// @Tags 工单统计
// @Param period query string true "period" enums(week, month)
// @Success 200 {object} response.Response{data=dto.OrderCountByPeriodResponse} "{"code": 200, "data": {...}}"
// @Router /api/v1/statistics/orders/count [get]
// @Security Bearer
func (e OrderStatistics) GetOrderCountByPeriod(c *gin.Context) {
	s := service.OrderStatistics{}
	period := c.Query("period")
	req := dto.OrderCountByPeriodResponse{}

	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	data, err := s.GetOrderCountByPeriod(period)
	if err != nil {
		e.Error(500, err, "查询统计数据失败")
		return
	}

	e.OK(data, "查询成功")
}

// GetOrderRatingsByPeriod godoc
// @Summary      获取评分统计
// @Description  根据周或月统计评分数据
// @Tags         OrderStatistics
// @Accept       json
// @Produce      json
// @Param        period query string true "统计周期，支持 week 或 month"
// @Success      200 {object} dto.OrderRatingsResponse "成功返回评分统计数据"
// @Failure      500 {object} string "查询评分统计数据失败"
// @Router       /ratings [get]
func (e OrderStatistics) GetOrderRatingsByPeriod(c *gin.Context) {
	s := service.OrderStatistics{}
	period := c.Query("period")
	req := dto.OrderRatingsResponse{}

	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	data, err := s.GetOrderRatings(period)
	if err != nil {
		e.Error(500, err, "查询评分统计数据失败")
		return
	}

	e.OK(data, "查询成功")
}
