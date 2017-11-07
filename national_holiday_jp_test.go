package goholiday

import (
	"testing"
	"time"
	"github.com/yut-kt/goholiday/config"
	"github.com/yut-kt/goholiday/nholidays"
)

const (
	Fmt    = config.DateFormat
	ErrFmt = "Checking %s is incorrect."
)

func TestIsNationalHoliday(t *testing.T) {
	nationalHolidayOfWeekday, err1 := time.Parse(Fmt, "2017-10-09")
	nationalHolidayOfHoliday, err2 := time.Parse(Fmt, "2017-09-23")
	weekday, err3 := time.Parse(Fmt, "2017-10-10")
	holiday, err4 := time.Parse(Fmt, "2017-09-24")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date.")
	}

	if !IsNationalHoliday(nationalHolidayOfWeekday) {
		t.Errorf(ErrFmt, "nationalHolidayOfWeekday")
		t.Log("nationalHolidayOfWeekday: ", nationalHolidayOfWeekday.String())
	}
	if !IsNationalHoliday(nationalHolidayOfHoliday) {
		t.Errorf(ErrFmt, "nationalHolidayOfHoliday")
		t.Log("nationalHolidayOfHoliday: ", nationalHolidayOfHoliday.String())
	}
	if IsNationalHoliday(weekday) {
		t.Errorf(ErrFmt, "weekday")
		t.Log("weekday: ", weekday.String())
	}
	if IsNationalHoliday(holiday) {
		t.Errorf(ErrFmt, "holiday")
		t.Log("holiday: ", holiday.String())
	}
}

func BenchmarkIsNationalHoliday(b *testing.B) {
	var date []time.Time
	for n := 0; n < b.N; n++ {
		date = append(date, time.Now().AddDate(0, 0, n))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsNationalHoliday(date[i])
	}
}

func TestIsBusinessDay(t *testing.T) {
	nationalHolidayOfWeekday, err1 := time.Parse(Fmt, "2017-10-09")
	nationalHolidayOfHoliday, err2 := time.Parse(Fmt, "2017-09-23")
	weekday, err3 := time.Parse(Fmt, "2017-10-10")
	holiday, err4 := time.Parse(Fmt, "2017-09-24")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date.")
	}

	if IsBusinessDay(nationalHolidayOfWeekday) {
		t.Errorf(ErrFmt, "nationalHolidayOfWeekday")
		t.Log("nationalHolidayOfWeekday: ", nationalHolidayOfWeekday.String())
	}
	if IsBusinessDay(nationalHolidayOfHoliday) {
		t.Errorf(ErrFmt, "nationalHolidayOfHoliday")
		t.Log("nationalHolidayOfHoliday: ", nationalHolidayOfHoliday.String())
	}
	if !IsBusinessDay(weekday) {
		t.Errorf(ErrFmt, "weekday")
		t.Log("weekday: ", weekday.String())
	}
	if IsBusinessDay(holiday) {
		t.Errorf(ErrFmt, "holiday")
		t.Log("holiday: ", holiday.String())
	}
}

func BenchmarkIsWeekDay(b *testing.B) {
	var date []time.Time
	for n := 0; n < b.N; n++ {
		date = append(date, time.Now().AddDate(0, 0, n))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsBusinessDay(date[i])
	}
}

