// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rdate

import (
	"testing"
	"time"
)

func TestTimeRules_Perform(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		rule     TimeRule
		expected time.Time
	}{
		{
			name:     "AsIs",
			rule:     &timeRuleAsIs{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
		},
		{
			name:     "StartOfThisDay",
			rule:     &timeRuleStartOfThisDay{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.FixedZone("UTC-8", -8*60*60)),
			expected: time.Date(2019, 12, 11, 0, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60)),
		},
		{
			name:     "EndOfThisDay",
			rule:     &timeRuleEndOfThisDay{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 11, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfPrevDay",
			rule:     &timeRuleStartOfPrevDay{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.FixedZone("UTC+4", 4*60*60)),
			expected: time.Date(2019, 12, 10, 0, 0, 0, 0, time.FixedZone("UTC+4", 4*60*60)),
		},
		{
			name:     "EndOfPrevDay",
			rule:     &timeRuleEndOfPrevDay{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 10, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfThisWeek",
			rule:     &timeRuleStartOfThisWeek{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.FixedZone("UTC-1", -1*60*60)),
			expected: time.Date(2019, 12, 9, 0, 0, 0, 0, time.FixedZone("UTC-1", -1*60*60)),
		},
		{
			name:     "EndOfThisWeek",
			rule:     &timeRuleEndOfThisWeek{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 15, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfThisWeek(Sunday)",
			rule:     &timeRuleStartOfThisWeek{},
			date:     time.Date(2020, 8, 9, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, 8, 3, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfThisWeek(Sunday)",
			rule:     &timeRuleEndOfThisWeek{},
			date:     time.Date(2020, 8, 9, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, 8, 9, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfThisWeekS(the first day of the week is Sunday)",
			rule:     &timeRuleStartOfThisWeekS{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.FixedZone("UTC+8", 8*60*60)),
			expected: time.Date(2019, 12, 8, 0, 0, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
		},
		{
			name:     "EndOfThisWeekS(the first day of the week is Sunday)",
			rule:     &timeRuleEndOfThisWeekS{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 14, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfThisMonth",
			rule:     &timeRuleStartOfThisMonth{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfThisMonth",
			rule:     &timeRuleEndOfThisMonth{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfThisQuart",
			rule:     &timeRuleStartOfThisQuart{},
			date:     time.Date(2019, 7, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfThisQuart",
			rule:     &timeRuleEndOfThisQuart{},
			date:     time.Date(2019, 7, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfThisHalfYear",
			rule:     &timeRuleStartOfThisHalfYear{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfThisHalfYear",
			rule:     &timeRuleEndOfThisHalfYear{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfThisYear",
			rule:     &timeRuleStartOfThisYear{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfThisYear",
			rule:     &timeRuleEndOfThisYear{},
			date:     time.Date(2019, 3, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfPrevWeek",
			rule:     &timeRuleStartOfPrevWeek{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfPrevWeek",
			rule:     &timeRuleEndOfPrevWeek{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 8, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfPrevWeek(Sunday)",
			rule:     &timeRuleStartOfPrevWeek{},
			date:     time.Date(2020, 8, 9, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, 7, 27, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfPrevWeek(Sunday)",
			rule:     &timeRuleEndOfPrevWeek{},
			date:     time.Date(2020, 8, 9, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, 8, 2, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfPrevWeekS(the first day of the week is Sunday)",
			rule:     &timeRuleStartOfPrevWeekS{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfPrevWeekS(the first day of the week is Sunday)",
			rule:     &timeRuleEndOfPrevWeekS{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 12, 7, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfPrevMonth",
			rule:     &timeRuleStartOfPrevMonth{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 11, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfPrevMonth",
			rule:     &timeRuleEndOfPrevMonth{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 11, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfPrevQuart",
			rule:     &timeRuleStartOfPrevQuart{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfPrevQuart",
			rule:     &timeRuleEndOfPrevQuart{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfPrevHalfYear",
			rule:     &timeRuleStartOfPrevHalfYear{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfPrevHalfYear",
			rule:     &timeRuleEndOfPrevHalfYear{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2019, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name:     "StartOfPrevYear",
			rule:     &timeRuleStartOfPrevYear{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "EndOfPrevYear",
			rule:     &timeRuleEndOfPrevYear{},
			date:     time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC),
			expected: time.Date(2018, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.rule.Calculate(tc.date)
			if !actual.Equal(tc.expected) {
				t.Errorf("Date = %s; expected %s", actual, tc.expected)
			}
		})
	}
}
