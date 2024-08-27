package cn

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCompareTimeRangesMoM(t *testing.T) {
	type args struct {
		granularity string
		now         time.Time
		originStart time.Time
	}
	originStart, _ := time.Parse("2006-01-02 15:04:05", "2024-07-31 18:10:22")
	now, _ := time.Parse("2006-01-02 15:04:05", "2024-07-31 18:40:59")
	fmt.Println(now)
	tests := []struct {
		args args
	}{
		{
			args: args{
				granularity: ByMinute,
				now:         now,
				originStart: originStart,
			},
		},
		{
			args: args{
				granularity: ByHour,
				now:         now,
				originStart: originStart,
			},
		},
		{
			args: args{
				granularity: ByDay,
				now:         now,
				originStart: originStart,
			},
		},
		{
			args: args{
				granularity: ByWeek,
				now:         now,
				originStart: originStart,
			},
		},
		{
			args: args{
				granularity: ByMonth,
				now:         now,
				originStart: originStart,
			},
		},
		{
			args: args{
				granularity: ByQuarter,
				now:         now,
				originStart: originStart,
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotCurrentRange, gotPreviousRange := CompareTimeRangesMoM(tt.args.granularity, tt.args.originStart, tt.args.now)
			fmt.Println(tt.args.granularity, gotCurrentRange, gotPreviousRange)
		})
	}
}

func TestCompareTimeRangesYoY(t *testing.T) {
	type args struct {
		granularity string
		now         time.Time
		originStart time.Time
	}
	originStart, _ := time.Parse("2006-01-02 15:04:05", "2024-08-17 18:10:22")
	now, _ := time.Parse("2006-01-02 15:04:05", "2024-08-26 18:40:59")
	fmt.Println(now)
	tests := []struct {
		args args
	}{
		{
			args: args{
				granularity: ByMinute,
				now:         now,
				originStart: originStart,
			},
		},
		{
			args: args{
				granularity: ByHour,
				now:         now,
				originStart: originStart,
			},
		},
		{
			args: args{
				granularity: ByDay,
				now:         now,
				originStart: originStart,
			},
		},
		{
			args: args{
				granularity: ByWeek,
				now:         now,
				originStart: originStart,
			},
		},
		{
			args: args{
				granularity: ByMonth,
				now:         now,
				originStart: originStart,
			},
		},
		{
			args: args{
				granularity: ByQuarter,
				now:         now,
				originStart: originStart,
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotCurrentRange, gotPreviousRange := CompareTimeRangesYoY(tt.args.granularity, tt.args.originStart, tt.args.now)
			fmt.Println(tt.args.granularity, gotCurrentRange, gotPreviousRange)
		})
	}
}
