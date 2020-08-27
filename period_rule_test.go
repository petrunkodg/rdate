package rdate

import (
	"testing"
	"time"
)

func TestPeriodRules_Perform(t *testing.T) {
	testCases := []struct {
		name         string
		rule         PeriodRule
		date         time.Time
		expectedFrom time.Time
		expectedTo   time.Time
	}{
		{
			name:         "ThisDay",
			rule:         &periodRuleThisDay{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 12, 11, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 11, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "ThisWeek",
			rule:         &periodRuleThisWeek{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 12, 9, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 15, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "ThisMonth",
			rule:         &periodRuleThisMonth{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "ThisQuart",
			rule:         &periodRuleThisQuart{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 10, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "ThisHalfYear",
			rule:         &periodRuleThisHalfYear{},
			date:         time.Date(2019, 8, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "ThisYear",
			rule:         &periodRuleThisYear{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PrevDay",
			rule:         &periodRulePrevDay{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 12, 10, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 10, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PrevWeek",
			rule:         &periodRulePrevWeek{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 12, 2, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 12, 8, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PrevMonth",
			rule:         &periodRulePrevMonth{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 11, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 11, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PrevQuart",
			rule:         &periodRulePrevQuart{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PrevHalfYear",
			rule:         &periodRulePrevHalfYear{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2019, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:         "PrevYear",
			rule:         &periodRulePrevYear{},
			date:         time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expectedFrom: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2018, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
	}

	tf := NewTimeFactory()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualFrom, actualTo := tc.rule.Calculate(tc.date, tf)

			if !actualFrom.Time().Equal(tc.expectedFrom) {
				t.Errorf("Period.From = %s; expected %s", actualFrom.Time(), tc.expectedFrom)
			}

			if !actualTo.Time().Equal(tc.expectedTo) {
				t.Errorf("Period.To = %s; expected %s", actualTo.Time(), tc.expectedTo)
			}
		})
	}
}
