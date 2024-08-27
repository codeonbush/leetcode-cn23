package cn

import "time"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const (
	ByMinute   = "BY_MINUTE"
	ByHour     = "BY_HOUR"
	ByDay      = "BY_DAY"
	ByWeek     = "BY_WEEK"
	ByMonth    = "BY_MONTH"
	ByQuarter  = "BY_QUARTER"
	ByCutRange = "BY_CUT_RANGE"
	AmountTo   = "AMOUNT_TO"
)

// TimeRange 表示一个时间范围
type TimeRange struct {
	Start time.Time
	End   time.Time
}

// CompareTimeRangesMoM 获取环比时间范围
func CompareTimeRangesMoM(granularity string, originStart, originEnd time.Time) (currentRange TimeRange, previousRange TimeRange) {

	switch granularity {
	case ByMinute:
		currentRange = TimeRange{
			Start: originEnd.Truncate(time.Minute),
			End:   originEnd,
		}
		previousRange = TimeRange{
			Start: originEnd.Add(-time.Minute).Truncate(time.Minute),
			End:   originEnd.Truncate(time.Minute).Add(-time.Second),
		}

	case ByHour:
		currentStart := time.Date(originEnd.Year(), originEnd.Month(), originEnd.Day(), originEnd.Hour(), 0, 0, 0, originEnd.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		previousStart := currentStart.Add(-time.Hour)

		currentRange = TimeRange{
			Start: currentStart,
			End:   originEnd,
		}
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(originEnd.Sub(currentStart)),
		}

	case ByDay:
		currentStart := time.Date(originEnd.Year(), originEnd.Month(), originEnd.Day(), 0, 0, 0, 0, originEnd.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		currentRange = TimeRange{
			Start: currentStart,
			End:   originEnd,
		}
		previousStart := currentStart.Add(-24 * time.Hour)

		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(currentRange.End.Sub(currentRange.Start)),
		}

	case ByWeek:
		// 计算本周的开始时间，从周一 00:00 开始
		offset := int(time.Monday - originEnd.Weekday())
		if offset > 0 {
			offset = -6 // 如果今天是周日，则回退到上周一
		}
		currentStart := time.Date(originEnd.Year(), originEnd.Month(), originEnd.Day()+offset, 0, 0, 0, 0, originEnd.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		previousStart := currentStart.AddDate(0, 0, -7)

		currentRange = TimeRange{
			Start: currentStart,
			End:   originEnd,
		}
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(originEnd.Sub(currentStart)),
		}
	case ByMonth:
		// 本月开始时间
		currentStart := time.Date(originEnd.Year(), originEnd.Month(), 1, 0, 0, 0, 0, originEnd.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		// 上月开始时间
		previousStart := currentStart.AddDate(0, -1, 0)

		// 检查上个月的日期是否有效
		if previousStart.Day() != currentStart.Day() {
			// 如果无效，返回上个月的最后一天
			previousStart = time.Date(previousStart.Year(), previousStart.Month(), 0, previousStart.Hour(), previousStart.Minute(), previousStart.Second(), 0, previousStart.Location())
		}

		// 计算上月的结束时间
		previousEnd := previousStart.AddDate(0, 1, 0).Add(-time.Nanosecond) // 上月最后一刻

		// 当前时间范围
		currentRange = TimeRange{
			Start: currentStart,
			End:   originEnd,
		}

		// 上月对应的时间范围
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(originEnd.Sub(currentStart)),
		}

		// 确保上月的结束时间不超过上月的实际结束时间
		if previousRange.End.After(previousEnd) {
			previousRange.End = previousEnd
		}
	case ByQuarter:
		// 计算本季度的开始月份
		// 本季度开始时间
		currentStart := time.Date(originEnd.Year(), getQuarterStartMonth(originEnd.Month()), 1, 0, 0, 0, 0, originEnd.Location())
		currentQuarterStart := time.Date(originEnd.Year(), getQuarterStartMonth(originEnd.Month()), 1, 0, 0, 0, 0, originEnd.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}

		previousStart := currentStart.AddDate(0, -3, 0)

		// 检查上季度的日期是否有效
		if previousStart.Month() != getQuarterStartMonth(currentQuarterStart.Month()) {
			// 如果无效，返回上季度的最后一天
			previousStart = time.Date(previousStart.Year(), previousStart.Month()+3, 0, previousStart.Hour(), previousStart.Minute(), previousStart.Second(), 0, previousStart.Location())
		}

		// 当前季度的时间范围
		currentRange = TimeRange{
			Start: currentStart,
			End:   originEnd,
		}

		// 计算上季度的实际结束时间
		previousEnd := previousStart.AddDate(0, 3, 0).Add(-time.Nanosecond) // 上季度的实际结束时间为上季度的最后一刻

		// 上季度的时间范围
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(originEnd.Sub(currentStart)),
		}

		// 确保上季度的结束时间不超过上季度的实际结束时间
		if previousRange.End.After(originEnd) {
			previousRange.End = previousEnd
		}

	default:
	}

	return
}

