package main

import (
	"fmt"
	"time"
)

// TimeRange 表示一个时间范围
type TimeRange struct {
	Start time.Time
	End   time.Time
}

// CompareTimeRanges 返回当前时间范围与上一时间范围
func CompareTimeRanges(granularity string) (currentRange TimeRange, previousRange TimeRange) {
	now := time.Now()

	switch granularity {
	case "minute":
		currentRange = TimeRange{
			Start: now.Truncate(time.Minute),
			End:   now,
		}
		previousRange = TimeRange{
			Start: now.Add(-time.Minute).Truncate(time.Minute),
			End:   now.Truncate(time.Minute),
		}

	case "hour":
		currentStart := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
		previousStart := currentStart.Add(-time.Hour)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(time.Duration(now.Minute()) * time.Minute),
		}

	case "day":
		currentStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		previousStart := currentStart.Add(-24 * time.Hour)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(time.Duration(now.Hour())*time.Hour + time.Duration(now.Minute())*time.Minute),
		}

	case "week":
		// 计算本周的开始时间，从周一 00:00 开始
		offset := int(time.Monday - now.Weekday())
		if offset > 0 {
			offset = -6 // 如果今天是周日，则回退到上周一
		}
		currentStart := time.Date(now.Year(), now.Month(), now.Day()+offset, 0, 0, 0, 0, now.Location())
		previousStart := currentStart.AddDate(0, 0, -7)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(time.Duration(now.Sub(currentStart))),
		}

	case "month":
		currentStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		previousStart := currentStart.AddDate(0, -1, 0)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(time.Duration(now.Sub(currentStart))),
		}

	case "quarter":
		month := (now.Month()-1)/3*3 + 1
		currentStart := time.Date(now.Year(), month, 1, 0, 0, 0, 0, now.Location())
		previousStart := currentStart.AddDate(0, -3, 0)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		previousRange = TimeRange{
			Start: previousStart,
			End:   previousStart.Add(time.Duration(now.Sub(currentStart))),
		}

	default:
		fmt.Println("Unsupported granularity")
	}

	return
}

// 获取当前日期是当前月的第几周
func weekOfMonth(t time.Time) int {
	// 找到当前月的第一天
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	// 计算当前日期在这个月的第几周
	_, week := t.ISOWeek()
	_, firstWeek := firstOfMonth.ISOWeek()
	return week - firstWeek + 1
}

// 获取指定月的周数
func weeksInMonth(t time.Time) int {
	// 获取下个月的第一天
	firstOfNextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	// 获取这个月的最后一天
	lastDayOfMonth := firstOfNextMonth.Add(-time.Nanosecond)
	return weekOfMonth(lastDayOfMonth)
}

// CompareTimeRanges 返回当前时间范围与同比时间范围
func CompareTimeRangesT(granularity string) (currentRange TimeRange, yearOverYearRange TimeRange) {
	now := time.Now()

	switch granularity {
	case "minute":
		currentRange = TimeRange{
			Start: now.Truncate(time.Minute),
			End:   now,
		}
		yearOverYearStart := now.AddDate(0, 0, -1).Truncate(time.Minute)
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(time.Minute),
		}

	case "hour":
		currentStart := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
		yearOverYearStart := currentStart.AddDate(0, 0, -1)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(time.Duration(now.Minute()) * time.Minute),
		}

	case "day":
		currentStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		yearOverYearStart := currentStart.AddDate(0, 0, -7)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(time.Duration(now.Hour())*time.Hour + time.Duration(now.Minute())*time.Minute),
		}

	case "week":
		// 获取当前日期在本月的第几周
		currentWeekOfMonth := weekOfMonth(now)
		// 计算本周的开始时间，从周一 00:00 开始
		offset := int(time.Monday - now.Weekday())
		if offset > 0 {
			offset = -6
		}
		currentWeekStart := time.Date(now.Year(), now.Month(), now.Day()+offset, 0, 0, 0, 0, now.Location())

		// 找到上个月的对应周
		previousMonth := now.AddDate(0, -1, 0)
		weeksInPrevMonth := weeksInMonth(previousMonth)

		// 如果上个月没有对应的周，则使用上个月的最后一周
		previousWeekOfMonth := currentWeekOfMonth
		if currentWeekOfMonth > weeksInPrevMonth {
			previousWeekOfMonth = weeksInPrevMonth
		}

		previousMonthWeekStart := time.Date(previousMonth.Year(), previousMonth.Month(), 1, 0, 0, 0, 0, now.Location())
		previousWeekStart := previousMonthWeekStart.AddDate(0, 0, (previousWeekOfMonth-1)*7)
		offset = int(time.Monday - previousWeekStart.Weekday())
		if offset > 0 {
			offset = -6
		}
		previousWeekStart = previousWeekStart.AddDate(0, 0, offset)
		previousWeekEnd := previousWeekStart.Add(time.Duration(now.Sub(currentWeekStart)))

		currentRange = TimeRange{
			Start: currentWeekStart,
			End:   now,
		}
		yearOverYearRange = TimeRange{
			Start: previousWeekStart,
			End:   previousWeekEnd,
		}

	case "month":
		currentStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		yearOverYearStart := currentStart.AddDate(-1, 0, 0)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(time.Duration(now.Sub(currentStart))),
		}

	case "quarter":
		month := (now.Month()-1)/3*3 + 1
		currentStart := time.Date(now.Year(), month, 1, 0, 0, 0, 0, now.Location())
		yearOverYearStart := currentStart.AddDate(-1, 0, 0)

		currentRange = TimeRange{
			Start: currentStart,
			End:   now,
		}
		yearOverYearRange = TimeRange{
			Start: yearOverYearStart,
			End:   yearOverYearStart.Add(time.Duration(now.Sub(currentStart))),
		}

	default:
		fmt.Println("Unsupported granularity")
	}

	return
}

func main() {
	granuarities := []string{"minute", "hour", "day", "week", "month", "quarter"}
	for _, granuarity := range granuarities {
		currentRange, previousRange := CompareTimeRanges(granuarity)
		fmt.Printf("Current %s range: %v - %v\n", granuarity, currentRange.Start, currentRange.End)
		fmt.Printf("Previous %s range: %v - %v\n", granuarity, previousRange.Start, previousRange.End)
	}
	fmt.Println("========================")
	for _, granuarity := range granuarities {
		currentRange, previousRange := CompareTimeRangesT(granuarity)
		fmt.Printf("Current %s range: %v - %v\n", granuarity, currentRange.Start, currentRange.End)
		fmt.Printf("Previous %s range: %v - %v\n", granuarity, previousRange.Start, previousRange.End)
	}

}
