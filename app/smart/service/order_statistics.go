// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/smart/models"
	"go-admin/app/smart/service/dto"
	"time"
)

// OrderStatistics 服务结构体
type OrderStatistics struct {
	service.Service
}

// GetStatistics 获取工单统计数据
func (e *OrderStatistics) GetStatistics() (map[string]interface{}, error) {
	var totalOrders int64
	var completedOrders int64
	var pendingOrders int64
	var currentProcessingOrders int64
	var monthlyOrders int64

	// 查询总工单数
	err := e.Orm.Model(&models.OrderWorks{}).Count(&totalOrders).Error
	if err != nil {
		return nil, err
	}

	// 查询已完结的工单
	err = e.Orm.Model(&models.OrderWorks{}).Where("status IN (?, ?)", "termination", "manual-termination").Count(&completedOrders).Error
	if err != nil {
		return nil, err
	}

	// 查询待办的工单
	statuses := []string{"under-way", "reopen", "reject"}
	err = e.Orm.Model(&models.OrderWorks{}).Where("status IN ?", statuses).Count(&pendingOrders).Error
	if err != nil {
		return nil, err
	}

	// 查询当前处理的工单
	err = e.Orm.Model(&models.OrderWorks{}).Where("status = ?", "under-way").Count(&currentProcessingOrders).Error
	if err != nil {
		return nil, err
	}

	// 获取当前月
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// 查询当月工单数
	err = e.Orm.Model(&models.OrderWorks{}).Where("created_at >= ?", startOfMonth).Count(&monthlyOrders).Error
	if err != nil {
		return nil, err
	}

	// 计算已过天数
	daysPassed := now.Day()

	// 计算日均工单数
	var dailyAverage float64
	if daysPassed > 0 {
		dailyAverage = float64(monthlyOrders) / float64(daysPassed)
	} else {
		dailyAverage = 0
	}

	// 计算完成率
	completionRate := 0.0
	if totalOrders > 0 {
		completionRate = float64(completedOrders) / float64(totalOrders) * 100
	}

	// 计算总工单数和当前处理工单数的周同比、日同比
	totalOrdersWeekOverWeek, totalOrdersDayOverDay, currentHandlerOrdersWeekOverWeek, currentHandlerOrdersDayOverDay, err := e.calculateComparisons(totalOrders, currentProcessingOrders)
	if err != nil {
		return nil, err
	}

	// 统计数据整合
	data := map[string]interface{}{
		"totalOrders":                      totalOrders,
		"completedOrders":                  completedOrders,
		"completionRate":                   completionRate,
		"pendingOrders":                    pendingOrders,
		"currentProcessingOrders":          currentProcessingOrders,
		"totalOrdersWeekOverWeek":          totalOrdersWeekOverWeek,
		"totalOrdersDayOverDay":            totalOrdersDayOverDay,
		"currentHandlerOrdersWeekOverWeek": currentHandlerOrdersWeekOverWeek,
		"currentHandlerOrdersDayOverDay":   currentHandlerOrdersDayOverDay,
		"dailyAverage":                     dailyAverage, // 添加日均工单数
	}
	return data, nil
}

// calculateComparisons 计算总工单数和当前处理工单数的周同比、日同比
func (e *OrderStatistics) calculateComparisons(totalOrders, currentProcessingOrders int64) (float64, float64, float64, float64, error) {
	var yesterdayTotal int64
	var lastWeekTotal int64
	var yesterdayCurrentProcessing int64
	var lastWeekCurrentProcessing int64

	// 查询前一天和前一周的总工单数
	err := e.Orm.Model(&models.OrderWorks{}).
		Where("created_at >= CURDATE() - INTERVAL 1 DAY").Count(&yesterdayTotal).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	err = e.Orm.Model(&models.OrderWorks{}).
		Where("created_at >= CURDATE() - INTERVAL 7 DAY").Count(&lastWeekTotal).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// 查询前一天和前一周的当前处理工单数
	err = e.Orm.Model(&models.OrderWorks{}).
		Where("status = ? AND created_at >= CURDATE() - INTERVAL 1 DAY", "under-way").Count(&yesterdayCurrentProcessing).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	err = e.Orm.Model(&models.OrderWorks{}).
		Where("status = ? AND created_at >= CURDATE() - INTERVAL 7 DAY", "under-way").Count(&lastWeekCurrentProcessing).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// 计算周同比和日同比
	totalOrdersWeekOverWeek := float64(totalOrders) / float64(lastWeekTotal) * 100
	totalOrdersDayOverDay := float64(totalOrders) / float64(yesterdayTotal) * 100

	currentHandlerOrdersWeekOverWeek := float64(currentProcessingOrders) / float64(lastWeekCurrentProcessing) * 100
	currentHandlerOrdersDayOverDay := float64(currentProcessingOrders) / float64(yesterdayCurrentProcessing) * 100

	return totalOrdersWeekOverWeek, totalOrdersDayOverDay, currentHandlerOrdersWeekOverWeek, currentHandlerOrdersDayOverDay, nil
}

