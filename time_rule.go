// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rdate

import "time"

type TimeRule interface {
	Calculate(pivot time.Time) time.Time
	Shortcut() TimeShortcut
}

var (
	thisDayStartRule = &timeRuleStartOfThisDay{}
	thisDayEndRule   = &timeRuleEndOfThisDay{}
)

type timeRuleAsIs struct{}

func (r *timeRuleAsIs) Calculate(pivot time.Time) time.Time {
	return pivot
}

func (r *timeRuleAsIs) Shortcut() TimeShortcut { return TimeAsIs }

type timeRuleStartOfThisDay struct{}

func (r *timeRuleStartOfThisDay) Calculate(pivot time.Time) time.Time {
	return time.Date(pivot.Year(), pivot.Month(), pivot.Day(), 0, 0, 0, 0,
		pivot.Location())
}

func (r *timeRuleStartOfThisDay) Shortcut() TimeShortcut { return TimeStartOfThisDay }

type timeRuleEndOfThisDay struct{}

func (r *timeRuleEndOfThisDay) Calculate(pivot time.Time) time.Time {
	return time.Date(pivot.Year(), pivot.Month(), pivot.Day(), 23, 59, 59, 999999999,
		pivot.Location())
}

func (r *timeRuleEndOfThisDay) Shortcut() TimeShortcut { return TimeEndOfThisDay }

type timeRuleStartOfPrevDay struct{}

func (r *timeRuleStartOfPrevDay) Calculate(pivot time.Time) time.Time {
	return thisDayStartRule.Calculate(pivot).AddDate(0, 0, -1)
}

func (r *timeRuleStartOfPrevDay) Shortcut() TimeShortcut { return TimeStartOfPrevDay }

type timeRuleEndOfPrevDay struct{}

func (r *timeRuleEndOfPrevDay) Calculate(pivot time.Time) time.Time {
	return thisDayEndRule.Calculate(pivot).AddDate(0, 0, -1)
}

func (r *timeRuleEndOfPrevDay) Shortcut() TimeShortcut { return TimeEndOfPrevDay }

type timeRuleStartOfThisWeek struct{}

func (r *timeRuleStartOfThisWeek) Calculate(pivot time.Time) time.Time {
	ts := thisDayStartRule.Calculate(pivot)
	if ts.Weekday() == time.Sunday {
		return ts.AddDate(0, 0, -6)
	}
	return ts.AddDate(0, 0, -(int(ts.Weekday()) - 1))
}

func (r *timeRuleStartOfThisWeek) Shortcut() TimeShortcut {
	return TimeStartOfThisWeek
}

type timeRuleEndOfThisWeek struct{}

func (r *timeRuleEndOfThisWeek) Calculate(pivot time.Time) time.Time {
	ts := thisDayEndRule.Calculate(pivot)
	if ts.Weekday() == time.Sunday {
		return ts
	}
	return ts.AddDate(0, 0, 7-int(ts.Weekday()))
}

func (r *timeRuleEndOfThisWeek) Shortcut() TimeShortcut {
	return TimeEndOfThisWeek
}

type timeRuleStartOfPrevWeek struct{}

func (r *timeRuleStartOfPrevWeek) Calculate(pivot time.Time) time.Time {
	ts := thisDayStartRule.Calculate(pivot).AddDate(0, 0, -7)
	if ts.Weekday() == time.Sunday {
		return ts.AddDate(0, 0, -6)
	}

	return ts.AddDate(0, 0, -(int(ts.Weekday()) - 1))
}

func (r *timeRuleStartOfPrevWeek) Shortcut() TimeShortcut { return TimeStartOfPrevWeek }

type timeRuleEndOfPrevWeek struct{}

func (r *timeRuleEndOfPrevWeek) Calculate(pivot time.Time) time.Time {
	ts := thisDayEndRule.Calculate(pivot)
	if ts.Weekday() == time.Sunday {
		ts = ts.AddDate(0, 0, -6)
	} else {
		ts = ts.AddDate(0, 0, -(int(ts.Weekday()) - 1))
	}
	return ts.AddDate(0, 0, -1)
}

func (r *timeRuleEndOfPrevWeek) Shortcut() TimeShortcut { return TimeEndOfPrevWeek }

type timeRuleStartOfThisMonth struct{}

func (r *timeRuleStartOfThisMonth) Calculate(pivot time.Time) time.Time {
	return time.Date(pivot.Year(), pivot.Month(), 1, 0, 0, 0, 0, pivot.Location())
}

