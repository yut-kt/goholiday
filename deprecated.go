// Package goholiday provides Package ome simple calculation functions.
// These are functions to calculate business days.
package goholiday

import (
	"time"

	"github.com/yut-kt/goholiday/nholidays/jp"
)

var jpHoliday = Goholiday{jp.New(), map[string]struct{}{}}

// Deprecated: Moved to Goholiday instance (goholiday.go).
// SetUniqueHolidays is a function to set unique holidays.
func SetUniqueHolidays(ts []time.Time) {
	jpHoliday.uniqueHolidays = map[string]struct{}{}
	jpHoliday.SetUniqueHolidays(ts)
}

// Deprecated: Moved to Goholiday instance (goholiday.go).
// IsNationalHoliday is a function to decide whether t given national holiday.
func IsNationalHoliday(t time.Time) bool {
	return jpHoliday.IsNationalHoliday(t)
}

// Deprecated: Moved to Goholiday instance (goholiday.go).
// IsHoliday is a function to decide whether t given holiday.
func IsHoliday(t time.Time) bool {
	return jpHoliday.IsHoliday(t)
}

// Deprecated: Moved to Goholiday instance (goholiday.go).
// IsBusinessDay is a function to decide whether t given business day.
func IsBusinessDay(t time.Time) bool {
	return jpHoliday.IsBusinessDay(t)
}

// Deprecated: Moved to Goholiday instance (goholiday.go).
// BusinessDaysBefore is a function that calculates bds business days before given t
func BusinessDaysBefore(t time.Time, bds int) time.Time {
	return jpHoliday.BusinessDaysBefore(t, bds)
}

// Deprecated: Moved to Goholiday instance (goholiday.go).
// BusinessDaysAfter is a function that calculates bds business days after given t
func BusinessDaysAfter(t time.Time, bds int) time.Time {
	return jpHoliday.BusinessDaysAfter(t, bds)
}