// 获取季度开始的月份
func getQuarterStartMonth(month time.Month) time.Month {
	switch month {
	case time.January, time.February, time.March:
		return time.January
	case time.April, time.May, time.June:
		return time.April
	case time.July, time.August, time.September:
		return time.July
	case time.October, time.November, time.December:
		return time.October
	}
	return time.January // 默认返回一月
}

// CompareTimeRangesYoY 获取同比时间范围
func CompareTimeRangesYoY(granularity string, originStart, originEnd time.Time) (currentRange TimeRange, yearOverYearRange TimeRange) {

	switch granularity {
	case ByMinute:
		currentRange = TimeRange{
			Start: originEnd.Truncate(time.Minute),
			End:   originEnd,
		}
		yearOverYearStart := originEnd.AddDate(0, 0, -1).Truncate(time.Minute)
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(time.Minute).Add(-time.Second),
		}

	case ByHour:
		currentStart := time.Date(originEnd.Year(), originEnd.Month(), originEnd.Day(), originEnd.Hour(), 0, 0, 0, originEnd.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		yearOverYearStart := currentStart.AddDate(0, 0, -1)

		currentRange = TimeRange{
			Start: currentStart,
			End:   originEnd,
		}
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(originEnd.Sub(currentStart)),
		}

	case ByDay:
		currentStart := time.Date(originEnd.Year(), originEnd.Month(), originEnd.Day(), 0, 0, 0, 0, originEnd.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		yearOverYearStart := currentStart.AddDate(0, 0, -7)

		currentRange = TimeRange{
			Start: currentStart,
			End:   originEnd,
		}
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(originEnd.Sub(currentStart)),
		}

	case ByWeek:
		startOfWeek := originEnd.Truncate(24 * time.Hour).Add(-(time.Duration(originEnd.Weekday()-time.Monday) * 24 * time.Hour))
		if originStart.After(startOfWeek) {
			startOfWeek = originStart
		}
		// 当前分钟
		currentMinute := originEnd

		// 计算四周前开始时间（从周一 00:00 开始至四周前的对应时间）
		fourWeeksAgoStart := startOfWeek.AddDate(0, 0, -28) // 向前计算 28 天

		// 设置当前时间区间
		currentRange = TimeRange{
			Start: startOfWeek,
			End:   currentMinute,
		}

		// 设置四周前的同比时间区间
		yearOverYearRange = TimeRange{
			Start: fourWeeksAgoStart,
			End:   fourWeeksAgoStart.Add(currentMinute.Sub(startOfWeek)),
		}

	case ByMonth:
		currentStart := time.Date(originEnd.Year(), originEnd.Month(), 1, 0, 0, 0, 0, originEnd.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		yearOverYearStart := currentStart.AddDate(-1, 0, 0)

		currentRange = TimeRange{
			Start: currentStart,
			End:   originEnd,
		}
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(originEnd.Sub(currentStart)),
		}

	case ByQuarter:
		// 计算当前季度的开始月份
		month := (originEnd.Month()-1)/3*3 + 1
		currentStart := time.Date(originEnd.Year(), month, 1, 0, 0, 0, 0, originEnd.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		// 计算去年同季度的开始时间
		yearOverYearStart := currentStart.AddDate(-1, 0, 0)

		// 当前季度的时间范围
		currentRange = TimeRange{
			Start: currentStart,
			End:   originEnd,
		}

		// 去年同季度的时间范围
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(originEnd.Sub(currentStart)),
		}

	default:
	}

	return
}
