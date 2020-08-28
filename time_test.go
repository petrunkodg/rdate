package rdate_test

import (
	"testing"
	"time"

	"github.com/petrunkodg/rdate"
)

func TestNewTimeFactory(t *testing.T) {
	f1 := rdate.NewTimeFactory()
	f2 := rdate.NewTimeFactory()

	if f1 == f2 {
		t.Errorf("both the factories are the same")
	}
}

func TestRequireTime(t *testing.T) {
	testCases := []struct {
		name     string
		pivot    time.Time
		sc       rdate.TimeShortcut
		expected time.Time
	}{
		{
			name:     "ok",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfPrevDay,
			expected: time.Date(2019, 12, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "fail",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       "PrevDay111",
			expected: time.Time{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := rdate.RequireTime(tc.pivot, tc.sc)
			timeEqual(t, actual, tc.expected)
		})
	}
}

func TestNewTime(t *testing.T) {
	testCases := []struct {
		name       string
		pivot      time.Time
		sc         rdate.TimeShortcut
		expected   time.Time
		expectedOK bool
	}{
		{
			name:       "ok",
			pivot:      time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:         rdate.TimeStartOfPrevDay,
			expected:   time.Date(2019, 12, 10, 0, 0, 0, 0, time.UTC),
			expectedOK: true,
		},
		{
			name:       "fail",
			pivot:      time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:         "prev day111",
			expected:   time.Time{},
			expectedOK: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, ok := rdate.NewTime(tc.pivot, tc.sc)
			if ok != tc.expectedOK {
				t.Errorf("expected %t but it isn't %t", tc.expectedOK, ok)
			}

			timeEqual(t, actual, tc.expected)
		})
	}
}

