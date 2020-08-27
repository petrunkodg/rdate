package rdate_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/petrunkodg/rdate"
)

const format = "January 2, 2006"

type customPeriodStringer struct{}

func (s *customPeriodStringer) String(from, to rdate.Time, sc rdate.PeriodShortcut) string {
	p := ""

	switch sc {
	case rdate.PeriodThisDay:
		p = "this day"
	case rdate.PeriodThisWeek:
		p = "this week"
	case rdate.PeriodThisMonth:
		p = "this month"
	case rdate.PeriodThisQuart:
		p = "this quarter"
	case rdate.PeriodThisHalfYear:
		p = "this half year"
	case rdate.PeriodThisYear:
		p = "this year"
	case rdate.PeriodPrevDay:
		p = "previous day"
	case rdate.PeriodPrevWeek:
		p = "previous week"
	case rdate.PeriodPrevMonth:
		p = "previous month"
	case rdate.PeriodPrevQuart:
		p = "previous quarter"
	case rdate.PeriodPrevHalfYear:
		p = "previous half year"
	case rdate.PeriodPrevYear:
		p = "previous year"
	default:
	}

	var b strings.Builder

	if len(p) > 0 {
		b.WriteString(p)
		b.WriteString(" is ")
	}

	b.WriteString(fmt.Sprintf("(%s — %s)", from.Time().Format(format),
		to.Time().Format(format)))

	return b.String()
}

func ExamplePeriodStringer_replacing() {
	pf := rdate.NewPeriodFactory()

	pf.SetStringer(&customPeriodStringer{})

	ts := time.Date(2020, 8, 11, 0, 2, 1, 6, time.UTC)

	fmt.Println("today is", ts.Format(format))
	// today is August 11, 2020

	p, ok := pf.Make(ts, "prev week")
	if !ok {
		fmt.Println("the shortcut doesn't exist")
	}
	fmt.Println(p)
	// previous week is (August 3, 2020 — August 9, 2020)

	p = pf.Require(ts, rdate.PeriodPrevQuart)
	fmt.Println(p)
	// previous quarter is (April 1, 2020 — June 30, 2020)

	p = pf.Require(ts, rdate.PeriodThisMonth)
	fmt.Println(p)
	// this month is (August 1, 2020 — August 31, 2020)

	// Output:
	// today is August 11, 2020
	// previous week is (August 3, 2020 — August 9, 2020)
	// previous quarter is (April 1, 2020 — June 30, 2020)
	// this month is (August 1, 2020 — August 31, 2020)
}
