package goholiday

import (
	"testing"
	"time"
)

const (
	errFmt = "Checking %v is incorrect."
)

func TestIsNationalHoliday(t *testing.T) {
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

	for _, c := range cases {
		if IsNationalHoliday(c.date) != c.expect {
			t.Errorf(errFmt, c.dateType)
		}
	}
}

func BenchmarkIsNationalHoliday(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsNationalHoliday(date[i])
	}
}

func TestIsHoliday(t *testing.T) {
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
			expect:   true,
			dateType: "holiday",
		},
	}

	for _, c := range cases {
		if IsHoliday(c.date) != c.expect {
			t.Errorf(errFmt, c.dateType)
		}
	}
}

func BenchmarkIsHoliday(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsHoliday(date[i])
	}
}

func TestIsBusinessDay(t *testing.T) {
	cases := []struct {
		date     time.Time
		expect   bool
		dateType string
	}{
		{
			date:     time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
			expect:   false,
			dateType: "national holiday of weekday",
		},
		{
			date:     time.Date(2017, 9, 23, 0, 0, 0, 0, time.Local),
			expect:   false,
			dateType: "national holiday of holiday",
		},
		{
			date:     time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
			expect:   true,
			dateType: "weekday",
		},
		{
			date:     time.Date(2017, 9, 24, 0, 0, 0, 0, time.Local),
			expect:   false,
			dateType: "holiday",
		},
	}

	for _, c := range cases {
		if IsBusinessDay(c.date) != c.expect {
			t.Errorf(errFmt, c.dateType)
		}
	}
}

func BenchmarkIsBusinessDay(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsBusinessDay(date[i])
	}
}

func TestBusinessDaysBefore(t *testing.T) {
	cases := []struct {
		date   time.Time
		days   int
		expect time.Time
	}{
		{
			date:   time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 5, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 8, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		},
	}

	for _, c := range cases {
		if BusinessDaysBefore(c.date, c.days).Format(dFmt) != c.expect.Format(dFmt) {
			t.Errorf(errFmt, c.date.String())
		}
	}
}

func BenchmarkBusinessDaysBefore(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BusinessDaysBefore(date[i], i)
	}
}

func TestBusinessDaysAfter(t *testing.T) {
	cases := []struct {
		date   time.Time
		days   int
		expect time.Time
	}{
		{
			date:   time.Date(2017, 10, 5, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 5, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 7, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		},
	}

	for _, c := range cases {
		if BusinessDaysAfter(c.date, c.days).Format(dFmt) != c.expect.Format(dFmt) {
			t.Errorf(errFmt, c.date.String())
		}
	}
}

func BenchmarkBusinessDaysAfter(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BusinessDaysAfter(date[i], i)
	}
}

// set unique holidays
func TestIsHolidayUnique(t *testing.T) {
	SetUniqueHolidays([]time.Time{time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)})

	date := time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)
	expect := true
	if IsHoliday(date) != expect {
		t.Errorf(errFmt, "unique holiday")
	}
}

// set unique holidays
func TestIsBusinessDayUnique(t *testing.T) {
	SetUniqueHolidays([]time.Time{time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)})

	date := time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)
	expect := false
	if IsBusinessDay(date) != expect {
		t.Errorf(errFmt, "unique holiday")
	}
}

// set unique holidays
func TestBusinessDaysBeforeUnique(t *testing.T) {
	uhs := []time.Time{
		time.Date(2017, 10, 5, 0, 0, 0, 0, time.Local),
		time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
	}
	SetUniqueHolidays(uhs)

	cases := []struct {
		date   time.Time
		days   int
		expect time.Time
	}{
		{
			date:   time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 4, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 3, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 4, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 3, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 4, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 8, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 4, 0, 0, 0, 0, time.Local),
		},
	}

	for _, c := range cases {
		if BusinessDaysBefore(c.date, c.days).Format(dFmt) != c.expect.Format(dFmt) {
			t.Errorf(errFmt, c.date.String())
		}
	}
}

// set unique holidays
func TestBusinessDaysAfterUnique(t *testing.T) {
	uhs := []time.Time{
		time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		time.Date(2017, 10, 12, 0, 0, 0, 0, time.Local),
		time.Date(2017, 10, 13, 0, 0, 0, 0, time.Local),
	}
	SetUniqueHolidays(uhs)

	cases := []struct {
		date   time.Time
		days   int
		expect time.Time
	}{
		{
			date:   time.Date(2017, 10, 5, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 5, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 16, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 7, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
		},
	}

	for _, c := range cases {
		if BusinessDaysAfter(c.date, c.days).Format(dFmt) != c.expect.Format(dFmt) {
			t.Errorf(errFmt, c.date.String())
		}
	}
}
