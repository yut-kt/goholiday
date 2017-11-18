package goholiday

import (
	"testing"
	"time"
)

const (
	errFmt = "Checking %s is incorrect."
)

func TestIsNationalHoliday(t *testing.T) {
	nationalHolidayOfWeekday, err1 := time.Parse(dFmt, "2017-10-09")
	nationalHolidayOfHoliday, err2 := time.Parse(dFmt, "2017-09-23")
	weekday, err3 := time.Parse(dFmt, "2017-10-10")
	holiday, err4 := time.Parse(dFmt, "2017-09-24")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date.")
	}

	if !IsNationalHoliday(nationalHolidayOfWeekday) {
		t.Errorf(errFmt, "nationalHolidayOfWeekday")
		t.Log("nationalHolidayOfWeekday: ", nationalHolidayOfWeekday.String())
	}
	if !IsNationalHoliday(nationalHolidayOfHoliday) {
		t.Errorf(errFmt, "nationalHolidayOfHoliday")
		t.Log("nationalHolidayOfHoliday: ", nationalHolidayOfHoliday.String())
	}
	if IsNationalHoliday(weekday) {
		t.Errorf(errFmt, "weekday")
		t.Log("weekday: ", weekday.String())
	}
	if IsNationalHoliday(holiday) {
		t.Errorf(errFmt, "holiday")
		t.Log("holiday: ", holiday.String())
	}
}

func BenchmarkIsNationalHoliday(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsNationalHoliday(date[i])
	}
}

func TestIsHoliday(t *testing.T) {
	nationalHolidayOfWeekday, err1 := time.Parse(dFmt, "2017-10-09")
	nationalHolidayOfHoliday, err2 := time.Parse(dFmt, "2017-09-23")
	weekday, err3 := time.Parse(dFmt, "2017-10-10")
	holiday, err4 := time.Parse(dFmt, "2017-09-24")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date.")
	}

	if !IsHoliday(nationalHolidayOfWeekday) {
		t.Errorf(errFmt, "nationalHolidayOfWeekday")
		t.Log("nationalHolidayOfWeekday: ", nationalHolidayOfWeekday.String())
	}
	if !IsHoliday(nationalHolidayOfHoliday) {
		t.Errorf(errFmt, "nationalHolidayOfHoliday")
		t.Log("nationalHolidayOfHoliday: ", nationalHolidayOfHoliday.String())
	}
	if IsHoliday(weekday) {
		t.Errorf(errFmt, "weekday")
		t.Log("weekday: ", weekday.String())
	}
	if !IsHoliday(holiday) {
		t.Errorf(errFmt, "holiday")
		t.Log("holiday: ", holiday.String())
	}
}

func TestIsBusinessDay(t *testing.T) {
	nationalHolidayOfWeekday, err1 := time.Parse(dFmt, "2017-10-09")
	nationalHolidayOfHoliday, err2 := time.Parse(dFmt, "2017-09-23")
	weekday, err3 := time.Parse(dFmt, "2017-10-10")
	holiday, err4 := time.Parse(dFmt, "2017-09-24")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date.")
	}

	if IsBusinessDay(nationalHolidayOfWeekday) {
		t.Errorf(errFmt, "nationalHolidayOfWeekday")
		t.Log("nationalHolidayOfWeekday: ", nationalHolidayOfWeekday.String())
	}
	if IsBusinessDay(nationalHolidayOfHoliday) {
		t.Errorf(errFmt, "nationalHolidayOfHoliday")
		t.Log("nationalHolidayOfHoliday: ", nationalHolidayOfHoliday.String())
	}
	if !IsBusinessDay(weekday) {
		t.Errorf(errFmt, "weekday")
		t.Log("weekday: ", weekday.String())
	}
	if IsBusinessDay(holiday) {
		t.Errorf(errFmt, "holiday")
		t.Log("holiday: ", holiday.String())
	}
}

func BenchmarkIsWeekDay(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsBusinessDay(date[i])
	}
}

