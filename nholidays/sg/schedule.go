package sg

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
			"2022-02-01": "Chinese New Year",
			"2022-02-02": "Chinese New Year Holiday",
			"2022-04-15": "Good Friday",
			"2022-05-02": "Labour Day (in lieu)",
			"2022-05-03": "Hari Raya Puasa",
			"2022-05-16": "Vesak Day (in lieu)",
			"2022-07-10": "Hari Raya Haji",
			"2022-07-11": "Hari Raya Haji (in lieu)",
			"2022-08-09": "National Day",
			"2022-10-24": "Deepavali",
			"2022-12-25": "Christmas Day",
			"2022-12-26": "Christmas Day (in lieu)",
			"2023-01-01": "New Year's Day",
			"2023-01-22": "Chinese New Year",
			"2023-01-23": "Chinese New Year",
			"2023-01-24": "Chinese New Year (in lieu)",
			"2023-04-07": "Good Friday",
			"2023-04-22": "Hari Raya Puasa",
			"2023-05-01": "Labour Day",
			"2023-06-03": "Vesak Day",
			"2023-06-29": "Hari Raya Haji",
			"2023-08-09": "National Day",
			"2023-11-12": "Deepavali",
			"2023-11-13": "Deepavali (in lieu)",
			"2023-12-25": "Christmas Day",
		},
	)
}
