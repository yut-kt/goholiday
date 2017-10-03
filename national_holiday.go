package main

import (
	"time"
	"github.com/holiday-jp/config"
	"sort"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type NationalHoliday struct {
	Date string `yaml:"date"`
	Name string `yaml:"name"`
}

type NationalHolidays []*NationalHoliday

func (nh NationalHoliday) isOrLater(t time.Time) bool {
	pt, err := time.Parse(config.DateFormat, nh.Date)
	if err != nil {
		panic(err)
	}

	return pt.After(t) || pt.Equal(t)
}

func isNationalHoliday(t time.Time) (bool, error) {
	nationalHolidays, err := fetchNationalHolidays()
	if err != nil {
		return false, err
	}

	index := sort.Search(len(nationalHolidays), func(i int) bool { return nationalHolidays[i].isOrLater(t) })
	return nationalHolidays[index].Date == t.Format(config.DateFormat), nil
}

func fetchNationalHolidays() (NationalHolidays, error) {
	buf, err := ioutil.ReadFile("data/national_holidays.yaml")
	if err != nil {
		return nil, err
	}

	var nh NationalHolidays
	err = yaml.Unmarshal(buf, &nh)
	if err != nil {
		return nil, err
	}

	return nh, nil
}
