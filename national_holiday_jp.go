// Packages goholiday provides Package ome simple calcuration functions.
// These are functions to calculate business days.
package goholiday

import (
	"sort"
	"time"
	"gopkg.in/yaml.v2"
	"github.com/yut-kt/goholiday/config"
	"github.com/yut-kt/goholiday/entity"
	"github.com/yut-kt/goholiday/nholidays"
)

var nhs entity.NationalHolidays
func init() {
	if err := yaml.Unmarshal(nholidays.JpYaml, &nhs); err != nil {
		panic(err)
	}
}

// IsNationalHoliday is a function to decide whether t given national holiday.
func IsNationalHoliday(t time.Time) bool {
	t = t.In(config.JST)
	index := sort.Search(len(nhs), func(i int) bool { return nhs[i].IsOrAfter(t) })
	if index == len(nhs) {
		return false
	}
	return nhs[index].Date == t.Format(config.DateFormat)
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
		if t = t.AddDate(0, 0, course); IsBusinessDay(t) {
			tbds++
		}
	}
	return t
}