func TestBusinessDaysBefore(t *testing.T) {
	date1, err1 := time.Parse(dFmt, "2017-10-11")
	date2, err2 := time.Parse(dFmt, "2017-10-10")
	date3, err3 := time.Parse(dFmt, "2017-10-09")
	date4, err4 := time.Parse(dFmt, "2017-10-08")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date.")
	}

	if BusinessDaysBefore(date1, 1).Format(dFmt) != "2017-10-10" {
		t.Errorf(errFmt, "date1-2")
		t.Log("date1-1: ", BusinessDaysBefore(date1, 1).Format(dFmt))
	}
	if BusinessDaysBefore(date1, 2).Format(dFmt) != "2017-10-06" {
		t.Errorf(errFmt, "date1-2")
		t.Log("date1-2: ", BusinessDaysBefore(date1, 2).Format(dFmt))
	}
	if BusinessDaysBefore(date2, 1).Format(dFmt) != "2017-10-06" {
		t.Errorf(errFmt, "date2-1")
		t.Log("date2-1: ", BusinessDaysBefore(date2, 1).Format(dFmt))
	}
	if BusinessDaysBefore(date2, 2).Format(dFmt) != "2017-10-05" {
		t.Errorf(errFmt, "date2-2")
		t.Log("date2-2: ", BusinessDaysBefore(date2, 2).Format(dFmt))
	}
	if BusinessDaysBefore(date3, 1).Format(dFmt) != "2017-10-06" {
		t.Errorf(errFmt, "date3")
		t.Log("date3: ", BusinessDaysBefore(date3, 1).Format(dFmt))
	}
	if BusinessDaysBefore(date4, 1).Format(dFmt) != "2017-10-06" {
		t.Errorf(errFmt, "date4")
		t.Log("date4: ", BusinessDaysBefore(date4, 1).Format(dFmt))
	}
}

func BenchmarkBusinessDaysBefore(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BusinessDaysBefore(date[i], i)
	}
}

func TestBusinessDaysAfter(t *testing.T) {
	date1, err1 := time.Parse(dFmt, "2017-10-05")
	date2, err2 := time.Parse(dFmt, "2017-10-06")
	date3, err3 := time.Parse(dFmt, "2017-10-07")
	date4, err4 := time.Parse(dFmt, "2017-10-09")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date.")
	}

	if BusinessDaysAfter(date1, 1).Format(dFmt) != "2017-10-06" {
		t.Errorf(errFmt, "date1-1")
		t.Log("date1-1: ", BusinessDaysAfter(date1, 1).Format(dFmt))
	}
	if BusinessDaysAfter(date1, 2).Format(dFmt) != "2017-10-10" {
		t.Errorf(errFmt, "date1-2")
		t.Log("date1-2: ", BusinessDaysAfter(date1, 2).Format(dFmt))
	}
	if BusinessDaysAfter(date2, 1).Format(dFmt) != "2017-10-10" {
		t.Errorf(errFmt, "date2-1")
		t.Log("date2-1: ", BusinessDaysAfter(date2, 1).Format(dFmt))
	}
	if BusinessDaysAfter(date2, 2).Format(dFmt) != "2017-10-11" {
		t.Errorf(errFmt, "date2-2")
		t.Log("date2-2: ", BusinessDaysAfter(date2, 2).Format(dFmt))
	}
	if BusinessDaysAfter(date3, 1).Format(dFmt) != "2017-10-10" {
		t.Errorf(errFmt, "date3")
		t.Log("date3: ", BusinessDaysAfter(date3, 1).Format(dFmt))
	}
	if BusinessDaysAfter(date4, 1).Format(dFmt) != "2017-10-10" {
		t.Errorf(errFmt, "date4")
		t.Log("date4: ", BusinessDaysAfter(date4, 1).Format(dFmt))
	}
}

func BenchmarkBusinessDaysAfter(b *testing.B) {
	date := make([]time.Time, b.N)
	for n := 0; n < b.N; n++ {
		date[n] = time.Now().AddDate(0, 0, n)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BusinessDaysAfter(date[i], i)
	}
}

