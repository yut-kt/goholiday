// Packages goholiday provides Package ome simple calcuration functions.
// These are functions to calculate business days.
package goholiday

import (
	"sort"
	"time"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"gopkg.in/yaml.v2"
	"github.com/yut-kt/goholiday/config"
	"github.com/yut-kt/goholiday/entity"
	"github.com/yut-kt/goholiday/data"
=======
	"github.com/goholiday/entity"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/goholiday/config"
>>>>>>> [add]godoc
=======
=======
	"os"
>>>>>>> [fix]import os
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/yut-kt/goholiday/config"
	"github.com/yut-kt/goholiday/entity"
>>>>>>> [fix]dir path
=======
	"gopkg.in/yaml.v2"
	"github.com/yut-kt/goholiday/config"
	"github.com/yut-kt/goholiday/entity"
	"github.com/yut-kt/goholiday/data"
>>>>>>> [update]Support for binaryization
)

// IsNationalHoliday is a function to decide whether t given national holiday.
func IsNationalHoliday(t time.Time) bool {
	nhs, err := fetchNationalHolidays()
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	t = t.In(config.JST)
=======
>>>>>>> [update]Reduced count of file reads
	index := sort.Search(len(nhs), func(i int) bool { return nhs[i].IsOrAfter(t) })
	return nhs[index].Date == t.Format(config.DateFormat)
}

// IsBusinessDay is a function to decide whether t given business day.
func IsBusinessDay(t time.Time) bool {
	return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday && !IsNationalHoliday(t)
}

// BusinessDaysBefore is a function that calculates bds business days before given t
func BusinessDaysBefore(t time.Time, bds int) time.Time {
<<<<<<< HEAD
	return travelBusinessDays(t.In(config.JST), bds, false)
=======
	return travelBusinessDays(t, bds, false)
>>>>>>> [add]godoc
}

// BusinessDaysAfter is a function that calculates bds business days after given t
func BusinessDaysAfter(t time.Time, bds int) time.Time {
<<<<<<< HEAD
	return travelBusinessDays(t.In(config.JST), bds, true)
=======
	return travelBusinessDays(t, bds, true)
>>>>>>> [add]godoc
}

<<<<<<< HEAD
func travelBusinessDays(t time.Time, bds int, isFuture bool) time.Time {
=======
func isNationalHoliday(t time.Time, nhs entity.NationalHolidays) bool {
	index := sort.Search(len(nhs), func(i int) bool { return nhs[i].IsOrAfter(t) })
	return nhs[index].Date == t.Format(config.DateFormat)
}

func isBusinessDay(t time.Time, hs entity.NationalHolidays) bool {
	return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday && !isNationalHoliday(t, hs)
}

func travelBusinessDays(d time.Time, bds int, isFuture bool) time.Time {
	course := map[bool]int{true: 1, false: -1}[isFuture]
>>>>>>> [update]Reduced count of file reads
	nhs, err := fetchNationalHolidays()
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
<<<<<<< HEAD
	course := map[bool]int{true: 1, false: -1}[isFuture]
	for tbds := 0; tbds != bds; {
		if t = t.AddDate(0, 0, course); isBusinessDay(t, nhs) {
=======
	for tbds := 0; tbds != businessDays; {
		date = date.AddDate(0, 0, course)
		if IsBusinessDay(date) {
>>>>>>> [add]godoc
			tbds++
		}
	}
	return t
=======
	for tbds := 0; tbds != bds; {
		d = d.AddDate(0, 0, course)
		if isBusinessDay(d, nhs) {
			tbds++
		}
	}
	return d
>>>>>>> [update]Reduced count of file reads
}

<<<<<<< HEAD
func isNationalHoliday(t time.Time, nhs entity.NationalHolidays) bool {
	index := sort.Search(len(nhs), func(i int) bool { return nhs[i].IsOrAfter(t) })
	return nhs[index].Date == t.Format(config.DateFormat)
}

func isBusinessDay(t time.Time, hs entity.NationalHolidays) bool {
	return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday && !isNationalHoliday(t, hs)
}
=======
func fetchNationalHolidays() (entity.NationalHolidays, error) {
<<<<<<< HEAD
	goholidayRoot := os.Getenv("GOPATH") + "/src/github.com/yut-kt/goholiday"
	buf, err := ioutil.ReadFile(goholidayRoot + "/data/national_holidays_jp.yaml")
	if err != nil {
		return nil, err
	}
>>>>>>> [fix]open path

func fetchNationalHolidays() (entity.NationalHolidays, error) {
	var nh entity.NationalHolidays
	if err := yaml.Unmarshal(data.NationalHolidaysJpYaml, &nh); err != nil {
=======
	var nh entity.NationalHolidays
	err := yaml.Unmarshal(data.NationalHolidaysJpYaml, &nh)
	if err != nil {
>>>>>>> [update]Support for binaryization
		return nil, err
	}
	return nh, nil
}