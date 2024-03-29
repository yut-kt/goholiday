package goholiday_test

import (
	"fmt"
	"time"

	"github.com/yut-kt/goholiday/nholidays"

	"github.com/yut-kt/goholiday/nholidays/uk"

	"github.com/yut-kt/goholiday"
	"github.com/yut-kt/goholiday/nholidays/jp"
)

func ExampleGoholiday() {
	// Japan Schedule
	ghj := goholiday.New(jp.New())
	fmt.Println(ghj.IsNationalHoliday(time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)))

	// England Schedule
	ghe := goholiday.New(uk.NewEngland())
	fmt.Println(ghe.IsNationalHoliday(time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)))

	// Original Schedule
	mySchedule := nholidays.New(
		map[time.Weekday]struct{}{},
		map[string]string{"2022-05-01": "my holiday"},
	)
	ghMine := goholiday.New(mySchedule)
	fmt.Println(ghMine.IsNationalHoliday(time.Date(2022, 5, 1, 0, 0, 0, 0, time.Local)))

	// Output:
	// true
	// true
	// true
}

func ExampleNew() {
	gh := goholiday.New(jp.New())
	fmt.Printf("%T", gh)
	// Output:
	// *goholiday.Goholiday
}

func ExampleGoholiday_IsNationalHoliday() {
	targets := []time.Time{
		// national holiday of weekday"
		time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
		// national holiday of holiday
		time.Date(2017, 9, 23, 0, 0, 0, 0, time.Local),
		// business day
		time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		// Sunday
		time.Date(2017, 9, 24, 0, 0, 0, 0, time.Local),
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
	targets := []time.Time{
		// national holiday of weekday"
		time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
		// national holiday of holiday
		time.Date(2017, 9, 23, 0, 0, 0, 0, time.Local),
		// business day
		time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		// Sunday
		time.Date(2017, 9, 24, 0, 0, 0, 0, time.Local),
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
	targets := []time.Time{
		// national holiday of weekday"
		time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
		// national holiday of holiday
		time.Date(2017, 9, 23, 0, 0, 0, 0, time.Local),
		// business day
		time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		// Sunday
		time.Date(2017, 9, 24, 0, 0, 0, 0, time.Local),
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
	cases := []struct {
		date   time.Time
		days   int
		expect time.Time
	}{
		{
			date:   time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 5, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 8, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
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
	cases := []struct {
		date   time.Time
		days   int
		expect time.Time
	}{
		{
			date:   time.Date(2017, 10, 5, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 5, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local),
			days:   2,
			expect: time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 7, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
		},
		{
			date:   time.Date(2017, 10, 9, 0, 0, 0, 0, time.Local),
			days:   1,
			expect: time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local),
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
	gh := goholiday.New(jp.New())
	gh.SetUniqueHolidays([]time.Time{time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)})
	date := time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)
	fmt.Println(gh.IsHoliday(date))
	fmt.Println(gh.IsBusinessDay(date))
	fmt.Println(gh.IsNationalHoliday(date))
	// Output:
	// true
	// false
	// false
}
