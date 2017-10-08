# goholiday

## Install
```bash
$ go get github.com/yut-kt/holiday-jp
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

