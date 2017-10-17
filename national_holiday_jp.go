// Packages goholiday provides Package ome simple calcuration functions.
// These are functions to calculate business days.
package goholiday

import (
	"sort"
	"time"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/yut-kt/goholiday/config"
	"github.com/yut-kt/goholiday/entity"
)

// IsNationalHoliday is a function to decide whether t given national holiday.
func IsNationalHoliday(t time.Time) bool {
	nhs, err := fetchNationalHolidays()
	if err != nil {
		panic(err)
	}

	index := sort.Search(len(nhs), func(i int) bool { return nhs[i].IsOrAfter(t) })
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

func isNationalHoliday(t time.Time, nhs entity.NationalHolidays) bool {
	index := sort.Search(len(nhs), func(i int) bool { return nhs[i].IsOrAfter(t)})
	return nhs[index].Date == t.Format(config.DateFormat)
}

func isBusinessDay(t time.Time, hs entity.NationalHolidays) bool {
	return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday && !isNationalHoliday(t, hs)
}

func travelBusinessDays(d time.Time, bds int, isFuture bool) time.Time {
	course := map[bool]int{true: 1, false: -1}[isFuture]
	nhs, err := fetchNationalHolidays()
	if err != nil {
		panic(err)
	}

	for tbds := 0; tbds != bds; {
		d = d.AddDate(0, 0, course)
		if isBusinessDay(d, nhs) {
			tbds++
		}
	}
	return d
}

func fetchNationalHolidays() (entity.NationalHolidays, error) {
	goholidayRoot := os.Getenv("GOPATH") + "/src/github.com/yut-kt/goholiday"
	buf, err := ioutil.ReadFile(goholidayRoot + "/data/national_holidays_jp.yaml")
	if err != nil {
		return nil, err
	}

	var nh entity.NationalHolidays
	err = yaml.Unmarshal(buf, &nh)
	if err != nil {
		return nil, err
	}

	return nh, nil
}
