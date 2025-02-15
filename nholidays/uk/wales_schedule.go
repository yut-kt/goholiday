package uk

import (
	"time"

	"github.com/yut-kt/goholiday/nholidays"
)

func NewWales() *nholidays.ScheduleImpl {
	return nholidays.New(
		map[time.Weekday]struct{}{
			time.Saturday: {},
			time.Sunday:   {},
		},
		map[string]string{
			"2022-01-01": "New Year's Day",
			"2022-01-03": "New Year's Day (in lieu)",
			"2022-04-15": "Good Friday",
			"2022-04-18": "Easter Monday",
			"2022-05-02": "Early May Bank Holiday",
			"2022-06-02": "Spring Bank Holiday",
			"2022-06-03": "Queen's Platinum Jubilee Bank Holiday",
			"2022-08-29": "Summer Bank Holiday",
			"2022-12-25": "Christmas Day",
			"2022-12-26": "Boxing Day",
			"2022-12-27": "Christmas Day (in lieu)",
			"2023-01-01": "New Year's Day",
			"2023-01-02": "New Year's Day",
			"2023-04-07": "Good Friday",
			"2023-04-10": "Easter Monday",
			"2023-05-01": "Early May Bank Holiday",
			"2023-05-29": "Spring Bank Holiday",
			"2023-08-28": "Summer Bank Holiday",
			"2023-12-25": "Christmas Day",
			"2023-12-26": "Boxing Day",
		},
	)
}