// GetOrderCountByPeriod 按周或按月统计工单数量
func (e *OrderStatistics) GetOrderCountByPeriod(period string) (dto.OrderCountByPeriodResponse, error) {
	var orders []struct {
		Date  string
		Count int
	}
	var personalRankings []struct {
		Creator string
		Count   int
	}

	response := dto.OrderCountByPeriodResponse{}

	var err error
	if period == "week" {
		err = e.Orm.Model(&models.OrderWorks{}).
			Select("DATE_FORMAT(created_at, '%Y-%u') as date, COUNT(*) as count").
			Group("DATE_FORMAT(created_at, '%Y-%u')").
			Order("date ASC"). // 添加排序
			Find(&orders).Error
	} else if period == "month" {
		err = e.Orm.Model(&models.OrderWorks{}).
			Select("DATE_FORMAT(created_at, '%Y-%m') as date, COUNT(*) as count").
			Group("DATE_FORMAT(created_at, '%Y-%m')").
			Order("date ASC"). // 添加排序
			Find(&orders).Error
	} else {
		return response, errors.New("invalid period")
	}

	if err != nil {
		return response, err
	}

	// 统计个人提交工单数量
	err = e.Orm.Model(&models.OrderWorks{}).
		Select("u.nick_name as creator, COUNT(*) as count").
		Joins("JOIN sys_user u ON u.user_id = order_works.create_by").
		Group("u.nick_name").
		Order("count DESC").
		Find(&personalRankings).Error
	if err != nil {
		return response, err
	}

	// 将结果赋值给响应
	response.OrderStats = make([]dto.OrderCountStat, len(orders))
	for i, order := range orders {
		response.OrderStats[i] = dto.OrderCountStat{
			Date:  order.Date,  // 确保这里使用正确的字段
			Count: order.Count, // 确保这里使用正确的字段
		}
	}
	response.SubmissionRanking = make([]dto.SubmissionRanking, len(personalRankings))
	for i, ranking := range personalRankings {
		response.SubmissionRanking[i] = dto.SubmissionRanking{
			Name:  ranking.Creator,
			Total: ranking.Count,
		}
	}

	return response, nil
}

// GetOrderRatings 按周或按月统计评分数据
func (e *OrderStatistics) GetOrderRatings(period string) (dto.OrderRatingsResponse, error) {
	var ratings []struct {
		Date    string
		Average float64
	}
	var personalRankings []struct {
		Creator string
		Count   int
	}

	response := dto.OrderRatingsResponse{}

	var err error
	if period == "week" {
		err = e.Orm.Model(&models.OrderRating{}).
			Select("DATE_FORMAT(order_works.created_at, '%Y-%u') as date, AVG(order_rating.rating) as average").
			Joins("JOIN order_works ON order_works.id = order_rating.order_id").
			Group("DATE_FORMAT(order_works.created_at, '%Y-%u')").
			Find(&ratings).Error
	} else if period == "month" {
		err = e.Orm.Model(&models.OrderRating{}).
			Select("DATE_FORMAT(order_works.created_at, '%Y-%m') as date, AVG(order_rating.rating) as average").
			Joins("JOIN order_works ON order_works.id = order_rating.order_id").
			Group("DATE_FORMAT(order_works.created_at, '%Y-%m')").
			Find(&ratings).Error
	} else {
		return response, errors.New("invalid period")
	}

	if err != nil {
		return response, err
	}

	// 统计个人评分数量，获取用户姓名
	err = e.Orm.Model(&models.OrderRating{}).
		Select("sys_user.nick_name as creator, COUNT(*) as count").
		Joins("JOIN order_works ON order_works.id = order_rating.order_id").
		Joins("JOIN sys_user ON sys_user.user_id = order_works.create_by"). // 根据需要调整连接条件
		Group("sys_user.user_id").                                          // 依据用户ID分组
		Order("count DESC").
		Find(&personalRankings).Error
	if err != nil {
		return response, err
	}

	// 将结果赋值给响应
	response.RatingsStats = make([]dto.RatingStat, len(ratings))
	for i, rating := range ratings {
		response.RatingsStats[i] = dto.RatingStat{
			Date:    rating.Date,
			Average: rating.Average,
		}
	}

	response.RatingRanking = make([]dto.RatingRanking, len(personalRankings))
	for i, ranking := range personalRankings {
		response.RatingRanking[i] = dto.RatingRanking{
			Name:  ranking.Creator, // 现在是昵称
			Total: ranking.Count,
		}
	}

	return response, nil
}
