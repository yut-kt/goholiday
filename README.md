# goholiday

[![v0.1.5](https://img.shields.io/badge/package-v0.1.5-ff69b4.svg)]()
[![GoDoc](https://godoc.org/github.com/yut-kt/goholiday?status.svg)](https://godoc.org/github.com/yut-kt/goholiday)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/yut-kt/goholiday/master/LICENSE)
[![coverage](https://img.shields.io/badge/coverage-96%25-green.svg)](https://raw.githubusercontent.com/yut-kt/goholiday/master/coverage/v0.1.5)
[![Go Report Card](https://goreportcard.com/badge/github.com/yut-kt/goholiday)](https://goreportcard.com/report/github.com/yut-kt/goholiday)

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

#### func  IsNationalHoliday

```go
func IsNationalHoliday(t time.Time) bool
```
IsNationalHoliday is a function to decide whether t given national holiday

#### func  IsWeekDay

```go
func IsBusinessday(t time.Time) bool
```
IsBusinessday is a function to decide whether t given businessday

#### func  BusinessDaysBefore

```go
func BusinessDaysBefore(t time.Time, businessDays int) time.Time
```
BusinessDaysBefore is a function that calculates bds business days before given t

#### func  BusinessDaysAfter

```go
func BusinessDaysAfter(t time.Time, businessDays int) time.Time
```
BusinessDaysAfter is a function that calculates bds business days after given t


## Contribution

1. Fork ([https://github.com/yut-kt/goholiday/fork](https://github.com/yut-kt/goholiday/fork))
2. Checkout the latest version of branch
3. Create a feature branch
4. Commit your changes
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s` or `goimports -s`
7. Create new Pull Request