func TestDefaultTimeFactory_existenceOfRules(t *testing.T) {
	testCases := []struct {
		name     string
		pivot    time.Time
		sc       rdate.TimeShortcut
		expected time.Time
	}{
		{
			name:     "TimeAsIs",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeAsIs,
			expected: time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
		},
		{
			name:     "TimeStartOfThisDay",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfThisDay,
			expected: time.Date(2019, 12, 11, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfThisDay",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfThisDay,
			expected: time.Date(2019, 12, 11, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfPrevDay",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfPrevDay,
			expected: time.Date(2019, 12, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfPrevDay",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfPrevDay,
			expected: time.Date(2019, 12, 10, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfThisWeek",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfThisWeek,
			expected: time.Date(2019, 12, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfThisWeek",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfThisWeek,
			expected: time.Date(2019, 12, 15, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfThisMonth",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfThisMonth,
			expected: time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfThisMonth",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfThisMonth,
			expected: time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfThisQuart",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfThisQuart,
			expected: time.Date(2019, 10, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfThisQuart",
			pivot:    time.Date(2019, 7, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfThisQuart,
			expected: time.Date(2019, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfThisHalfYear",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfThisHalfYear,
			expected: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfThisHalfYear",
			pivot:    time.Date(2019, 3, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfThisHalfYear,
			expected: time.Date(2019, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfThisYear",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfThisYear,
			expected: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfThisYear",
			pivot:    time.Date(2019, 3, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfThisYear,
			expected: time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfPrevWeek",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfPrevWeek,
			expected: time.Date(2019, 12, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfPrevWeek",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfPrevWeek,
			expected: time.Date(2019, 12, 8, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfPrevMonth",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfPrevMonth,
			expected: time.Date(2019, 11, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfPrevMonth",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfPrevMonth,
			expected: time.Date(2019, 11, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfPrevQuart",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfPrevQuart,
			expected: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfPrevQuart",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfPrevQuart,
			expected: time.Date(2019, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfPrevHalfYear",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfPrevHalfYear,
			expected: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfPrevHalfYear",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfPrevHalfYear,
			expected: time.Date(2019, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "TimeStartOfPrevYear",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeStartOfPrevYear,
			expected: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "TimeEndOfPrevYear",
			pivot:    time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			sc:       rdate.TimeEndOfPrevYear,
			expected: time.Date(2018, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, ok := rdate.NewTime(tc.pivot, tc.sc)

			if !ok {
				t.Errorf("expected ok but it isn't")
			}

			timeEqual(t, actual, tc.expected)
		})
	}
}

type testTimeRule struct{}

func (r *testTimeRule) Calculate(pivot time.Time) time.Time {
	return time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC)
}

func (r *testTimeRule) Shortcut() rdate.TimeShortcut {
	return "my test time"
}

func TestTimeFactory_Extend(t *testing.T) {
	pivot := time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC)

	f := rdate.NewTimeFactory()

	tm, ok := f.Make(pivot, "my test time")

	if ok {
		t.Errorf("expected ok = false but it's true")
	}
	if !tm.IsZero() {
		t.Errorf("expected tm has a zero-value but it doesn't")
	}

	f.Extend([]rdate.TimeRule{&testTimeRule{}})

	tm, ok = f.Make(pivot, "my test time")

	if !ok {
		t.Errorf("expected ok but it isn't")
	}
	timeEqual(t, tm, time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC))
}

type testTimeStringer struct{}

func (s *testTimeStringer) String(t time.Time) string { return "test stringer" }

func TestTimeFactory_SetStringer(t *testing.T) {
	expected := []string{
		"2010-02-28 00:00:00",
		"test stringer",
	}

	pivot := time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC)

	f := rdate.NewTimeFactory()

	tm, ok := f.Make(pivot, rdate.TimeStartOfPrevDay)

	if !ok {
		t.Errorf("expected ok but it isn't")
	}
	if tm.String() != expected[0] {
		t.Errorf("expected '%s' but there is '%s'", expected[0], tm.String())
	}

	f.SetStringer(&testTimeStringer{})

	tm, ok = f.Make(pivot, rdate.TimeStartOfPrevDay)

	if !ok {
		t.Errorf("expected ok but it isn't")
	}
	if tm.String() != expected[1] {
		t.Errorf("expected '%s' but there is '%s'", expected[1], tm.String())
	}
}

func TestTimeFactory_SetNilStringer(t *testing.T) {
	tf := rdate.NewTimeFactory()
	tf.SetStringer(nil)

	tm := tf.Require(time.Now(), rdate.TimeAsIs)
	if tm.String() != "" {
		t.Errorf("expected an empty string")
	}
}

func TestTimeFactory_SetStartOfWeek(t *testing.T) {
	testCases := []struct {
		name           string
		ts             time.Time
		sc             rdate.TimeShortcut
		expectedMonday time.Time
		expectedSunday time.Time
	}{
		{
			name:           "StartOfThisWeek",
			ts:             time.Date(2020, 7, 8, 0, 2, 1, 6, time.UTC),
			sc:             rdate.TimeStartOfThisWeek,
			expectedMonday: time.Date(2020, 7, 6, 0, 0, 0, 0, time.UTC),
			expectedSunday: time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			name:           "StartOfPrevWeek",
			ts:             time.Date(2020, 7, 8, 0, 2, 1, 6, time.UTC),
			sc:             rdate.TimeStartOfPrevWeek,
			expectedMonday: time.Date(2020, 6, 29, 0, 0, 0, 0, time.UTC),
			expectedSunday: time.Date(2020, 6, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			name:           "TimeEndOfPrevWeek",
			ts:             time.Date(2020, 7, 8, 0, 2, 1, 6, time.UTC),
			sc:             rdate.TimeEndOfPrevWeek,
			expectedMonday: time.Date(2020, 7, 5, 23, 59, 59, 999999999, time.UTC),
			expectedSunday: time.Date(2020, 7, 4, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := rdate.NewTimeFactory()
			f.SetStartOfWeek(rdate.StartOfWeekMonday)

			actual, ok := f.Make(tc.ts, tc.sc)

			if !ok || !tc.expectedMonday.Equal(actual.Time()) {
				t.Errorf("Time = %s; expected %s", actual.Time(), tc.expectedMonday)
			}

			f.SetStartOfWeek(rdate.StartOfWeekSunday)

			actual, ok = f.Make(tc.ts, tc.sc)

			if !ok || !tc.expectedSunday.Equal(actual.Time()) {
				t.Errorf("Time = %s; expected %s", actual.Time(), tc.expectedSunday)
			}
		})
	}
}

func TestSetDefaultTimeFactory(t *testing.T) {
	pivot := time.Date(2010, 3, 1, 0, 2, 1, 6, time.UTC)

	tm, ok := rdate.NewTime(pivot, "my test time")

	if ok {
		t.Errorf("expected ok = false but it's true")
	}
	if !tm.IsZero() {
		t.Errorf("expected tm has a zero-value but doesn't")
	}

	f := rdate.NewTimeFactory()
	f.Extend([]rdate.TimeRule{&testTimeRule{}})

	rdate.SetDefaultTimeFactory(f)

	tm, ok = rdate.NewTime(pivot, "my test time")

	if !ok {
		t.Errorf("expected ok but it isn't")
	}
	timeEqual(t, tm, time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC))
}

func timeEqual(t *testing.T, actual rdate.Time, expected time.Time) {
	t.Helper()

	if !expected.Equal(actual.Time()) {
		t.Errorf("Date = %s; expected %s", actual.Time(), expected)
	}
}
