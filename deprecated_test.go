package goholiday_test

import (
	"testing"
	"time"

	"github.com/yut-kt/goholiday"
)

const errFmt = "Checking %v is incorrect."

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
		if goholiday.IsNationalHoliday(c.date) != c.expect {
			t.Errorf(errFmt, c.dateType)
		}
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
		if goholiday.IsHoliday(c.date) != c.expect {
			t.Errorf(errFmt, c.dateType)
		}
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
		if goholiday.IsBusinessDay(c.date) != c.expect {
			t.Errorf(errFmt, c.dateType)
		}
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
		if goholiday.BusinessDaysBefore(c.date, c.days).Format("2006-01-02") != c.expect.Format("2006-01-02") {
			t.Errorf(errFmt, c.date.String())
		}
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
		if goholiday.BusinessDaysAfter(c.date, c.days).Format("2006-01-02") != c.expect.Format("2006-01-02") {
			t.Errorf(errFmt, c.date.String())
		}
	}
}

// set unique holidays
func TestIsHolidayUnique(t *testing.T) {
	goholiday.SetUniqueHolidays([]time.Time{time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)})

	date := time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)
	if !goholiday.IsHoliday(date) {
		t.Errorf(errFmt, "unique holiday")
	}
}

// set unique holidays
func TestIsBusinessDayUnique(t *testing.T) {
	goholiday.SetUniqueHolidays([]time.Time{time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)})

	date := time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)
	if goholiday.IsBusinessDay(date) {
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
	goholiday.SetUniqueHolidays(uhs)

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
		if goholiday.BusinessDaysBefore(c.date, c.days).Format("2006-01-02") != c.expect.Format("2006-01-02") {
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
	goholiday.SetUniqueHolidays(uhs)

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
		if goholiday.BusinessDaysAfter(c.date, c.days).Format("2006-01-02") != c.expect.Format("2006-01-02") {
			t.Errorf(errFmt, c.date.String())
		}
	}
}
