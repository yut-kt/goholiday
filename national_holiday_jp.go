// Package goholiday provides Package ome simple calculation functions.
// These are functions to calculate business days.
package goholiday

import (
	"github.com/yut-kt/goholiday/config"
	"github.com/yut-kt/goholiday/nholidays"
	"sort"
	"time"
)

const dFmt = config.DateFormat

var (
	uhs []time.Time
	jst = config.JST
)

// SetUniqueHolidays is a function to set unique holidays.
func SetUniqueHolidays(ts []time.Time) {
	sort.Slice(ts, func(i, j int) bool { return ts[i].Before(ts[j]) })
	uhs = ts
}

// IsNationalHoliday is a function to decide whether t given national holiday.
func IsNationalHoliday(t time.Time) bool {
	t = t.In(jst)
	return searchNHoliday(t)
}

// IsHoliday is a function to decide whether t given holiday.
func IsHoliday(t time.Time) bool {
	t = t.In(jst)
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday || searchNHoliday(t) || isUniqueHoliday(t)
}

// IsBusinessDay is a function to decide whether t given business day.
func IsBusinessDay(t time.Time) bool {
	t = t.In(jst)
	return !IsHoliday(t)
}

func searchNHoliday(t time.Time) bool {
	_, exist := nholidays.Jp[t.Format(dFmt)]
	return exist
}

// BusinessDaysBefore is a function that calculates bds business days before given t
func BusinessDaysBefore(t time.Time, bds int) time.Time {
	return travelBusinessDays(t, bds, false)
}

// BusinessDaysAfter is a function that calculates bds business days after given t
func BusinessDaysAfter(t time.Time, bds int) time.Time {
	return travelBusinessDays(t, bds, true)
}

func travelBusinessDays(t time.Time, bds int, isFuture bool) time.Time {
	t = t.In(jst)
	course := map[bool]int{true: 1, false: -1}[isFuture]
	for tbds := 0; tbds != bds; {
		if t = t.AddDate(0, 0, course); IsBusinessDay(t) && !isUniqueHoliday(t) {
			tbds++
		}
	}
	return t
}

func isUniqueHoliday(t time.Time) bool {
	index := sort.Search(len(uhs), func(i int) bool { return uhs[i].After(t) || uhs[i].Format(dFmt) == t.Format(dFmt) })
	if index == len(uhs) {
		return false
	}
	return uhs[index].Format(dFmt) == t.Format(dFmt)
}
