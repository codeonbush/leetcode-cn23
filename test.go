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
func CompareTimeRanges(granularity string, now time.Time) (currentRange TimeRange, previousRange TimeRange) {

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
		// 本月开始时间
		currentStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

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
	case "quarter":
		// 计算本季度的开始月份
		month := (now.Month()-1)/3*3 + 1
		currentStart := time.Date(now.Year(), month, 1, 0, 0, 0, 0, now.Location())
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
			End:   previousEnd,
		}

		// 确保上季度的结束时间不超过当前时间
		if previousRange.End.After(now) {
			previousRange.End = now
		}

	default:
		fmt.Println("Unsupported granularity")
	}

	return
}

// CompareTimeRanges 返回当前时间范围与同比时间范围
func CompareTimeRangesT(granularity string, now time.Time) (currentRange TimeRange, yearOverYearRange TimeRange) {

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
		//这段代码计算同比两个时间范围的开始和结束时间，修改这段代码逻辑，
		//本周至当前分钟，对比本周向前取至四周前的对应周的对应时间。
		//如：在 8.5，时间范围选8.1~8.5，粒度选「按周」，本周期为8.5 0:00到当前分钟，月同比为 7.8 0:00 到对应当前分钟
		// 获取当前时间
		// 本周开始时间，从周一 00:00 开始
		startOfWeek := now.Truncate(24 * time.Hour).Add(-(time.Duration(now.Weekday()-time.Monday) * 24 * time.Hour))

		// 当前分钟
		currentMinute := now

		// 计算四周前开始时间（从周一 00:00 开始至四周前的对应时间）
		fourWeeksAgoStart := startOfWeek.AddDate(0, 0, -28) // 向前计算 28 天
		fourWeeksAgoMinute := fourWeeksAgoStart.Truncate(24 * time.Hour).Add(currentMinute.Sub(currentMinute.Truncate(time.Minute)))

		// 设置当前时间区间
		currentRange = TimeRange{
			Start: startOfWeek,
			End:   currentMinute,
		}

		// 设置四周前的同比时间区间
		yearOverYearRange = TimeRange{
			Start: fourWeeksAgoMinute,
			End:   fourWeeksAgoMinute.Add(currentMinute.Sub(startOfWeek)),
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
			End:   yearOverYearStart.Add(now.Sub(currentStart)),
		}

	case "quarter":
		// 计算当前季度的开始月份
		month := (now.Month()-1)/3*3 + 1
		currentStart := time.Date(now.Year(), month, 1, 0, 0, 0, 0, now.Location())

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
			End:   yearOverYearStart.Add(time.Duration(now.Sub(currentStart))),
		}

	default:
		fmt.Println("Unsupported granularity")
	}

	return
}

func main() {
	granuarities := []string{"minute", "hour", "day", "week", "month", "quarter"}
	now, err := time.ParseInLocation("2006-01-02 15:04:05", "2024-07-31 15:18:01", time.Local)
	fmt.Println(err)
	for _, granuarity := range granuarities {
		currentRange, previousRange := CompareTimeRanges(granuarity, now)
		fmt.Printf("Current %s range: %v - %v\n", granuarity, currentRange.Start, currentRange.End)
		fmt.Printf("Previous %s range: %v - %v\n", granuarity, previousRange.Start, previousRange.End)
	}
	fmt.Println("========================")
	for _, granuarity := range granuarities {
		currentRange, previousRange := CompareTimeRangesT(granuarity, now)
		fmt.Printf("Current %s range: %v - %v\n", granuarity, currentRange.Start, currentRange.End)
		fmt.Printf("Previous %s range: %v - %v\n", granuarity, previousRange.Start, previousRange.End)
	}

}
