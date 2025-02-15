package goholiday_test

import (
	"testing"
	"time"

	"github.com/yut-kt/goholiday/nholidays"
	"github.com/yut-kt/goholiday/nholidays/jp"
	"github.com/yut-kt/goholiday/nholidays/uk"

	"github.com/yut-kt/goholiday"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Japan
		goholiday.New(jp.New())
		// England in UK
		goholiday.New(uk.NewEngland())
		// Original Schedule
		goholiday.New(nholidays.New(
			map[time.Weekday]struct{}{},
			map[string]string{"2022-05-01": "my holiday"},
		))
	}
}

func getDatesBy2020() []time.Time {
	days := 366
	dates := make([]time.Time, days)
	for i := 0; i < days; i++ {
		dates[i] = time.Date(2020, 01, i+1, 0, 0, 0, 0, time.Local)
	}
	return dates
}

func BenchmarkGoholiday_IsNationalHoliday(b *testing.B) {
	dates := getDatesBy2020()
	gh := goholiday.New(jp.New())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, date := range dates {
			gh.IsNationalHoliday(date)
		}
	}
}

func BenchmarkGoholiday_IsHoliday(b *testing.B) {
	dates := getDatesBy2020()
	gh := goholiday.New(jp.New())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, date := range dates {
			gh.IsHoliday(date)
		}
	}
}

func BenchmarkGoholiday_IsBusinessDay(b *testing.B) {
	dates := getDatesBy2020()
	gh := goholiday.New(jp.New())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, date := range dates {
			gh.IsBusinessDay(date)
		}
	}
}

func BenchmarkGoholiday_SetUniqueHolidays(b *testing.B) {
	dates := getDatesBy2020()
	gh := goholiday.New(jp.New())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gh.SetUniqueHolidays(dates)
	}
}

func BenchmarkGoholiday_BusinessDaysAfter(b *testing.B) {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	afterDays := 365
	gh := goholiday.New(jp.New())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gh.BusinessDaysAfter(start, afterDays)
	}
}

func BenchmarkGoholiday_BusinessDaysBefore(b *testing.B) {
	start := time.Date(2020, 12, 31, 0, 0, 0, 0, time.Local)
	beforeDays := 365
	gh := goholiday.New(jp.New())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gh.BusinessDaysBefore(start, beforeDays)
	}
}
