package goholiday_test

import (
	"time"
	"github.com/yut-kt/goholiday"
	"github.com/yut-kt/goholiday/nholidays/jp"
	"fmt"
)

func getLocaleJP() *time.Location {
	locale, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	return locale
}

func ExampleNew() {
	gh := goholiday.New(jp.New())
	fmt.Printf("%T", gh)
	// Output:
	// *goholiday.Goholiday
}

func ExampleGoholiday_IsNationalHoliday() {
	localeJP := getLocaleJP()
	targets := []time.Time{
		// national holiday of weekday"
		time.Date(2017, 10, 9, 0, 0, 0, 0, localeJP),
		// national holiday of holiday
		time.Date(2017, 9, 23, 0, 0, 0, 0, localeJP),
		// business day
		time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP),
		// Sunday
		time.Date(2017, 9, 24, 0, 0, 0, 0, localeJP),
	}

	gh := goholiday.New(jp.New())
	for _, target := range targets {
		fmt.Println(gh.IsNationalHoliday(target))
	}
	// Output:
	// true
	// true
	// false
	// false
}

func ExampleGoholiday_IsHoliday() {
	localeJP := getLocaleJP()
	targets := []time.Time{
		// national holiday of weekday"
		time.Date(2017, 10, 9, 0, 0, 0, 0, localeJP),
		// national holiday of holiday
		time.Date(2017, 9, 23, 0, 0, 0, 0, localeJP),
		// business day
		time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP),
		// Sunday
		time.Date(2017, 9, 24, 0, 0, 0, 0, localeJP),
	}
	gh := goholiday.New(jp.New())
	for _, target := range targets {
		fmt.Println(gh.IsHoliday(target))
	}
	// Output:
	// true
	// true
	// false
	// true
}

func ExampleGoholiday_IsBusinessDay() {
	localeJP := getLocaleJP()
	targets := []time.Time{
		// national holiday of weekday"
		time.Date(2017, 10, 9, 0, 0, 0, 0, localeJP),
		// national holiday of holiday
		time.Date(2017, 9, 23, 0, 0, 0, 0, localeJP),
		// business day
		time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP),
		// Sunday
		time.Date(2017, 9, 24, 0, 0, 0, 0, localeJP),
	}
	gh := goholiday.New(jp.New())
	for _, target := range targets {
		fmt.Println(gh.IsBusinessDay(target))
	}
	// Output:
	// false
	// false
	// true
	// false
}

func ExampleGoholiday_BusinessDaysBefore() {
	localeJP := getLocaleJP()
	cases := []struct {
		date   time.Time
		days   int
		expect time.Time
	}{
		{
			date:   time.Date(2017, 10, 11, 0, 0, 0, 0, localeJP),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP),
		},
		{
			date:   time.Date(2017, 10, 11, 0, 0, 0, 0, localeJP),
			days:   2,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, localeJP),
		},
		{
			date:   time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, localeJP),
		},
		{
			date:   time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP),
			days:   2,
			expect: time.Date(2017, 10, 5, 0, 0, 0, 0, localeJP),
		},
		{
			date:   time.Date(2017, 10, 9, 0, 0, 0, 0, localeJP),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, localeJP),
		},
		{
			date:   time.Date(2017, 10, 8, 0, 0, 0, 0, localeJP),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, localeJP),
		},
	}

	gh := goholiday.New(jp.New())
	for _, c := range cases {
		fmt.Println(gh.BusinessDaysBefore(c.date, c.days).Format("2006-01-02") == c.expect.Format("2006-01-02"))
	}
	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}

func ExampleGoholiday_BusinessDaysAfter() {
	localeJP := getLocaleJP()
	cases := []struct {
		date   time.Time
		days   int
		expect time.Time
	}{
		{
			date:   time.Date(2017, 10, 5, 0, 0, 0, 0, localeJP),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, localeJP),
		},
		{
			date:   time.Date(2017, 10, 5, 0, 0, 0, 0, localeJP),
			days:   2,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP),
		},
		{
			date:   time.Date(2017, 10, 6, 0, 0, 0, 0, localeJP),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP),
		},
		{
			date:   time.Date(2017, 10, 6, 0, 0, 0, 0, localeJP),
			days:   2,
			expect: time.Date(2017, 10, 11, 0, 0, 0, 0, localeJP),
		},
		{
			date:   time.Date(2017, 10, 7, 0, 0, 0, 0, localeJP),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP),
		},
		{
			date:   time.Date(2017, 10, 9, 0, 0, 0, 0, localeJP),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP),
		},
	}
	
	gh := goholiday.New(jp.New())
	for _, c := range cases {
		fmt.Println(gh.BusinessDaysAfter(c.date, c.days).Format("2006-01-02") == c.expect.Format("2006-01-02"))
	}
	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}

func ExampleGoholiday_SetUniqueHolidays() {
	localeJP := getLocaleJP()
	gh := goholiday.New(jp.New())
	gh.SetUniqueHolidays([]time.Time{time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)})
	date := time.Date(2017, 10, 10, 0, 0, 0, 0, localeJP)
	fmt.Println(gh.IsHoliday(date))
	fmt.Println(gh.IsBusinessDay(date))
	fmt.Println(gh.IsNationalHoliday(date))
	// Output:
	// true
	// false
	// false
}
