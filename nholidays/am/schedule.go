package am

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
			"2022-01-02": "New Year Holiday",
			"2022-01-06": "Armenian Christmas",
			"2022-01-28": "National Army Day",
			"2022-03-08": "International Women's Day",
			"2022-04-24": "Armenian Remembrance Day",
			"2022-05-01": "Labour Day",
			"2022-05-09": "Victory and Peace Day",
			"2022-05-28": "1st Republic Day",
			"2022-07-05": "Constitution Day",
			"2022-09-21": "Independence Day",
			"2022-12-31": "New Year's Eve",
			"2023-01-01": "New Year's Day",
			"2023-01-02": "New Year Holiday",
			"2023-01-06": "Armenian Christmas",
			"2023-01-28": "National Army Day",
			"2023-03-08": "International Women's Day",
			"2023-04-24": "Armenian Remembrance Day",
			"2023-05-01": "Labour Day",
			"2023-05-09": "Victory and Peace Day",
			"2023-05-28": "1st Republic Day",
		},
	)
}