func (r *timeRuleStartOfThisMonth) Shortcut() TimeShortcut {
	return TimeStartOfThisMonth
}

type timeRuleEndOfThisMonth struct{}

func (r *timeRuleEndOfThisMonth) Calculate(pivot time.Time) time.Time {
	return time.Date(pivot.Year(), pivot.Month(), 1, 23, 59, 59, 999999999, pivot.Location()).
		AddDate(0, 1, -1)
}

func (r *timeRuleEndOfThisMonth) Shortcut() TimeShortcut { return TimeEndOfThisMonth }

type timeRuleStartOfPrevMonth struct{}

func (r *timeRuleStartOfPrevMonth) Calculate(pivot time.Time) time.Time {
	ts := pivot.AddDate(0, -1, 0)
	return time.Date(ts.Year(), ts.Month(), 1, 0, 0, 0, 0, pivot.Location())
}

func (r *timeRuleStartOfPrevMonth) Shortcut() TimeShortcut {
	return TimeStartOfPrevMonth
}

type timeRuleEndOfPrevMonth struct{}

func (r *timeRuleEndOfPrevMonth) Calculate(pivot time.Time) time.Time {
	return time.Date(pivot.Year(), pivot.Month(), 1, 23, 59, 59, 999999999, pivot.Location()).
		AddDate(0, 0, -1)
}

func (r *timeRuleEndOfPrevMonth) Shortcut() TimeShortcut { return TimeEndOfPrevMonth }

type timeRuleStartOfThisQuart struct{}

func (r *timeRuleStartOfThisQuart) Calculate(pivot time.Time) time.Time {
	quartNum := ((pivot.Month() - 1) / 3)
	month := quartNum*3 + 1
	return time.Date(pivot.Year(), month, 1, 0, 0, 0, 0, pivot.Location())
}

func (r *timeRuleStartOfThisQuart) Shortcut() TimeShortcut {
	return TimeStartOfThisQuart
}

type timeRuleEndOfThisQuart struct{}

func (r *timeRuleEndOfThisQuart) Calculate(pivot time.Time) time.Time {
	quartNum := int((pivot.Month() - 1) / 3)
	offset := quartNum*3 + 3
	return time.Date(pivot.Year(), 1, 1, 23, 59, 59, 999999999, pivot.Location()).
		AddDate(0, offset, -1)
}

func (r *timeRuleEndOfThisQuart) Shortcut() TimeShortcut { return TimeEndOfThisQuart }

type timeRuleStartOfPrevQuart struct{}

func (r *timeRuleStartOfPrevQuart) Calculate(pivot time.Time) time.Time {
	ts := pivot.AddDate(0, -3, 0)
	quartNum := ((ts.Month() - 1) / 3)
	month := quartNum*3 + 1
	return time.Date(ts.Year(), month, 1, 0, 0, 0, 0, pivot.Location())
}

func (r *timeRuleStartOfPrevQuart) Shortcut() TimeShortcut {
	return TimeStartOfPrevQuart
}

type timeRuleEndOfPrevQuart struct{}

func (r *timeRuleEndOfPrevQuart) Calculate(pivot time.Time) time.Time {
	quartNum := ((pivot.Month() - 1) / 3)
	month := quartNum*3 + 1
	return time.Date(pivot.Year(), month, 1, 23, 59, 59, 999999999, pivot.Location()).
		AddDate(0, 0, -1)
}

func (r *timeRuleEndOfPrevQuart) Shortcut() TimeShortcut { return TimeEndOfPrevQuart }

type timeRuleStartOfThisHalfYear struct{}

func (r *timeRuleStartOfThisHalfYear) Calculate(pivot time.Time) time.Time {
	halfNum := ((pivot.Month() - 1) / 6)
	month := halfNum*6 + 1
	return time.Date(pivot.Year(), month, 1, 0, 0, 0, 0, pivot.Location())
}

func (r *timeRuleStartOfThisHalfYear) Shortcut() TimeShortcut {
	return TimeStartOfThisHalfYear
}

type timeRuleEndOfThisHalfYear struct{}

func (r *timeRuleEndOfThisHalfYear) Calculate(pivot time.Time) time.Time {
	halfNum := int(((pivot.Month() - 1) / 6))
	offset := halfNum*6 + 6

	return time.Date(pivot.Year(), 1, 1, 23, 59, 59, 999999999, pivot.Location()).
		AddDate(0, offset, -1)
}

func (r *timeRuleEndOfThisHalfYear) Shortcut() TimeShortcut {
	return TimeEndOfThisHalfYear
}

