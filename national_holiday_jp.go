// Packages goholiday provides Package ome simple calcuration functions.
// These are functions to calculate business days.
package goholiday

import (
	"github.com/yut-kt/goholiday/config"
	"github.com/yut-kt/goholiday/entity"
	"github.com/yut-kt/goholiday/nholidays"
	"gopkg.in/yaml.v2"
	"sort"
	"time"
)

const DFmt = config.DateFormat

var nhs entity.NationalHolidays
var uhs []time.Time

func init() {
	if err := yaml.Unmarshal(nholidays.JpYaml, &nhs); err != nil {
		panic(err)
	}
}

func setUniqueHolidays(ts []time.Time) {
	sort.Slice(ts, func(i, j int) bool { return ts[i].Before(ts[j]) })
	uhs = ts
}

// IsNationalHoliday is a function to decide whether t given national holiday.
func IsNationalHoliday(t time.Time) bool {
	t = t.In(config.JST)
	index := sort.Search(len(nhs), func(i int) bool { return nhs[i].IsOrAfter(t) })
	if index == len(nhs) {
		return false
	}
	return nhs[index].Date == t.Format(DFmt)
}

// IsBusinessDay is a function to decide whether t given business day.
func IsBusinessDay(t time.Time) bool {
	return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday && !IsNationalHoliday(t)
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
	t = t.In(config.JST)
	course := map[bool]int{true: 1, false: -1}[isFuture]
	for tbds := 0; tbds != bds; {
		if t = t.AddDate(0, 0, course); IsBusinessDay(t) && isNotUniqueHoliday(t) {
			tbds++
		}
	}
	return t
}

func isNotUniqueHoliday(t time.Time) bool {
	index := sort.Search(len(uhs), func(i int) bool { return uhs[i].After(t) || uhs[i].Format(DFmt) == t.Format(DFmt) })
	if index == len(uhs) {
		return true
	}
	return uhs[index].Format(DFmt) != t.Format(DFmt)
}
