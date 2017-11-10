# goholiday

[![v0.1.3](https://img.shields.io/badge/package-v0.1.3-ff69b4.svg)](https://github.com/yut-kt/goholiday/tree/v0.1.3)
[![GoDoc](https://godoc.org/github.com/yut-kt/goholiday?status.svg)](https://godoc.org/github.com/yut-kt/goholiday)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/yut-kt/goholiday/v0.1.3/LICENSE)
[![coverage](https://img.shields.io/badge/coverage-96%25-green.svg)](https://github.com/yut-kt/goholiday/coverage/v0.1.3)

**Functions to calculate and judge about business days in Japan.**
Now we are dealing with only Japanese holidays but we plan to deal with other national holidays.

## Install
```bash
$ go get github.com/yut-kt/goholiday
```

## Import
```go
import (
  "github.com/yut-kt/goholiday"
)
```

## Usage

### func SetUniqueHolidays
```go
func SetUniqueHolidays(ts []time.Time) void
```
SetUniqueHolidays is a function to set unique holidays

#### func  BusinessDaysAfter

```go
func BusinessDaysAfter(t time.Time, businessDays int) time.Time
```
BusinessDaysAfter is a function that calculates bds business days after given t


#### func  BusinessDaysBefore

```go
func BusinessDaysBefore(t time.Time, businessDays int) time.Time
```
BusinessDaysBefore is a function that calculates bds business days before given t

#### func  IsNationalHoliday

```go
func IsNationalHoliday(t time.Time) bool
```
IsBusinessDay is a function to decide whether t given business day

#### func  IsWeekDay

```go
func IsWeekDay(t time.Time) bool
```
IsNationalHoliday is a function to decide whether t given national holiday