// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rdate

import (
	"sync"
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

// PeriodFactory is used to make new Period objects by passing the pivot and the shortcut.
// Method NewPeriod and RequirePeriod uses the default period factory which is created
// by calling NewPeriodFactory during the init.
type PeriodFactory interface {
	// Make creates a new Period object by using the rule which is found (or not)
	// by the given PeriodShortcut.
	// If the rule is not found, ok will be false and t will be a zero-value of Period.
	Make(pivot time.Time, sc PeriodShortcut) (p Period, ok bool)

	// Require creates new Period object by using the rule which is found (or not)
	// by the given PeriodShortcut.
	// If the rule is not found, the result will be a zero-value of Period type.
	// This method should be used only if you are sure about existence of given shortcut.
	Require(pivot time.Time, sc PeriodShortcut) Period

	// Extend appends new rules (or replaces existing ones if there are any rules
	// with the same shortcuts) to the period factory.
	Extend(rules []PeriodRule)

	// SetTimeFactory sets your own TimeFactory which will be passed to
	// a rule Calculate method during the calculation of a Make call.
	SetTimeFactory(tf TimeFactory)

	// SetStringer sets your own PeriodStringer implementation
	// for every new Period object which is created by this factory.
	SetStringer(s PeriodStringer)
}

type unsafePeriodFactory struct {
	rules map[PeriodShortcut]PeriodRule
	tf    TimeFactory
	s     PeriodStringer
}

func newUnsafePeriodFactory(rules []PeriodRule, tf TimeFactory,
	s PeriodStringer) PeriodFactory {
	f := &unsafePeriodFactory{
		rules: map[PeriodShortcut]PeriodRule{},
		tf:    tf,
		s:     s,
	}

	f.Extend(rules)

	return f
}

// Make implements the PeriodFactory Make method.
func (f *unsafePeriodFactory) Make(pivot time.Time, sc PeriodShortcut) (p Period, ok bool) {
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

// Require implements the PeriodFactory Require method.
func (f *unsafePeriodFactory) Require(pivot time.Time, sc PeriodShortcut) Period {
	p, ok := f.Make(pivot, sc)
	if !ok {
		return Period{}
	}

	return p
}

// SetTimeFactory implements the PeriodFactory SetTimeFactory method.
func (f *unsafePeriodFactory) SetTimeFactory(tf TimeFactory) {
	f.tf = tf
}

// SetStringer implements the PeriodFactory SetStringer method.
func (f *unsafePeriodFactory) SetStringer(s PeriodStringer) {
	f.s = s
}

// Extend implements the PeriodFactory Extend method.
func (f *unsafePeriodFactory) Extend(rules []PeriodRule) {
	for _, r := range rules {
		f.rules[r.Shortcut()] = r
	}
}

var defaultPeriodRules = []PeriodRule{
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
}

var defaultPeriodFactory = NewPeriodFactory()

// safePeriodFactory is a decorator which wraps a period factory for concurrent use
// by multiple goroutines.
type safePeriodFactory struct {
	f  PeriodFactory
	rw sync.RWMutex
}

// Make implements the PeriodFactory Make method.
func (f *safePeriodFactory) Make(pivot time.Time, sc PeriodShortcut) (p Period, ok bool) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	return f.f.Make(pivot, sc)
}

// Require implements the PeriodFactory Require method.
func (f *safePeriodFactory) Require(pivot time.Time, sc PeriodShortcut) Period {
	f.rw.RLock()
	defer f.rw.RUnlock()

	return f.f.Require(pivot, sc)
}

// SetTimeFactory implements the PeriodFactory SetTimeFactory method.
func (f *safePeriodFactory) SetTimeFactory(tf TimeFactory) {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.f.SetTimeFactory(tf)
}

// SetStringer implements the PeriodFactory SetStringer method.
func (f *safePeriodFactory) SetStringer(s PeriodStringer) {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.f.SetStringer(s)
}

// Extend implements the PeriodFactory Extend method.
func (f *safePeriodFactory) Extend(rules []PeriodRule) {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.f.Extend(rules)
}

func newSafePeriodFactory(f PeriodFactory) PeriodFactory {
	return &safePeriodFactory{f: f}
}

// NewPeriodFactory creates a period factory which is ready to extend and
// safe for concurrent use by multiple goroutines.
func NewPeriodFactory() PeriodFactory {
	return newSafePeriodFactory(
		NewNonblockingPeriodFactory())
}

// NewNonblockingPeriodFactory creates an unsafe period factory which is ready to extend.
// It means that there is no synchronization mechanisms in it.
// On the one hand extending of this factory is not a thread-safe operation,
// but on the other hand it makes Make and Require methods to be non-blocking.
// (eliminates the cache coherency problem)
//
// It might be useful when your application is under high load and your period factory
// doesn't use Extend method at all or use it once during the init.
func NewNonblockingPeriodFactory() PeriodFactory {
	return newUnsafePeriodFactory(defaultPeriodRules,
		NewNonblockingTimeFactory(), &defaultPeriodStringer{})
}

// SetDefaultPeriodFactory sets your own period factory as default.
//
// After that you can use NewPeriod or RequirePeriod functions
// without concreting a factory like that rdate.NewPeriod(...)
// or rdate.RequirePeriod(...))
func SetDefaultPeriodFactory(f PeriodFactory) {
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
