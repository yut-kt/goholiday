// Code generated by jp_generator.go; DO NOT EDIT.
// Based on holiday from https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv.

package jp

import "time"

type ScheduleImpl struct {
	nationalHolidays map[string]string
	weekdayHolidays  map[time.Weekday]struct{}
	location         *time.Location
}

func New() *ScheduleImpl {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	return &ScheduleImpl{
		nationalHolidays: map[string]string{
			"2000-01-01": "元日",
			"2000-01-10": "成人の日",
			"2000-02-11": "建国記念の日",
			"2000-03-20": "春分の日",
			"2000-04-29": "みどりの日",
			"2000-05-03": "憲法記念日",
			"2000-05-04": "休日",
			"2000-05-05": "こどもの日",
			"2000-07-20": "海の日",
			"2000-09-15": "敬老の日",
			"2000-09-23": "秋分の日",
			"2000-10-09": "体育の日",
			"2000-11-03": "文化の日",
			"2000-11-23": "勤労感謝の日",
			"2000-12-23": "天皇誕生日",
			"2001-01-01": "元日",
			"2001-01-08": "成人の日",
			"2001-02-11": "建国記念の日",
			"2001-02-12": "休日",
			"2001-03-20": "春分の日",
			"2001-04-29": "みどりの日",
			"2001-04-30": "休日",
			"2001-05-03": "憲法記念日",
			"2001-05-04": "休日",
			"2001-05-05": "こどもの日",
			"2001-07-20": "海の日",
			"2001-09-15": "敬老の日",
			"2001-09-23": "秋分の日",
			"2001-09-24": "休日",
			"2001-10-08": "体育の日",
			"2001-11-03": "文化の日",
			"2001-11-23": "勤労感謝の日",
			"2001-12-23": "天皇誕生日",
			"2001-12-24": "休日",
			"2002-01-01": "元日",
			"2002-01-14": "成人の日",
			"2002-02-11": "建国記念の日",
			"2002-03-21": "春分の日",
			"2002-04-29": "みどりの日",
			"2002-05-03": "憲法記念日",
			"2002-05-04": "休日",
			"2002-05-05": "こどもの日",
			"2002-05-06": "休日",
			"2002-07-20": "海の日",
			"2002-09-15": "敬老の日",
			"2002-09-16": "休日",
			"2002-09-23": "秋分の日",
			"2002-10-14": "体育の日",
			"2002-11-03": "文化の日",
			"2002-11-04": "休日",
			"2002-11-23": "勤労感謝の日",
			"2002-12-23": "天皇誕生日",
			"2003-01-01": "元日",
			"2003-01-13": "成人の日",
			"2003-02-11": "建国記念の日",
			"2003-03-21": "春分の日",
			"2003-04-29": "みどりの日",
			"2003-05-03": "憲法記念日",
			"2003-05-05": "こどもの日",
			"2003-07-21": "海の日",
			"2003-09-15": "敬老の日",
			"2003-09-23": "秋分の日",
			"2003-10-13": "体育の日",
			"2003-11-03": "文化の日",
			"2003-11-23": "勤労感謝の日",
			"2003-11-24": "休日",
			"2003-12-23": "天皇誕生日",
			"2004-01-01": "元日",
			"2004-01-12": "成人の日",
			"2004-02-11": "建国記念の日",
			"2004-03-20": "春分の日",
			"2004-04-29": "みどりの日",
			"2004-05-03": "憲法記念日",
			"2004-05-04": "休日",
			"2004-05-05": "こどもの日",
			"2004-07-19": "海の日",
			"2004-09-20": "敬老の日",
			"2004-09-23": "秋分の日",
			"2004-10-11": "体育の日",
			"2004-11-03": "文化の日",
			"2004-11-23": "勤労感謝の日",
			"2004-12-23": "天皇誕生日",
			"2005-01-01": "元日",
			"2005-01-10": "成人の日",
			"2005-02-11": "建国記念の日",
			"2005-03-20": "春分の日",
			"2005-03-21": "休日",
			"2005-04-29": "みどりの日",
			"2005-05-03": "憲法記念日",
			"2005-05-04": "休日",
			"2005-05-05": "こどもの日",
			"2005-07-18": "海の日",
			"2005-09-19": "敬老の日",
			"2005-09-23": "秋分の日",
			"2005-10-10": "体育の日",
			"2005-11-03": "文化の日",
			"2005-11-23": "勤労感謝の日",
			"2005-12-23": "天皇誕生日",
			"2006-01-01": "元日",
			"2006-01-02": "休日",
			"2006-01-09": "成人の日",
			"2006-02-11": "建国記念の日",
			"2006-03-21": "春分の日",
			"2006-04-29": "みどりの日",
			"2006-05-03": "憲法記念日",
			"2006-05-04": "休日",
			"2006-05-05": "こどもの日",
			"2006-07-17": "海の日",
			"2006-09-18": "敬老の日",
			"2006-09-23": "秋分の日",
			"2006-10-09": "体育の日",
			"2006-11-03": "文化の日",
			"2006-11-23": "勤労感謝の日",
			"2006-12-23": "天皇誕生日",
			"2007-01-01": "元日",
			"2007-01-08": "成人の日",
			"2007-02-11": "建国記念の日",
			"2007-02-12": "休日",
			"2007-03-21": "春分の日",
			"2007-04-29": "昭和の日",
			"2007-04-30": "休日",
			"2007-05-03": "憲法記念日",
			"2007-05-04": "みどりの日",
			"2007-05-05": "こどもの日",
			"2007-07-16": "海の日",
			"2007-09-17": "敬老の日",
			"2007-09-23": "秋分の日",
			"2007-09-24": "休日",
			"2007-10-08": "体育の日",
			"2007-11-03": "文化の日",
			"2007-11-23": "勤労感謝の日",
			"2007-12-23": "天皇誕生日",
			"2007-12-24": "休日",
			"2008-01-01": "元日",
			"2008-01-14": "成人の日",
			"2008-02-11": "建国記念の日",
			"2008-03-20": "春分の日",
			"2008-04-29": "昭和の日",
			"2008-05-03": "憲法記念日",
			"2008-05-04": "みどりの日",
			"2008-05-05": "こどもの日",
			"2008-05-06": "休日",
			"2008-07-21": "海の日",
			"2008-09-15": "敬老の日",
			"2008-09-23": "秋分の日",
			"2008-10-13": "体育の日",
			"2008-11-03": "文化の日",
			"2008-11-23": "勤労感謝の日",
			"2008-11-24": "休日",
			"2008-12-23": "天皇誕生日",
			"2009-01-01": "元日",
			"2009-01-12": "成人の日",
			"2009-02-11": "建国記念の日",
			"2009-03-20": "春分の日",
			"2009-04-29": "昭和の日",
			"2009-05-03": "憲法記念日",
			"2009-05-04": "みどりの日",
			"2009-05-05": "こどもの日",
			"2009-05-06": "休日",
			"2009-07-20": "海の日",
			"2009-09-21": "敬老の日",
			"2009-09-22": "休日",
			"2009-09-23": "秋分の日",
			"2009-10-12": "体育の日",
			"2009-11-03": "文化の日",
			"2009-11-23": "勤労感謝の日",
			"2009-12-23": "天皇誕生日",
			"2010-01-01": "元日",
			"2010-01-11": "成人の日",
			"2010-02-11": "建国記念の日",
			"2010-03-21": "春分の日",
			"2010-03-22": "休日",
			"2010-04-29": "昭和の日",
			"2010-05-03": "憲法記念日",
			"2010-05-04": "みどりの日",
			"2010-05-05": "こどもの日",
			"2010-07-19": "海の日",
			"2010-09-20": "敬老の日",
			"2010-09-23": "秋分の日",
			"2010-10-11": "体育の日",
			"2010-11-03": "文化の日",
			"2010-11-23": "勤労感謝の日",
			"2010-12-23": "天皇誕生日",
			"2011-01-01": "元日",
			"2011-01-10": "成人の日",
			"2011-02-11": "建国記念の日",
			"2011-03-21": "春分の日",
			"2011-04-29": "昭和の日",
			"2011-05-03": "憲法記念日",
			"2011-05-04": "みどりの日",
			"2011-05-05": "こどもの日",
			"2011-07-18": "海の日",
			"2011-09-19": "敬老の日",
			"2011-09-23": "秋分の日",
			"2011-10-10": "体育の日",
			"2011-11-03": "文化の日",
			"2011-11-23": "勤労感謝の日",
			"2011-12-23": "天皇誕生日",
			"2012-01-01": "元日",
			"2012-01-02": "休日",
			"2012-01-09": "成人の日",
			"2012-02-11": "建国記念の日",
			"2012-03-20": "春分の日",
			"2012-04-29": "昭和の日",
			"2012-04-30": "休日",
			"2012-05-03": "憲法記念日",
			"2012-05-04": "みどりの日",
			"2012-05-05": "こどもの日",
			"2012-07-16": "海の日",
			"2012-09-17": "敬老の日",
			"2012-09-22": "秋分の日",
			"2012-10-08": "体育の日",
			"2012-11-03": "文化の日",
			"2012-11-23": "勤労感謝の日",
			"2012-12-23": "天皇誕生日",
			"2012-12-24": "休日",
			"2013-01-01": "元日",
			"2013-01-14": "成人の日",
			"2013-02-11": "建国記念の日",
			"2013-03-20": "春分の日",
			"2013-04-29": "昭和の日",
			"2013-05-03": "憲法記念日",
			"2013-05-04": "みどりの日",
			"2013-05-05": "こどもの日",
			"2013-05-06": "休日",
			"2013-07-15": "海の日",
			"2013-09-16": "敬老の日",
			"2013-09-23": "秋分の日",
			"2013-10-14": "体育の日",
			"2013-11-03": "文化の日",
			"2013-11-04": "休日",
			"2013-11-23": "勤労感謝の日",
			"2013-12-23": "天皇誕生日",
			"2014-01-01": "元日",
			"2014-01-13": "成人の日",
			"2014-02-11": "建国記念の日",
			"2014-03-21": "春分の日",
			"2014-04-29": "昭和の日",
			"2014-05-03": "憲法記念日",
			"2014-05-04": "みどりの日",
			"2014-05-05": "こどもの日",
			"2014-05-06": "休日",
			"2014-07-21": "海の日",
			"2014-09-15": "敬老の日",
			"2014-09-23": "秋分の日",
			"2014-10-13": "体育の日",
			"2014-11-03": "文化の日",
			"2014-11-23": "勤労感謝の日",
			"2014-11-24": "休日",
			"2014-12-23": "天皇誕生日",
			"2015-01-01": "元日",
			"2015-01-12": "成人の日",
			"2015-02-11": "建国記念の日",
			"2015-03-21": "春分の日",
			"2015-04-29": "昭和の日",
			"2015-05-03": "憲法記念日",
			"2015-05-04": "みどりの日",
			"2015-05-05": "こどもの日",
			"2015-05-06": "休日",
			"2015-07-20": "海の日",
			"2015-09-21": "敬老の日",
			"2015-09-22": "休日",
			"2015-09-23": "秋分の日",
			"2015-10-12": "体育の日",
			"2015-11-03": "文化の日",
			"2015-11-23": "勤労感謝の日",
			"2015-12-23": "天皇誕生日",
			"2016-01-01": "元日",
			"2016-01-11": "成人の日",
			"2016-02-11": "建国記念の日",
			"2016-03-20": "春分の日",
			"2016-03-21": "休日",
			"2016-04-29": "昭和の日",
			"2016-05-03": "憲法記念日",
			"2016-05-04": "みどりの日",
			"2016-05-05": "こどもの日",
			"2016-07-18": "海の日",
			"2016-08-11": "山の日",
			"2016-09-19": "敬老の日",
			"2016-09-22": "秋分の日",
			"2016-10-10": "体育の日",
			"2016-11-03": "文化の日",
			"2016-11-23": "勤労感謝の日",
			"2016-12-23": "天皇誕生日",
			"2017-01-01": "元日",
			"2017-01-02": "休日",
			"2017-01-09": "成人の日",
			"2017-02-11": "建国記念の日",
			"2017-03-20": "春分の日",
			"2017-04-29": "昭和の日",
			"2017-05-03": "憲法記念日",
			"2017-05-04": "みどりの日",
			"2017-05-05": "こどもの日",
			"2017-07-17": "海の日",
			"2017-08-11": "山の日",
			"2017-09-18": "敬老の日",
			"2017-09-23": "秋分の日",
			"2017-10-09": "体育の日",
			"2017-11-03": "文化の日",
			"2017-11-23": "勤労感謝の日",
			"2017-12-23": "天皇誕生日",
			"2018-01-01": "元日",
			"2018-01-08": "成人の日",
			"2018-02-11": "建国記念の日",
			"2018-02-12": "休日",
			"2018-03-21": "春分の日",
			"2018-04-29": "昭和の日",
			"2018-04-30": "休日",
			"2018-05-03": "憲法記念日",
			"2018-05-04": "みどりの日",
			"2018-05-05": "こどもの日",
			"2018-07-16": "海の日",
			"2018-08-11": "山の日",
			"2018-09-17": "敬老の日",
			"2018-09-23": "秋分の日",
			"2018-09-24": "休日",
			"2018-10-08": "体育の日",
			"2018-11-03": "文化の日",
			"2018-11-23": "勤労感謝の日",
			"2018-12-23": "天皇誕生日",
			"2018-12-24": "休日",
			"2019-01-01": "元日",
			"2019-01-14": "成人の日",
			"2019-02-11": "建国記念の日",
			"2019-03-21": "春分の日",
			"2019-04-29": "昭和の日",
			"2019-04-30": "休日",
			"2019-05-01": "休日（祝日扱い）",
			"2019-05-02": "休日",
			"2019-05-03": "憲法記念日",
			"2019-05-04": "みどりの日",
			"2019-05-05": "こどもの日",
			"2019-05-06": "休日",
			"2019-07-15": "海の日",
			"2019-08-11": "山の日",
			"2019-08-12": "休日",
			"2019-09-16": "敬老の日",
			"2019-09-23": "秋分の日",
			"2019-10-14": "体育の日（スポーツの日）",
			"2019-10-22": "休日（祝日扱い）",
			"2019-11-03": "文化の日",
			"2019-11-04": "休日",
			"2019-11-23": "勤労感謝の日",
			"2020-01-01": "元日",
			"2020-01-13": "成人の日",
			"2020-02-11": "建国記念の日",
			"2020-02-23": "天皇誕生日",
			"2020-02-24": "休日",
			"2020-03-20": "春分の日",
			"2020-04-29": "昭和の日",
			"2020-05-03": "憲法記念日",
			"2020-05-04": "みどりの日",
			"2020-05-05": "こどもの日",
			"2020-05-06": "休日",
			"2020-07-23": "海の日",
			"2020-07-24": "スポーツの日",
			"2020-08-10": "山の日",
			"2020-09-21": "敬老の日",
			"2020-09-22": "秋分の日",
			"2020-11-03": "文化の日",
			"2020-11-23": "勤労感謝の日",
			"2021-01-01": "元日",
			"2021-01-11": "成人の日",
			"2021-02-11": "建国記念の日",
			"2021-02-23": "天皇誕生日",
			"2021-03-20": "春分の日",
			"2021-04-29": "昭和の日",
			"2021-05-03": "憲法記念日",
			"2021-05-04": "みどりの日",
			"2021-05-05": "こどもの日",
			"2021-07-22": "海の日",
			"2021-07-23": "スポーツの日",
			"2021-08-08": "山の日",
			"2021-08-09": "休日",
			"2021-09-20": "敬老の日",
			"2021-09-23": "秋分の日",
			"2021-11-03": "文化の日",
			"2021-11-23": "勤労感謝の日",
			"2022-01-01": "元日",
			"2022-01-10": "成人の日",
			"2022-02-11": "建国記念の日",
			"2022-02-23": "天皇誕生日",
			"2022-03-21": "春分の日",
			"2022-04-29": "昭和の日",
			"2022-05-03": "憲法記念日",
			"2022-05-04": "みどりの日",
			"2022-05-05": "こどもの日",
			"2022-07-18": "海の日",
			"2022-08-11": "山の日",
			"2022-09-19": "敬老の日",
			"2022-09-23": "秋分の日",
			"2022-10-10": "スポーツの日",
			"2022-11-03": "文化の日",
			"2022-11-23": "勤労感謝の日",
			"2023-01-01": "元日",
			"2023-01-02": "休日",
			"2023-01-09": "成人の日",
			"2023-02-11": "建国記念の日",
			"2023-02-23": "天皇誕生日",
			"2023-03-21": "春分の日",
			"2023-04-29": "昭和の日",
			"2023-05-03": "憲法記念日",
			"2023-05-04": "みどりの日",
			"2023-05-05": "こどもの日",
			"2023-07-17": "海の日",
			"2023-08-11": "山の日",
			"2023-09-18": "敬老の日",
			"2023-09-23": "秋分の日",
			"2023-10-09": "スポーツの日",
			"2023-11-03": "文化の日",
			"2023-11-23": "勤労感謝の日",
			"2024-01-01": "元日",
			"2024-01-08": "成人の日",
			"2024-02-11": "建国記念の日",
			"2024-02-12": "休日",
			"2024-02-23": "天皇誕生日",
			"2024-03-20": "春分の日",
			"2024-04-29": "昭和の日",
			"2024-05-03": "憲法記念日",
			"2024-05-04": "みどりの日",
			"2024-05-05": "こどもの日",
			"2024-05-06": "休日",
			"2024-07-15": "海の日",
			"2024-08-11": "山の日",
			"2024-08-12": "休日",
			"2024-09-16": "敬老の日",
			"2024-09-22": "秋分の日",
			"2024-09-23": "休日",
			"2024-10-14": "スポーツの日",
			"2024-11-03": "文化の日",
			"2024-11-04": "休日",
			"2024-11-23": "勤労感謝の日",
			"2025-01-01": "元日",
			"2025-01-13": "成人の日",
			"2025-02-11": "建国記念の日",
			"2025-02-23": "天皇誕生日",
			"2025-02-24": "休日",
			"2025-03-20": "春分の日",
			"2025-04-29": "昭和の日",
			"2025-05-03": "憲法記念日",
			"2025-05-04": "みどりの日",
			"2025-05-05": "こどもの日",
			"2025-05-06": "休日",
			"2025-07-21": "海の日",
			"2025-08-11": "山の日",
			"2025-09-15": "敬老の日",
			"2025-09-23": "秋分の日",
			"2025-10-13": "スポーツの日",
			"2025-11-03": "文化の日",
			"2025-11-23": "勤労感謝の日",
			"2025-11-24": "休日",
			"2026-01-01": "元日",
			"2026-01-12": "成人の日",
			"2026-02-11": "建国記念の日",
			"2026-02-23": "天皇誕生日",
			"2026-03-20": "春分の日",
			"2026-04-29": "昭和の日",
			"2026-05-03": "憲法記念日",
			"2026-05-04": "みどりの日",
			"2026-05-05": "こどもの日",
			"2026-05-06": "休日",
			"2026-07-20": "海の日",
			"2026-08-11": "山の日",
			"2026-09-21": "敬老の日",
			"2026-09-22": "休日",
			"2026-09-23": "秋分の日",
			"2026-10-12": "スポーツの日",
			"2026-11-03": "文化の日",
			"2026-11-23": "勤労感謝の日",
		},
		weekdayHolidays: map[time.Weekday]struct{}{
			time.Saturday: {},
			time.Sunday:   {},
		},
		location: loc,
	}
}

func (s *ScheduleImpl) GetNationalHolidays() map[string]string {
	return s.nationalHolidays
}

func (s *ScheduleImpl) GetWeekdayHolidays() map[time.Weekday]struct{} {
	return s.weekdayHolidays
}

func (s *ScheduleImpl) GetLocation() *time.Location {
	return s.location
}