type timeRuleStartOfPrevHalfYear struct{}

func (r *timeRuleStartOfPrevHalfYear) Calculate(pivot time.Time) time.Time {
	ts := pivot.AddDate(0, -6, 0)
	halfNum := ((ts.Month() - 1) / 6)
	month := halfNum*6 + 1
	return time.Date(ts.Year(), month, 1, 0, 0, 0, 0, pivot.Location())
}

func (r *timeRuleStartOfPrevHalfYear) Shortcut() TimeShortcut {
	return TimeStartOfPrevHalfYear
}

type timeRuleEndOfPrevHalfYear struct{}

func (r *timeRuleEndOfPrevHalfYear) Calculate(pivot time.Time) time.Time {
	halfNum := ((pivot.Month() - 1) / 6)
	month := halfNum*6 + 1
	return time.Date(pivot.Year(), month, 1, 23, 59, 59, 999999999, pivot.Location()).
		AddDate(0, 0, -1)
}

func (r *timeRuleEndOfPrevHalfYear) Shortcut() TimeShortcut {
	return TimeEndOfPrevHalfYear
}

type timeRuleStartOfThisYear struct{}

func (r *timeRuleStartOfThisYear) Calculate(pivot time.Time) time.Time {
	return time.Date(pivot.Year(), 1, 1, 0, 0, 0, 0, pivot.Location())
}

func (r *timeRuleStartOfThisYear) Shortcut() TimeShortcut {
	return TimeStartOfThisYear
}

type timeRuleEndOfThisYear struct{}

func (r *timeRuleEndOfThisYear) Calculate(pivot time.Time) time.Time {
	return time.Date(pivot.Year(), 1, 1, 23, 59, 59, 999999999, pivot.Location()).
		AddDate(1, 0, -1)
}

func (r *timeRuleEndOfThisYear) Shortcut() TimeShortcut { return TimeEndOfThisYear }

type timeRuleStartOfPrevYear struct{}

func (r *timeRuleStartOfPrevYear) Calculate(pivot time.Time) time.Time {
	return time.Date(pivot.AddDate(-1, 0, 0).Year(), 1, 1, 0, 0, 0, 0, pivot.Location())
}

func (r *timeRuleStartOfPrevYear) Shortcut() TimeShortcut { return TimeStartOfPrevYear }

type timeRuleEndOfPrevYear struct{}

func (r *timeRuleEndOfPrevYear) Calculate(pivot time.Time) time.Time {
	return time.Date(pivot.Year(), 1, 1, 23, 59, 59, 999999999, pivot.Location()).
		AddDate(0, 0, -1)
}

func (r *timeRuleEndOfPrevYear) Shortcut() TimeShortcut { return TimeEndOfPrevYear }

type timeRuleStartOfThisWeekS struct{}

func (r *timeRuleStartOfThisWeekS) Calculate(pivot time.Time) time.Time {
	ts := thisDayStartRule.Calculate(pivot)
	return ts.AddDate(0, 0, -(int(ts.Weekday())))
}

func (r *timeRuleStartOfThisWeekS) Shortcut() TimeShortcut {
	return TimeStartOfThisWeek
}

type timeRuleEndOfThisWeekS struct{}

func (r *timeRuleEndOfThisWeekS) Calculate(pivot time.Time) time.Time {
	ts := thisDayEndRule.Calculate(pivot)
	return ts.AddDate(0, 0, 6-int(ts.Weekday()))
}

func (r *timeRuleEndOfThisWeekS) Shortcut() TimeShortcut {
	return TimeEndOfThisWeek
}

type timeRuleStartOfPrevWeekS struct{}

func (r *timeRuleStartOfPrevWeekS) Calculate(pivot time.Time) time.Time {
	ts := thisDayStartRule.Calculate(pivot).AddDate(0, 0, -7)
	return ts.AddDate(0, 0, -(int(ts.Weekday())))
}

func (r *timeRuleStartOfPrevWeekS) Shortcut() TimeShortcut { return TimeStartOfPrevWeek }

type timeRuleEndOfPrevWeekS struct{}

func (r *timeRuleEndOfPrevWeekS) Calculate(pivot time.Time) time.Time {
	ts := thisDayEndRule.Calculate(pivot)
	return ts.AddDate(0, 0, -(int(ts.Weekday()))-1)
}

func (r *timeRuleEndOfPrevWeekS) Shortcut() TimeShortcut { return TimeEndOfPrevWeek }
