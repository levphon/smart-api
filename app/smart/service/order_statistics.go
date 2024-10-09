// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"math"
	"smart-api/app/smart/models"
	"smart-api/app/smart/service/dto"
	"sort"
	"time"
)

// OrderStatistics 服务结构体
type OrderStatistics struct {
	service.Service
}

// GetStatistics 获取工单统计数据
func (e *OrderStatistics) GetStatistics() (map[string]interface{}, error) {
	var totalOrders, completedOrders, pendingOrders, currentProcessingOrders int64
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
	dailyAverage := 0.0
	if daysPassed > 0 {
		dailyAverage = float64(monthlyOrders) / float64(daysPassed)
	}

	// 保留两位小数
	dailyAverage = math.Round(dailyAverage*100) / 100

	// 计算完成率
	completionRate := 0
	if totalOrders > 0 {
		completionRate = int(float64(completedOrders) / float64(totalOrders) * 100)
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

// calculateComparisons 计算总工单数和当前处理工单数的周同比和日同比
func (e *OrderStatistics) calculateComparisons(totalOrders, currentProcessingOrders int64) (float64, float64, float64, float64, error) {
	var yesterdayTotal int64
	var lastWeekTotal int64
	var yesterdayCurrentProcessing int64
	var lastWeekCurrentProcessing int64

	// 查询前一天的总工单数
	err := e.Orm.Model(&models.OrderWorks{}).
		Where("created_at >= CURDATE() - INTERVAL 1 DAY").Count(&yesterdayTotal).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// 查询前一周的总工单数
	err = e.Orm.Model(&models.OrderWorks{}).
		Where("created_at >= CURDATE() - INTERVAL 7 DAY").Count(&lastWeekTotal).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// 查询前一天的当前处理工单数
	err = e.Orm.Model(&models.OrderWorks{}).
		Where("status = ? AND created_at >= CURDATE() - INTERVAL 1 DAY", "under-way").Count(&yesterdayCurrentProcessing).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// 查询前一周的当前处理工单数
	err = e.Orm.Model(&models.OrderWorks{}).
		Where("status = ? AND created_at >= CURDATE() - INTERVAL 7 DAY", "under-way").Count(&lastWeekCurrentProcessing).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// 计算周同比
	totalOrdersWeekOverWeek := calculatePercentage(totalOrders, lastWeekTotal)
	currentHandlerOrdersWeekOverWeek := calculatePercentage(currentProcessingOrders, lastWeekCurrentProcessing)

	// 计算日同比
	totalOrdersDayOverDay := calculatePercentage(totalOrders, yesterdayTotal)
	currentHandlerOrdersDayOverDay := calculatePercentage(currentProcessingOrders, yesterdayCurrentProcessing)

	return totalOrdersWeekOverWeek, totalOrdersDayOverDay, currentHandlerOrdersWeekOverWeek, currentHandlerOrdersDayOverDay, nil
}

// calculatePercentage 计算百分比，避免除以零
func calculatePercentage(numerator, denominator int64) float64 {
	if denominator == 0 {
		return 0.0 // 避免除以零
	}
	return float64(numerator) / float64(denominator) * 100
}

// GetOrderCountByPeriod 按周或按月统计工单数量
func (e *OrderStatistics) GetOrderCountByPeriod(period string) (dto.OrderCountByPeriodResponse, error) {
	var orders []struct {
		Date      string
		Total     int
		Completed int
		UnderWay  int
	}
	var personalRankings []struct {
		Creator string
		Count   int
	}

	response := dto.OrderCountByPeriodResponse{}

	var err error
	if period == "week" {
		err = e.Orm.Model(&models.OrderWorks{}).
			Select(`DATE_FORMAT(created_at, '%Y-%u') as date,
		         COUNT(*) as total,
		         SUM(CASE WHEN status = 'completed' THEN 1 ELSE 0 END) as completed,
		         SUM(CASE WHEN status = 'under-way' THEN 1 ELSE 0 END) as under_way`).
			Group("DATE_FORMAT(created_at, '%Y-%u')").
			Order("date ASC").
			Find(&orders).Error
	} else if period == "month" {
		err = e.Orm.Model(&models.OrderWorks{}).
			Select(`DATE_FORMAT(created_at, '%Y-%m') as date,
		         COUNT(*) as total,
		         SUM(CASE WHEN status = 'completed' THEN 1 ELSE 0 END) as completed,
		         SUM(CASE WHEN status = 'under-way' THEN 1 ELSE 0 END) as under_way`).
			Group("DATE_FORMAT(created_at, '%Y-%m')").
			Order("date ASC").
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
			Date:      order.Date,
			Total:     order.Total,
			Completed: order.Completed,
			UnderWay:  order.UnderWay,
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
		Total   int
		Average float64 // 确保这个字段被定义
	}

	response := dto.OrderRatingsResponse{}

	// 获取评分统计数据
	var err error
	if period == "week" {
		err = e.Orm.Model(&models.OrderRating{}).
			Select("DATE_FORMAT(order_works.created_at, '%Y-%u') as date, AVG(order_rating.rating) as average").
			Joins("JOIN order_works ON order_works.id = order_rating.order_id").
			Group("DATE_FORMAT(order_works.created_at, '%Y-%u')").
			Order("date ASC").
			Find(&ratings).Error
	} else if period == "month" {
		err = e.Orm.Model(&models.OrderRating{}).
			Select("DATE_FORMAT(order_works.created_at, '%Y-%m') as date, AVG(order_rating.rating) as average").
			Joins("JOIN order_works ON order_works.id = order_rating.order_id").
			Group("DATE_FORMAT(order_works.created_at, '%Y-%m')").
			Order("date ASC").
			Find(&ratings).Error
	} else {
		return response, errors.New("无效的时间周期")
	}

	if err != nil {
		return response, err
	}

	// 处理评分统计数据并赋值
	response.RatingsStats = make([]dto.RatingStat, len(ratings))
	for i, rating := range ratings {
		response.RatingsStats[i] = dto.RatingStat{
			Date:    rating.Date,
			Average: rating.Average,
		}
	}

	const k = 5.0 // 调节参数

	// 获取处理人的评分统计，包括工单数量和平均评分
	err = e.Orm.Model(&models.OrderRating{}).
		Select("sys_user.nick_name as creator, COUNT(order_rating.order_id) as total, AVG(order_rating.rating) as average").
		Joins("JOIN sys_user ON sys_user.user_id = order_rating.task_handler"). // 使用直接关联的 task_handler 字段
		Group("sys_user.user_id").
		Order("total ASC, average ASC"). // 优先按工单数量排序，再按平均评分排序
		Find(&personalRankings).Error
	if err != nil {
		return response, err
	}

	// 计算得分并赋值给响应
	response.RatingRanking = make([]dto.RatingRanking, len(personalRankings))
	for i, ranking := range personalRankings {
		// 使用加权得分公式计算最终得分
		score := (ranking.Average * float64(ranking.Total)) / (float64(ranking.Total) + k)

		// 将分数乘以 100 并保留两位小数
		score = math.Round(score * 100)

		// 如果需要将分数保留两位小数作为字符串展示，可以使用 strconv.FormatFloat
		// scoreStr := strconv.FormatFloat(score, 'f', 2, 64)  // 可选

		response.RatingRanking[i] = dto.RatingRanking{
			Name:  ranking.Creator,
			Score: score, // 使用加权得分
		}
	}

	// 排序排行榜（如果数据库已经按工单数量和平均分排序，可以省略）
	sort.Slice(response.RatingRanking, func(i, j int) bool {
		return response.RatingRanking[i].Score > response.RatingRanking[j].Score
	})

	return response, nil
}
