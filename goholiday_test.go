package goholiday

import (
	"testing"
	"time"
	"github.com/yut-kt/goholiday/nholidays/jp"
)

func TestGoholiday_IsNationalHoliday(t *testing.T) {
	cases := []struct {
		date     time.Time
		expect   bool
		dateType string
	}{
		{
			date:     time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
			expect:   true,
			dateType: "national holiday of weekday",
		},
		{
			date:     time.Date(2017, 9, 23, 0, 0, 0, 0, time.Local),
			expect:   true,
			dateType: "national holiday of holiday",
		},
		{
			date:     time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
			expect:   false,
			dateType: "weekday",
		},
		{
			date:     time.Date(2017, 9, 24, 0, 0, 0, 0, time.Local),
			expect:   false,
			dateType: "holiday",
		},
	}

	g := Goholiday{jp.New(), nil}
	for _, c := range cases {
		if g.IsNationalHoliday(c.date) != c.expect {
			t.Errorf(errFmt, c.dateType)
		}
	}
}

func BenchmarkGoholiday_IsNationalHoliday(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	g := Goholiday{jp.New(), nil}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.IsNationalHoliday(date[i])
	}
}

func BenchmarkGoholiday_IsNationalHolidayV2(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	g := Goholiday{jp.New(), nil}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.IsNationalHoliday(date[i])
	}
}