// set unique holidays
func TestBusinessDaysBeforeUnique(t *testing.T) {
	date1, err1 := time.Parse(dFmt, "2017-10-11")
	date2, err2 := time.Parse(dFmt, "2017-10-10")
	date3, err3 := time.Parse(dFmt, "2017-10-09")
	date4, err4 := time.Parse(dFmt, "2017-10-08")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date")
	}
	udate1, err1 := time.Parse(dFmt, "2017-10-10")
	udate2, err2 := time.Parse(dFmt, "2017-10-05")
	udate3, err3 := time.Parse(dFmt, "2017-10-06")
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error("Can`t parse date")
	}
	uhs := []time.Time{udate1, udate2, udate3}
	SetUniqueHolidays(uhs)

	if BusinessDaysBefore(date1, 1).Format(dFmt) != "2017-10-04" {
		t.Errorf(errFmt, "date1-2")
		t.Log("date1-1: ", BusinessDaysBefore(date1, 1).Format(dFmt))
	}
	if BusinessDaysBefore(date1, 2).Format(dFmt) != "2017-10-03" {
		t.Errorf(errFmt, "date1-2")
		t.Log("date1-2: ", BusinessDaysBefore(date1, 2).Format(dFmt))
	}
	if BusinessDaysBefore(date2, 1).Format(dFmt) != "2017-10-04" {
		t.Errorf(errFmt, "date2-1")
		t.Log("date2-1: ", BusinessDaysBefore(date2, 1).Format(dFmt))
	}
	if BusinessDaysBefore(date2, 2).Format(dFmt) != "2017-10-03" {
		t.Errorf(errFmt, "date2-2")
		t.Log("date2-2: ", BusinessDaysBefore(date2, 2).Format(dFmt))
	}
	if BusinessDaysBefore(date3, 1).Format(dFmt) != "2017-10-04" {
		t.Errorf(errFmt, "date3")
		t.Log("date3: ", BusinessDaysBefore(date3, 1).Format(dFmt))
	}
	if BusinessDaysBefore(date4, 1).Format(dFmt) != "2017-10-04" {
		t.Errorf(errFmt, "date4")
		t.Log("date4: ", BusinessDaysBefore(date4, 1).Format(dFmt))
	}
}

// set unique holidays
func TestBusinessDaysAfterUnique(t *testing.T) {
	date1, err1 := time.Parse(dFmt, "2017-10-05")
	date2, err2 := time.Parse(dFmt, "2017-10-06")
	date3, err3 := time.Parse(dFmt, "2017-10-07")
	date4, err4 := time.Parse(dFmt, "2017-10-09")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Error("Can`t parse date.")
	}
	udate1, err1 := time.Parse(dFmt, "2017-10-10")
	udate2, err2 := time.Parse(dFmt, "2017-10-12")
	udate3, err3 := time.Parse(dFmt, "2017-10-13")
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error("Can`t parse date")
	}
	uhs := []time.Time{udate1, udate2, udate3}
	SetUniqueHolidays(uhs)

	if BusinessDaysAfter(date1, 1).Format(dFmt) != "2017-10-06" {
		t.Errorf(errFmt, "date1-1")
		t.Log("date1-1: ", BusinessDaysAfter(date1, 1).Format(dFmt))
	}
	if BusinessDaysAfter(date1, 2).Format(dFmt) != "2017-10-11" {
		t.Errorf(errFmt, "date1-2")
		t.Log("date1-2: ", BusinessDaysAfter(date1, 2).Format(dFmt))
	}
	if BusinessDaysAfter(date2, 1).Format(dFmt) != "2017-10-11" {
		t.Errorf(errFmt, "date2-1")
		t.Log("date2-1: ", BusinessDaysAfter(date2, 1).Format(dFmt))
	}
	if BusinessDaysAfter(date2, 2).Format(dFmt) != "2017-10-16" {
		t.Errorf(errFmt, "date2-2")
		t.Log("date2-2: ", BusinessDaysAfter(date2, 2).Format(dFmt))
	}
	if BusinessDaysAfter(date3, 1).Format(dFmt) != "2017-10-11" {
		t.Errorf(errFmt, "date3")
		t.Log("date3: ", BusinessDaysAfter(date3, 1).Format(dFmt))
	}
	if BusinessDaysAfter(date4, 1).Format(dFmt) != "2017-10-11" {
		t.Errorf(errFmt, "date4")
		t.Log("date4: ", BusinessDaysAfter(date4, 1).Format(dFmt))
	}
}
