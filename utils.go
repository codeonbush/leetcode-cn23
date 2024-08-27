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
func CompareTimeRangesMoM(granularity string, originStart, now time.Time) (currentRange TimeRange, previousRange TimeRange) {

	switch granularity {
	case ByMinute:
		currentRange = TimeRange{
			Start: now.Truncate(time.Minute),
			End:   now,
		}
		previousRange = TimeRange{
			Start: now.Add(-time.Minute).Truncate(time.Minute),
			End:   now.Truncate(time.Minute).Add(-time.Second),
		}

	case ByHour:
		currentStart := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		previousStart := currentStart.Add(-time.Hour)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(now.Sub(currentStart)),
		}

	case ByDay:
		currentStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		previousStart := currentStart.Add(-24 * time.Hour)

		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(currentRange.End.Sub(currentRange.Start)),
		}

	case ByWeek:
		// 计算本周的开始时间，从周一 00:00 开始
		offset := int(time.Monday - now.Weekday())
		if offset > 0 {
			offset = -6 // 如果今天是周日，则回退到上周一
		}
		currentStart := time.Date(now.Year(), now.Month(), now.Day()+offset, 0, 0, 0, 0, now.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		previousStart := currentStart.AddDate(0, 0, -7)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(now.Sub(currentStart)),
		}
	case ByMonth:
		// 本月开始时间
		currentStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		// 上月开始时间
		previousStart := currentStart.AddDate(0, -1, 0)

		// 计算上月的结束时间
		previousEnd := previousStart.AddDate(0, 1, 0).Add(-time.Nanosecond) // 上月最后一刻

		// 当前时间范围
		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}

		// 上月对应的时间范围
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(now.Sub(currentStart)),
		}

		// 确保上月的结束时间不超过上月的实际结束时间
		if previousRange.End.After(previousEnd) {
			previousRange.End = previousEnd
		}
	case ByQuarter:
		// 计算本季度的开始月份
		month := (now.Month()-1)/3*3 + 1
		currentStart := time.Date(now.Year(), month, 1, 0, 0, 0, 0, now.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		previousStart := currentStart.AddDate(0, -3, 0)

		// 当前季度的时间范围
		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}

		// 计算上季度的实际结束时间
		previousEnd := previousStart.AddDate(0, 3, 0).Add(-time.Nanosecond) // 上季度的实际结束时间为上季度的最后一刻

		// 上季度的时间范围
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(now.Sub(currentStart)),
		}

		// 确保上季度的结束时间不超过上季度的实际结束时间
		if previousRange.End.After(now) {
			previousRange.End = previousEnd
		}

	default:
	}

	return
}

// CompareTimeRangesYoY 获取同比时间范围
func CompareTimeRangesYoY(granularity string, originStart, now time.Time) (currentRange TimeRange, yearOverYearRange TimeRange) {

	switch granularity {
	case ByMinute:
		currentRange = TimeRange{
			Start: now.Truncate(time.Minute),
			End:   now,
		}
		yearOverYearStart := now.AddDate(0, 0, -1).Truncate(time.Minute)
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(time.Minute).Add(-time.Second),
		}

	case ByHour:
		currentStart := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		yearOverYearStart := currentStart.AddDate(0, 0, -1)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(now.Sub(currentStart)),
		}

	case ByDay:
		currentStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		yearOverYearStart := currentStart.AddDate(0, 0, -7)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(now.Sub(currentStart)),
		}

	case ByWeek:
		startOfWeek := now.Truncate(24 * time.Hour).Add(-(time.Duration(now.Weekday()-time.Monday) * 24 * time.Hour))
		if originStart.After(startOfWeek) {
			startOfWeek = originStart
		}
		// 当前分钟
		currentMinute := now

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
		currentStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		yearOverYearStart := currentStart.AddDate(-1, 0, 0)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(now.Sub(currentStart)),
		}

	case ByQuarter:
		// 计算当前季度的开始月份
		month := (now.Month()-1)/3*3 + 1
		currentStart := time.Date(now.Year(), month, 1, 0, 0, 0, 0, now.Location())
		if originStart.After(currentStart) {
			currentStart = originStart
		}
		// 计算去年同季度的开始时间
		yearOverYearStart := currentStart.AddDate(-1, 0, 0)

		// 当前季度的时间范围
		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}

		// 去年同季度的时间范围
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(now.Sub(currentStart)),
		}

	default:
	}

	return
}
