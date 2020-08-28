package rdate_test

import (
	"testing"
	"time"

	"github.com/petrunkodg/rdate"
)

func TestNewPeriodFactory(t *testing.T) {
	f1 := rdate.NewPeriodFactory()
	f2 := rdate.NewPeriodFactory()

	if f1 == f2 {
		t.Errorf("both the factories are the same")
	}
}

func TestDefaultPediodFactory_existenceOfRules(t *testing.T) {
	testCases := []struct {
		name         string
		pivot        time.Time
		sc           rdate.PeriodShortcut
		expectedFrom time.Time
		expectedTo   time.Time
	}{
		{
			name:         "PeriodThisDay",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodThisDay,
			expectedFrom: time.Date(2019, 12, 11, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 11, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodThisWeek",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodThisWeek,
			expectedFrom: time.Date(2019, 12, 9, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 15, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodThisMonth",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodThisMonth,
			expectedFrom: time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodThisQuart",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodThisQuart,
			expectedFrom: time.Date(2019, 10, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodThisHalfYear",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodThisHalfYear,
			expectedFrom: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodThisYear",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodThisYear,
			expectedFrom: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodPrevDay",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodPrevDay,
			expectedFrom: time.Date(2019, 12, 10, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 10, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodPrevWeek",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodPrevWeek,
			expectedFrom: time.Date(2019, 12, 2, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 8, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodPrevMonth",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodPrevMonth,
			expectedFrom: time.Date(2019, 11, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 11, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodPrevQuart",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodPrevQuart,
			expectedFrom: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodPrevHalfYear",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodPrevHalfYear,
			expectedFrom: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PeriodPrevYear",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodPrevYear,
			expectedFrom: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2018, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, ok := rdate.NewPeriod(tc.pivot, tc.sc)

			if !ok {
				t.Errorf("expected ok but it isn't")
			}

			periodEqual(t, actual, tc.expectedFrom, tc.expectedTo)
		})
	}
}

func TestNewPeriod(t *testing.T) {
	testCases := []struct {
		name         string
		pivot        time.Time
		sc           rdate.PeriodShortcut
		expectedFrom time.Time
		expectedTo   time.Time
		expectedOK   bool
	}{
		{
			name:         "ok",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodPrevWeek,
			expectedFrom: time.Date(2019, 12, 2, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 8, 23, 59, 59, 999999999, time.UTC),
			expectedOK:   true,
		},
		{
			name:         "fail",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           "past week 1",
			expectedFrom: time.Time{},
			expectedTo:   time.Time{},
			expectedOK:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, ok := rdate.NewPeriod(tc.pivot, tc.sc)
			if ok != tc.expectedOK {
				t.Errorf("expected %t but it isn't %t", tc.expectedOK, ok)
			}

			periodEqual(t, actual, tc.expectedFrom, tc.expectedTo)
		})
	}
}

func TestRequirePeriod(t *testing.T) {
	testCases := []struct {
		name         string
		pivot        time.Time
		sc           rdate.PeriodShortcut
		expectedFrom time.Time
		expectedTo   time.Time
	}{
		{
			name:         "ok",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           rdate.PeriodPrevWeek,
			expectedFrom: time.Date(2019, 12, 2, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 8, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "fail",
			pivot:        time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:           "past week 1",
			expectedFrom: time.Time{},
			expectedTo:   time.Time{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := rdate.RequirePeriod(tc.pivot, tc.sc)
			periodEqual(t, actual, tc.expectedFrom, tc.expectedTo)
		})
	}
}

type testPeriodRule struct{}

func (p *testPeriodRule) Calculate(pivot time.Time,
	tf *rdate.TimeFactory) (from, to rdate.Time) {
	return tf.Require(time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC), rdate.TimeAsIs),
		tf.Require(time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC), rdate.TimeAsIs)
}

func (p *testPeriodRule) Shortcut() rdate.PeriodShortcut { return "my test period" }

func TestPeriodFactory_Extend(t *testing.T) {
	pivot := time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC)

	f := rdate.NewPeriodFactory()

	p, ok := f.Make(pivot, "my test period")

	if ok {
		t.Errorf("expected false but it is true")
	}
	if !p.IsZero() {
		t.Errorf("expected p has a zero-value but it doesn't")
	}

	f.Extend([]rdate.PeriodRule{&testPeriodRule{}})

	p, ok = f.Make(pivot, "my test period")

	if !ok {
		t.Errorf("expected ok but it isn't")
	}

	periodEqual(t, p,
		time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC),
		time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC))
}

