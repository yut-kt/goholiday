package nholidays

import "time"

type ScheduleImpl struct {
	weekdayHolidays  map[time.Weekday]struct{}
	nationalHolidays map[string]string
}

func (s *ScheduleImpl) GetWeekdayHolidays() map[time.Weekday]struct{} {
	return s.weekdayHolidays
}

func (s *ScheduleImpl) GetNationalHolidays() map[string]string {
	return s.nationalHolidays
}

func New(whs map[time.Weekday]struct{}, nhs map[string]string) *ScheduleImpl {
	return &ScheduleImpl{
		weekdayHolidays:  whs,
		nationalHolidays: nhs,
	}
}