func TestBusinessDaysBefore(t *testing.T) {
	date1, err1 := time.Parse(Fmt, "2017-10-11")
	date2, err2 := time.Parse(Fmt, "2017-10-10")
	date3, err3 := time.Parse(Fmt, "2017-10-09")
	date4, err4 := time.Parse(Fmt, "2017-10-08")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date.")
	}

	if BusinessDaysBefore(date1, 1).Format(Fmt) != "2017-10-10" {
		t.Errorf(ErrFmt, "date1-2")
		t.Log("date1-1: ", BusinessDaysBefore(date1, 1).Format(Fmt))
	}
	if BusinessDaysBefore(date1, 2).Format(Fmt) != "2017-10-06" {
		t.Errorf(ErrFmt, "date1-2")
		t.Log("date1-2: ", BusinessDaysBefore(date1, 2).Format(Fmt))
	}
	if BusinessDaysBefore(date2, 1).Format(Fmt) != "2017-10-06" {
		t.Errorf(ErrFmt, "date2-1")
		t.Log("date2-1: ", BusinessDaysBefore(date2, 1).Format(Fmt))
	}
	if BusinessDaysBefore(date2, 2).Format(Fmt) != "2017-10-05" {
		t.Errorf(ErrFmt, "date2-2")
		t.Log("date2-2: ", BusinessDaysBefore(date2, 2).Format(Fmt))
	}
	if BusinessDaysBefore(date3, 1).Format(Fmt) != "2017-10-06" {
		t.Errorf(ErrFmt, "date3")
		t.Log("date3: ", BusinessDaysBefore(date3, 1).Format(Fmt))
	}
	if BusinessDaysBefore(date4, 1).Format(Fmt) != "2017-10-06" {
		t.Errorf(ErrFmt, "date4")
		t.Log("date4: ", BusinessDaysBefore(date4, 1).Format(Fmt))
	}
}

func BenchmarkBusinessDaysBefore(b *testing.B) {
	var date []time.Time
	for n := 0; n < b.N; n++ {
		date = append(date, time.Now().AddDate(0, 0, n))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BusinessDaysBefore(date[i], i)
	}
}

func TestBusinessDaysAfter(t *testing.T) {
	date1, err1 := time.Parse(Fmt, "2017-10-05")
	date2, err2 := time.Parse(Fmt, "2017-10-06")
	date3, err3 := time.Parse(Fmt, "2017-10-07")
	date4, err4 := time.Parse(Fmt, "2017-10-09")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date.")
	}

	if BusinessDaysAfter(date1, 1).Format(Fmt) != "2017-10-06" {
		t.Errorf(ErrFmt, "date1-1")
		t.Log("date1-1: ", BusinessDaysAfter(date1, 1).Format(Fmt))
	}
	if BusinessDaysAfter(date1, 2).Format(Fmt) != "2017-10-10" {
		t.Errorf(ErrFmt, "date1-2")
		t.Log("date1-2: ", BusinessDaysAfter(date1, 2).Format(Fmt))
	}
	if BusinessDaysAfter(date2, 1).Format(Fmt) != "2017-10-10" {
		t.Errorf(ErrFmt, "date2-1")
		t.Log("date2-1: ", BusinessDaysAfter(date2, 1).Format(Fmt))
	}
	if BusinessDaysAfter(date2, 2).Format(Fmt) != "2017-10-11" {
		t.Errorf(ErrFmt, "date2-2")
		t.Log("date2-2: ", BusinessDaysAfter(date2, 2).Format(Fmt))
	}
	if BusinessDaysAfter(date3, 1).Format(Fmt) != "2017-10-10" {
		t.Errorf(ErrFmt, "date3")
		t.Log("date3: ", BusinessDaysAfter(date3, 1).Format(Fmt))
	}
	if BusinessDaysAfter(date4, 1).Format(Fmt) != "2017-10-10" {
		t.Errorf(ErrFmt, "date4")
		t.Log("date4: ", BusinessDaysAfter(date4, 1).Format(Fmt))
	}
}

func BenchmarkBusinessDaysAfter(b *testing.B) {
	var date []time.Time
	for n := 0; n < b.N; n++ {
		date = append(date, time.Now().AddDate(0, 0, n))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BusinessDaysAfter(date[i], i)
	}
}

func TestIsNationalHoliday2s(t *testing.T) {
	nholidays.JpYaml = []byte(`fail`)
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf(ErrFmt, "panic")
		}
	}()
	IsNationalHoliday(time.Now())
}

func TestBusinessDaysAfter2(t *testing.T) {
	nholidays.JpYaml = []byte(`fail`)
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf(ErrFmt, "panic")
		}
	}()
	BusinessDaysAfter(time.Now(), 1)
}
