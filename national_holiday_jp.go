// Packages goholiday provides Package ome simple calcuration functions.
// These are functions to calculate business days.
package goholiday

import (
	"sort"
	"time"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/yut-kt/goholiday/config"
	"github.com/yut-kt/goholiday/entity"
)

// IsNationalHoliday is a function to decide whether t given national holiday.
func IsNationalHoliday(t time.Time) bool {
	nationalHolidays, err := fetchNationalHolidays()
	if err != nil {
		panic(err)
	}

	index := sort.Search(len(nationalHolidays), func(i int) bool { return nationalHolidays[i].IsOrAfter(t) })
	return nationalHolidays[index].Date == t.Format(config.DateFormat)
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

func travelBusinessDays(date time.Time, businessDays int, isTravelFuture bool) time.Time {
	course := map[bool]int{true: 1, false: -1}[isTravelFuture]

	for tbds := 0; tbds != businessDays; {
		date = date.AddDate(0, 0, course)
		if IsBusinessDay(date) {
			tbds++
		}
	}
	return date
}

func fetchNationalHolidays() (entity.NationalHolidays, error) {
	buf, err := ioutil.ReadFile("data/national_holidays.yaml")
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
