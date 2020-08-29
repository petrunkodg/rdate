// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rdate

import (
	"time"
)

type PeriodShortcut string

const (
	PeriodThisDay      PeriodShortcut = "this day"
	PeriodThisWeek     PeriodShortcut = "this week"
	PeriodThisMonth    PeriodShortcut = "this month"
	PeriodThisQuart    PeriodShortcut = "this quart"
	PeriodThisHalfYear PeriodShortcut = "this half year"
	PeriodThisYear     PeriodShortcut = "this year"
	PeriodPrevDay      PeriodShortcut = "prev day"
	PeriodPrevWeek     PeriodShortcut = "prev week"
	PeriodPrevMonth    PeriodShortcut = "prev month"
	PeriodPrevQuart    PeriodShortcut = "prev quart"
	PeriodPrevHalfYear PeriodShortcut = "prev half year"
	PeriodPrevYear     PeriodShortcut = "prev year"
)

type PeriodFactory struct {
	rules map[PeriodShortcut]PeriodRule
	tf    *TimeFactory
	s     PeriodStringer
}

func newPeriodFactory(rules []PeriodRule, tf *TimeFactory, s PeriodStringer) *PeriodFactory {
	f := &PeriodFactory{
		rules: map[PeriodShortcut]PeriodRule{},
		tf:    tf,
		s:     s,
	}

	f.Extend(rules)

	return f
}

// Make creates a new Period object using the rule which is found (or not)
// by the given PeriodShortcut.
//
// If the rule is not found, ok will be false and t will be a zero-value of Period.
func (f *PeriodFactory) Make(pivot time.Time, sc PeriodShortcut) (p Period, ok bool) {
	r, ok := f.rules[sc]
	if !ok {
		return Period{}, false
	}

	from, to := r.Calculate(pivot, f.tf)

	return Period{
		from: from,
		to:   to,
		sc:   sc,
		s:    f.s,
	}, true
}

// Require creates new Period object using the rule which is found (or not)
// by the given PeriodShortcut.
//
// If the rule is not found, the result will be a zero-value of Period type.
//
// This method should be used only if you are sure about existence of given shortcut.
func (f *PeriodFactory) Require(pivot time.Time, sc PeriodShortcut) Period {
	p, ok := f.Make(pivot, sc)
	if !ok {
		return Period{}
	}

	return p
}

// SetTimeFactory sets your own TimeFactory which will be passed to
// a rule Calculate method during the calculation of a Make call.
func (f *PeriodFactory) SetTimeFactory(tf *TimeFactory) {
	f.tf = tf
}

// SetStringer sets your own PeriodStringer implementation
// for every new Period object which is created by this factory.
func (f *PeriodFactory) SetStringer(s PeriodStringer) {
	f.s = s
}

// Extend appends new rules or replaces existing ones
// (if there were any rules with the same shortcuts) to the period factory.
func (f *PeriodFactory) Extend(rules []PeriodRule) {
	for _, r := range rules {
		f.rules[r.Shortcut()] = r
	}
}

func (f *PeriodFactory) copy() *PeriodFactory {
	rules := make([]PeriodRule, 0, len(f.rules))
	for _, r := range f.rules {
		rules = append(rules, r)
	}

	return newPeriodFactory(rules, f.tf.copy(), f.s)
}

var defaultPeriodFactory = newPeriodFactory([]PeriodRule{
	&periodRuleThisDay{},
	&periodRuleThisWeek{},
	&periodRuleThisMonth{},
	&periodRuleThisQuart{},
	&periodRuleThisHalfYear{},
	&periodRuleThisYear{},
	&periodRulePrevDay{},
	&periodRulePrevWeek{},
	&periodRulePrevMonth{},
	&periodRulePrevQuart{},
	&periodRulePrevHalfYear{},
	&periodRulePrevYear{},
}, defaultTimeFactory, DefaultPeriodStringer)

// NewPeriodFactory returns a copy of the default time factory
func NewPeriodFactory() *PeriodFactory {
	return defaultPeriodFactory.copy()
}

// SetDefaultPeriodFactory sets your own period factory as default.
//
// After that you can use NewPeriod or RequirePeriod functions
// without concreting a factory like that rdate.NewPeriod(...)
// or rdate.RequirePeriod(...))
func SetDefaultPeriodFactory(f *PeriodFactory) {
	defaultPeriodFactory = f
}

type Period struct {
	from Time
	to   Time
	sc   PeriodShortcut
	s    PeriodStringer
}

// NewPeriod calls Make method of the default period factory.
func NewPeriod(pivot time.Time, sc PeriodShortcut) (p Period, ok bool) {
	return defaultPeriodFactory.Make(pivot, sc)
}

// RequirePeriod calls Require method of the default period factory.
func RequirePeriod(pivot time.Time, sc PeriodShortcut) Period {
	return defaultPeriodFactory.Require(pivot, sc)
}

// From is a getter of the from Time value of the type.
func (p Period) From() Time {
	return p.from
}

// To is a getter of the to Time value of the type.
func (p Period) To() Time {
	return p.to
}

func (p Period) String() string {
	if p.s == nil {
		return ""
	}

	return p.s.String(p.from, p.to, p.sc)
}

// IsZero reports if the value is a zero-value of the type
func (p Period) IsZero() bool {
	return p.s == nil && p.from.IsZero() && p.to.IsZero()
}