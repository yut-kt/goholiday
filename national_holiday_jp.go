// Package goholiday provides Package ome simple calculation functions.
// These are functions to calculate business days.
package goholiday

import (
	"time"

	"github.com/yut-kt/goholiday/nholidays/jp"
)

const dFmt = "2006-01-02"

var jpHoliday = Goholiday{jp.New(), map[string]struct{}{}}

// SetUniqueHolidays is a function to set unique holidays.
// deprecated:
// Moved to goholiday.go
func SetUniqueHolidays(ts []time.Time) {
	jpHoliday.uniqueHolidays = map[string]struct{}{}
	jpHoliday.SetUniqueHolidays(ts)
}

// IsNationalHoliday is a function to decide whether t given national holiday.
// deprecated:
// Moved to goholiday.go
func IsNationalHoliday(t time.Time) bool {
	return jpHoliday.IsNationalHoliday(t)
}

// IsHoliday is a function to decide whether t given holiday.
// deprecated:
// Moved to goholiday.go
func IsHoliday(t time.Time) bool {
	return jpHoliday.IsHoliday(t)
}

// IsBusinessDay is a function to decide whether t given business day.
// deprecated:
// Moved to goholiday.go
func IsBusinessDay(t time.Time) bool {
	return jpHoliday.IsBusinessDay(t)
}

// BusinessDaysBefore is a function that calculates bds business days before given t
// deprecated:
// Moved to goholiday.go
func BusinessDaysBefore(t time.Time, bds int) time.Time {
	return jpHoliday.BusinessDaysBefore(t, bds)
}

// BusinessDaysAfter is a function that calculates bds business days after given t
// deprecated:
// Moved to goholiday.go
func BusinessDaysAfter(t time.Time, bds int) time.Time {
	return jpHoliday.BusinessDaysAfter(t, bds)
}
