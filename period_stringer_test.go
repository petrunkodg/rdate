package rdate_test

import (
	"testing"
	"time"

	"github.com/petrunkodg/rdate"
)

func TestDefaultPeriodStringer(t *testing.T) {
	from := rdate.RequireTime(time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
		rdate.TimeAsIs)
	to := rdate.RequireTime(time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
		rdate.TimeAsIs)

	testCases := []struct {
		name     string
		sc       rdate.PeriodShortcut
		expected string
	}{
		{
			name:     "TimeAsIs",
			sc:       rdate.PeriodPrevWeek,
			expected: "2019-12-11 00:02:01 — 2019-12-11 00:02:01",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := rdate.DefaultPeriodStringer.String(from, to, tc.sc)
			if actual != tc.expected {
				t.Errorf("expected: '%s', but actual: '%s'", tc.expected, actual)
			}
		})
	}
}
