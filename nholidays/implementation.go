package nholidays

import "time"

type ScheduleImpl struct {
	location         *time.Location
	weekdayHolidays  map[time.Weekday]struct{}
	nationalHolidays map[string]string
}

func (s *ScheduleImpl) GetLocation() *time.Location {
	return s.location
}

func (s *ScheduleImpl) GetWeekdayHolidays() map[time.Weekday]struct{} {
	return s.weekdayHolidays
}

func (s *ScheduleImpl) GetNationalHolidays() map[string]string {
	return s.nationalHolidays
}

func New(l *time.Location, whs map[time.Weekday]struct{}, nhs map[string]string) *ScheduleImpl {
	return &ScheduleImpl{
		location:         l,
		weekdayHolidays:  whs,
		nationalHolidays: nhs,
	}
}
