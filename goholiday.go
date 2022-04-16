package goholiday

import "time"

const dateFormat = "2006-01-02"

type Schedule interface {
	GetNationalHolidays() map[string]string
	GetWeekdayHolidays() map[time.Weekday]struct{}
	GetLocation() *time.Location
}

type Goholiday struct {
	schedule       Schedule
	uniqueHolidays map[string]struct{}
}

func New(schedule Schedule) *Goholiday {
	return &Goholiday{
		schedule:       schedule,
		uniqueHolidays: map[string]struct{}{},
	}
}

func (g *Goholiday) IsNationalHoliday(t time.Time) bool {
	t = t.In(g.schedule.GetLocation())
	_, exist := g.schedule.GetNationalHolidays()[t.Format(dateFormat)]
	return exist
}

func (g *Goholiday) IsHoliday(t time.Time) bool {
	t = t.In(g.schedule.GetLocation())
	return g.isWeekdayHoliday(t) || g.IsNationalHoliday(t) || g.isUniqueHoliday(t)
}

func (g *Goholiday) isWeekdayHoliday(t time.Time) bool {
	_, exist := g.schedule.GetWeekdayHolidays()[t.Weekday()]
	return exist
}

func (g *Goholiday) SetUniqueHolidays(ts []time.Time) {
	for _, t := range ts {
		g.uniqueHolidays[t.Format(dateFormat)] = struct{}{}
	}
}

func (g *Goholiday) isUniqueHoliday(t time.Time) bool {
	_, exist := g.uniqueHolidays[t.Format(dateFormat)]
	return exist
}

func (g *Goholiday) IsBusinessDay(t time.Time) bool {
	t = t.In(g.schedule.GetLocation())
	return !g.IsHoliday(t)
}

func (g *Goholiday) BusinessDaysBefore(t time.Time, bds int) time.Time {
	return g.travelBusinessDays(t.In(g.schedule.GetLocation()), bds, -1)
}

func (g *Goholiday) BusinessDaysAfter(t time.Time, bds int) time.Time {
	return g.travelBusinessDays(t.In(g.schedule.GetLocation()), bds, 1)
}

func (g *Goholiday) travelBusinessDays(t time.Time, bds int, course int) time.Time {
	for tbds := 0; tbds != bds; {
		if t = t.AddDate(0, 0, course); !g.IsHoliday(t) {
			tbds++
		}
	}
	return t
}
