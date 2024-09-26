// @Author sunwenbo
// 2024/9/25 20:28
package dto

type OrderStatisticsResponse struct {
	TotalOrders                int     `json:"totalOrders"`                // 总工单数
	TotalOrdersWeekOverWeek    float64 `json:"totalOrdersWeekOverWeek"`    // 总工单数周同比
	TotalOrdersDayOverDay      float64 `json:"totalOrdersDayOverDay"`      // 总工单数日同比
	CompletedOrders            int     `json:"completedOrders"`            // 已完结工单
	PendingOrders              int     `json:"pendingOrders"`              // 待办工单
	CurrentHandlerOrders       int     `json:"currentHandlerOrders"`       // 当前处理的工单
	CurrentHandlerWeekOverWeek float64 `json:"currentHandlerWeekOverWeek"` // 当前处理工单周同比
	CurrentHandlerDayOverDay   float64 `json:"currentHandlerDayOverDay"`   // 当前处理工单日同比
	DailyAverage               int64   `json:"dailyAverage"`               // 日均工单数
}

type OrderCountByPeriodResponse struct {
	OrderStats        []OrderCountStat    `json:"orderStats"`        // 工单统计
	SubmissionRanking []SubmissionRanking `json:"submissionRanking"` // 个人提交排行榜
}

type OrderCountStat struct {
	Date      string `json:"date"`      // 日期
	Total     int    `json:"total"`     // 所有
	Completed int    `json:"completed"` // 已完成的
	UnderWay  int    `json:"underWay"`  // 进行中的
}

type SubmissionRanking struct {
	Name  string `json:"name"`  // 创建人
	Total int    `json:"total"` // 提交数量
}

type OrderRatingsResponse struct {
	RatingsStats  []RatingStat    `json:"ratingsStats"`  // 评分统计
	RatingRanking []RatingRanking `json:"ratingRanking"` // 个人评分排行榜
}

type RatingStat struct {
	Date    string  `json:"date"`    // 日期
	Average float64 `json:"average"` // 平均评分
}

type RatingRanking struct {
	Name  string  `json:"name"`  // 评价人
	Score float64 `json:"score"` // 平均评分
}
