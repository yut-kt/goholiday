package entity

import (
	"time"
	"github.com/yut-kt/goholiday/config"
)

type NationalHoliday struct {
	Date string `yaml:"date"`
	Name string `yaml:"name"`
}

type NationalHolidays []*NationalHoliday

var DFmt = config.DateFormat

func (nh NationalHoliday) IsOrAfter(t time.Time) bool {
	pt, err := time.Parse(DFmt, nh.Date)
	if err != nil {
		panic(err)
	}

	return nh.Date == t.Format(DFmt) || pt.After(t)
}