# goholiday

[![v0.1.6](https://img.shields.io/badge/package-v0.1.6-ff69b4.svg)]()
[![GoDoc](https://godoc.org/github.com/kenzo0107/goholiday?status.svg)](https://godoc.org/github.com/kenzo0107/goholiday)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/yut-kt/goholiday/master/LICENSE)
[![coverage](https://img.shields.io/badge/coverage-100%25-green.svg)](https://github.com/kenzo0107/goholiday/coverage/v0.1.6)
[![Go Report Card](https://goreportcard.com/badge/github.com/kenzo0107/goholiday)](https://goreportcard.com/report/github.com/kenzo0107/goholiday)

**Functions to calculate and judge about business days in Japan.**
Now we are dealing with only Japanese holidays but we plan to deal with other national holidays.

## Install
```bash
$ go get github.com/kenzo0107/goholiday
```

## Import
```go
import (
  "github.com/kenzo0107/goholiday"
)
```

## Usage

#### func  IsNationalHoliday
```go
func IsNationalHoliday(t time.Time) bool
```
IsNationalHoliday is a function to decide whether t given national holiday

### func SetUniqueHolidays
```go
func SetUniqueHolidays(ts []time.Time) void
```
SetUniqueHolidays is a function to set unique holidays

#### func IsHoliday
```go
func IsHoliday(t time.Time) bool
```
IsHoliday is a function to decide whether t given holiday

#### func  IsBusinessDay
```go
func IsBusinessDay(t time.Time) bool
```
IsBusinessDay is a function to decide whether t given business day

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

1. Fork ([https://github.com/kenzo0107/goholiday/fork](https://github.com/kenzo0107/goholiday/fork))
2. Checkout the latest version of branch
3. Create a feature branch
4. Commit your changes
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s` or `goimports -s`
7. Create new Pull Request

## License
goholiday is released under the [MIT License](https://raw.githubusercontent.com/yut-kt/goholiday/master/LICENSE).
