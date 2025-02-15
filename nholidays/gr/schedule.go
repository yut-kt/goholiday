package gr

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
			"2022-01-06": "Epiphany",
			"2022-01-25": "Public Holiday",
			"2022-01-26": "Public Holiday",
			"2022-03-07": "Ash Monday / Clean Monday",
			"2022-03-25": "Independence Day",
			"2022-04-22": "Orthodox Good Friday",
			"2022-04-24": "Orthodox Easter",
			"2022-04-25": "Orthodox Easter Monday",
			"2022-05-02": "May Day Holiday",
			"2022-06-12": "Orthodox Whit Sunday",
			"2022-06-13": "Orthodox Whit Monday",
			"2022-08-15": "Assumption Day",
			"2022-10-28": "Ochi Day",
			"2022-12-25": "Christmas Day",
			"2022-12-26": "Synaxis of the Mother of God",
			"2023-01-01": "New Year's Day",
			"2023-01-06": "Epiphany",
			"2023-02-27": "Ash Monday / Clean Monday",
			"2023-03-25": "Independence Day",
			"2023-04-14": "Orthodox Good Friday",
			"2023-04-16": "Orthodox Easter",
			"2023-04-17": "Orthodox Easter Monday",
			"2023-05-01": "May Day",
			"2023-06-04": "Orthodox Whit Sunday",
			"2023-06-05": "Orthodox Whit Monday",
			"2023-08-15": "Assumption Day",
			"2023-10-28": "Ochi Day",
			"2023-12-25": "Christmas Day",
			"2023-12-26": "Synaxis of the Mother of God",
		},
	)
}
