# goholiday

[![v0.1.0](https://img.shields.io/badge/package-v0.1.0-ff69b4.svg)]()
[![GoDoc](https://godoc.org/github.com/yut-kt/goholiday?status.svg)](https://godoc.org/github.com/yut-kt/goholiday)

## Install
```bash
$ go get github.com/yut-kt/goholiday
```

## Usage

#### func  BusinessDaysAfter

```go
func BusinessDaysAfter(t time.Time, businessDays int) time.Time
```

#### func  BusinessDaysBefore

```go
func BusinessDaysBefore(t time.Time, businessDays int) time.Time
```

#### func  IsNationalHoliday

```go
func IsNationalHoliday(t time.Time) bool
```

#### func  IsWeekDay

```go
func IsWeekDay(t time.Time) bool
```

## License
MIT

