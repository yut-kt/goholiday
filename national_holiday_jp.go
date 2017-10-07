package holidayJp

import (
	"github.com/holiday-jp/config"
	"github.com/holiday-jp/entity"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
	"time"
)

func IsNationalHoliday(t time.Time) bool {
	nationalHolidays, err := fetchNationalHolidays()
	if err != nil {
		panic(err)
	}

	index := sort.Search(len(nationalHolidays), func(i int) bool { return nationalHolidays[i].IsOrAfter(t) })
	return nationalHolidays[index].Date == t.Format(config.DateFormat)
}

func IsWeekDay(t time.Time) bool {
	return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday && !IsNationalHoliday(t)
}

func BusinessDaysBefore(t time.Time, businessDays int) time.Time {
	return travelBusinessDays(t, businessDays, false)
}

func BusinessDaysAfter(t time.Time, businessDays int) time.Time {
	return travelBusinessDays(t, businessDays, true)
}

func travelBusinessDays(date time.Time, businessDays int, isTravelFuture bool) time.Time {
	course := map[bool]int{true: 1, false: -1}[isTravelFuture]

	for tbds := 0; tbds != businessDays; date = date.AddDate(0, 0, course) {
		if IsWeekDay(date) {
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