type testPeriodStringer struct{}

func (s *testPeriodStringer) String(from, to rdate.Time,
	sc rdate.PeriodShortcut) string {
	return "test period stringer"
}

func TestPeriodFactory_SetStringer(t *testing.T) {
	expected := []string{
		"2010-02-22 00:00:00 — 2010-02-28 23:59:59",
		"test period stringer",
	}
	pivot := time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC)

	f := rdate.NewPeriodFactory()

	p, ok := f.Make(pivot, rdate.PeriodPrevWeek)

	if !ok {
		t.Errorf("expected ok but it isn't")
	}
	if p.String() != expected[0] {
		t.Errorf("expected '%s' but there is '%s'", expected[0], p.String())
	}

	f.SetStringer(&testPeriodStringer{})

	p, ok = f.Make(pivot, rdate.PeriodPrevWeek)

	if !ok {
		t.Errorf("expected ok but it isn't")
	}
	if p.String() != expected[1] {
		t.Errorf("expected '%s' but there is '%s'", expected[1], p.String())
	}
}

func TestPeriodFactory_SetNilStringer(t *testing.T) {
	pf := rdate.NewPeriodFactory()
	pf.SetStringer(nil)

	pm := pf.Require(time.Now(), rdate.PeriodPrevWeek)
	if pm.String() != "" {
		t.Errorf("expected an empty string")
	}
}

func TestPeriodFactory_SetTimeFactory(t *testing.T) {
	expected := []string{
		"2010-02-22 00:00:00 — 2010-02-28 23:59:59",
		"test stringer — test stringer",
	}
	pivot := time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC)

	f := rdate.NewPeriodFactory()
	p, ok := f.Make(pivot, rdate.PeriodPrevWeek)

	if !ok {
		t.Errorf("expected ok but it isn't")
	}
	if p.String() != expected[0] {
		t.Errorf("expected '%s' but there is '%s'", expected[0], p.String())
	}

	tf := rdate.NewTimeFactory()
	tf.SetStringer(&testTimeStringer{})

	f.SetTimeFactory(tf)

	p, ok = f.Make(pivot, rdate.PeriodPrevWeek)

	if !ok {
		t.Errorf("expected ok but it isn't")
	}
	if p.String() != expected[1] {
		t.Errorf("expected '%s' but there is '%s'", expected[1], p.String())
	}
}

func TestSetDefaultPeriodFactory(t *testing.T) {
	pivot := time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC)

	p, ok := rdate.NewPeriod(pivot, "my test period")

	if ok {
		t.Errorf("expected false but it is true")
	}

	if !p.IsZero() {
		t.Errorf("expected p has a zero-value but it doesn't")
	}

	f := rdate.NewPeriodFactory()
	f.Extend([]rdate.PeriodRule{&testPeriodRule{}})

	rdate.SetDefaultPeriodFactory(f)

	p, ok = rdate.NewPeriod(pivot, "my test period")

	if !ok {
		t.Errorf("expected ok but it isn't")
	}

	periodEqual(t, p,
		time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC),
		time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC))
}

func periodEqual(t *testing.T, actual rdate.Period, expectedFrom, expectedTo time.Time) {
	t.Helper()

	if !expectedFrom.Equal(actual.From().Time()) {
		t.Errorf("Date = %s; expected %s", actual.From().Time(), expectedFrom)
	}

	if !expectedTo.Equal(actual.To().Time()) {
		t.Errorf("Date = %s; expected %s", actual.To().Time(), expectedTo)
	}
}
