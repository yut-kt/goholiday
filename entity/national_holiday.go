package entity

import (
	"time"
<<<<<<< HEAD
	"github.com/yut-kt/goholiday/config"
=======
	"github.com/goholiday/config"
>>>>>>> [add]godoc
)

type NationalHoliday struct {
	Date string `yaml:"date"`
	Name string `yaml:"name"`
}

type NationalHolidays []*NationalHoliday

func (nh NationalHoliday) IsOrAfter(t time.Time) bool {
	pt, err := time.Parse(config.DateFormat, nh.Date)
	if err != nil {
		panic(err)
	}

	return nh.Date == t.Format(config.DateFormat) || pt.After(t)
}
