package do

import (
	"time"

	"github.com/yut-kt/goholiday/nholidays"
)

func New() *nholidays.ScheduleImpl {
	return nholidays.New(
		map[time.Weekday]struct{}{
			time.Saturday: {},
			time.Sunday:   {},
		},
		map[string]string{
			"2022-01-01": "New Year's Day",
			"2022-01-10": "Epiphany Holiday",
			"2022-01-21": "Lady of Altagracia Day",
			"2022-01-24": "Juan Pablo Duarte Day Holiday",
			"2022-02-27": "Independence Day",
			"2022-04-14": "Maundy Thursday",
			"2022-04-15": "Good Friday",
			"2022-05-01": "Labor Day",
			"2022-05-02": "Labor Day Holiday",
			"2022-06-16": "Corpus Christi",
			"2022-08-16": "Restoration Day",
			"2022-09-24": "Lady of Mercedes Day",
			"2022-11-06": "Constitution Day",
			"2022-12-25": "Christmas Day",
			"2023-01-01": "New Year's Day",
			"2023-01-09": "Epiphany Holiday",
			"2023-01-21": "Lady of Altagracia Day",
			"2023-01-30": "Juan Pablo Duarte Day Holiday",
			"2023-02-27": "Independence Day",
			"2023-04-06": "Maundy Thursday",
			"2023-04-07": "Good Friday",
			"2023-05-01": "Labor Day",
			"2023-06-08": "Corpus Christi",
		},
	)
}
