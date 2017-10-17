# goholiday

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
[![v0.1.2](https://img.shields.io/badge/package-v0.1.2-ff69b4.svg)](https://github.com/yut-kt/goholiday/tree/v0.1.2)
=======
[![v0.1.0](https://img.shields.io/github/package-json/v/badges/shields.svg)]()
>>>>>>> [update]readme
=======
[![v0.1.0](https://img.shields.io/badge/package-v0.1.0-ff69b4.svg)]()
>>>>>>> [update]version image
[![GoDoc](https://godoc.org/github.com/yut-kt/goholiday?status.svg)](https://godoc.org/github.com/yut-kt/goholiday)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/yut-kt/goholiday/v0.1.0/LICENSE)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
[![coverage](https://img.shields.io/badge/coverage-100%25-green.svg)](https://github.com/yut-kt/goholiday/coverage/v0.1.2)

**Functions to calculate and judge about business days in Japan.**
Now we are dealing with only Japanese holidays but we plan to deal with other national holidays.
=======
[![GoDoc](https://godoc.org/github.com/yut-kt/goholiday?status.svg)](https://godoc.org/github.com/yut-kt/goholiday)
>>>>>>> [add]godoc
=======
>>>>>>> [update]MIT License
=======
[![Go Report Card](https://goreportcard.com/badge/github.com/hybridgroup/gobot)](https://goreportcard.com/report/github.com/hybridgroup/gobot)
>>>>>>> [add]go report to readme
=======
>>>>>>> [fix]delete goreport from readme

**Functions to calculate and judge about business days in Japan.**

## Install
```bash
$ go get github.com/yut-kt/goholiday
<<<<<<< HEAD
```

## Import
```go
import (
  "github.com/yut-kt/goholiday"
)
=======
>>>>>>> [update]readme
```

## Import
```go
import (
  "github.com/yut-kt/goholiday"
)
```

## Usage

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
<<<<<<< HEAD
<<<<<<< HEAD
IsNationalHoliday is a function to decide whether t given national holiday
=======
=======
IsNationalHoliday is a function to decide whether t given national holiday
<<<<<<< HEAD
>>>>>>> [update]Add explanation to readme

## License
MIT

>>>>>>> [update]readme
=======
>>>>>>> [update]MIT License
