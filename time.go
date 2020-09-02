// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rdate

import (
	"sync"
	"time"
)

type TimeShortcut string

const (
	TimeAsIs                TimeShortcut = "as is"
	TimeStartOfThisDay      TimeShortcut = "start this day"
	TimeEndOfThisDay        TimeShortcut = "end this day"
	TimeStartOfThisWeek     TimeShortcut = "start this week"
	TimeEndOfThisWeek       TimeShortcut = "end this week"
	TimeStartOfThisMonth    TimeShortcut = "start this month"
	TimeEndOfThisMonth      TimeShortcut = "end this month"
	TimeStartOfThisQuart    TimeShortcut = "start this quart"
	TimeEndOfThisQuart      TimeShortcut = "end this quart"
	TimeStartOfThisHalfYear TimeShortcut = "start this half year"
	TimeEndOfThisHalfYear   TimeShortcut = "end this half year"
	TimeStartOfThisYear     TimeShortcut = "start this year"
	TimeEndOfThisYear       TimeShortcut = "end this year"
	TimeStartOfPrevDay      TimeShortcut = "start prev day"
	TimeEndOfPrevDay        TimeShortcut = "end prev day"
	TimeStartOfPrevWeek     TimeShortcut = "start prev week"
	TimeEndOfPrevWeek       TimeShortcut = "end prev week"
	TimeStartOfPrevMonth    TimeShortcut = "start prev month"
	TimeEndOfPrevMonth      TimeShortcut = "end prev month"
	TimeStartOfPrevQuart    TimeShortcut = "start prev quart"
	TimeEndOfPrevQuart      TimeShortcut = "end prev quart"
	TimeStartOfPrevHalfYear TimeShortcut = "start prev half year"
	TimeEndOfPrevHalfYear   TimeShortcut = "end prev half year"
	TimeStartOfPrevYear     TimeShortcut = "start prev year"
	TimeEndOfPrevYear       TimeShortcut = "end prev year"
)

type StartOfWeek int8

const (
	StartOfWeekMonday StartOfWeek = iota + 1
	StartOfWeekSunday
)

// TimeFactory is used to make new Time objects by passing the pivot and the shortcut.
// Method NewTime and RequireTime uses the default time factory which is created
// by calling NewTimeFactory during the init.
type TimeFactory interface {
	// Make creates a new Time object by using the rule which is found (or not)
	// by the given TimeShortcut.
	// If the rule is not found, ok will be false and t will be a zero-value of Time.
	Make(pivot time.Time, sc TimeShortcut) (t Time, ok bool)

	// Require creates new Time object by using the rule which is found (or not)
	// by the given TimeShortcut.
	// If the rule is not found, the result will be a zero-value of Time.
	// This method should be used only if you are sure about existence of given shortcut.
	Require(pivot time.Time, sc TimeShortcut) Time

	// Extend appends new rules (or replaces existing ones if there are any rules
	// with the same shortcuts) to the time factory.
	Extend(rules []TimeRule)

	// SetStartOfWeek sets the start of the week for the time factory.
	// It can be Monday or Sunday.
	// The default value is Monday.
	SetStartOfWeek(s StartOfWeek)

	// SetStringer sets your own TimeStringer implementation
	// for every new Time object which is created by this factory.
	SetStringer(s TimeStringer)
}

type unsafeTimeFactory struct {
	rules map[TimeShortcut]TimeRule
	s     TimeStringer
}

func newUnsafeTimeFactory(rules []TimeRule, s TimeStringer) TimeFactory {
	f := &unsafeTimeFactory{
		rules: map[TimeShortcut]TimeRule{},
		s:     s,
	}

	f.Extend(rules)

	return f
}

// Make implements the TimeFactory Make method.
func (f *unsafeTimeFactory) Make(pivot time.Time, sc TimeShortcut) (t Time, ok bool) {
	r, ok := f.rules[sc]
	if !ok {
		return Time{}, false
	}

	return Time{
		t: r.Calculate(pivot),
		s: f.s,
	}, true
}

// Require implements the TimeFactory Require method.
func (f *unsafeTimeFactory) Require(pivot time.Time, sc TimeShortcut) Time {
	t, ok := f.Make(pivot, sc)
	if !ok {
		t = Time{}
	}

	return t
}

// SetStringer implements the TimeFactory SetStringer method.
func (f *unsafeTimeFactory) SetStringer(s TimeStringer) {
	f.s = s
}

// Extend implements the TimeFactory Extend method.
func (f *unsafeTimeFactory) Extend(rules []TimeRule) {
	for _, r := range rules {
		f.rules[r.Shortcut()] = r
	}
}

// SetStartOfWeek implements the TimeFactory SetStartOfWeek method.
func (f *unsafeTimeFactory) SetStartOfWeek(s StartOfWeek) {
	var rules []TimeRule
	switch s {
	case StartOfWeekMonday:
		rules = []TimeRule{
			&timeRuleStartOfThisWeek{},
			&timeRuleEndOfThisWeek{},
			&timeRuleStartOfPrevWeek{},
			&timeRuleEndOfPrevWeek{},
		}
	case StartOfWeekSunday:
		rules = []TimeRule{
			&timeRuleStartOfThisWeekS{},
			&timeRuleEndOfThisWeekS{},
			&timeRuleStartOfPrevWeekS{},
			&timeRuleEndOfPrevWeekS{},
		}
	}

	f.Extend(rules)
}

var defaultRules = []TimeRule{
	&timeRuleAsIs{},
	thisDayStartRule,
	thisDayEndRule,
	&timeRuleStartOfPrevDay{},
	&timeRuleEndOfPrevDay{},
	&timeRuleStartOfThisWeek{},
	&timeRuleEndOfThisWeek{},
	&timeRuleStartOfPrevWeek{},
	&timeRuleEndOfPrevWeek{},
	&timeRuleStartOfThisMonth{},
	&timeRuleEndOfThisMonth{},
	&timeRuleStartOfPrevMonth{},
	&timeRuleEndOfPrevMonth{},
	&timeRuleStartOfThisQuart{},
	&timeRuleEndOfThisQuart{},
	&timeRuleStartOfPrevQuart{},
	&timeRuleEndOfPrevQuart{},
	&timeRuleStartOfThisHalfYear{},
	&timeRuleEndOfThisHalfYear{},
	&timeRuleStartOfPrevHalfYear{},
	&timeRuleEndOfPrevHalfYear{},
	&timeRuleStartOfThisYear{},
	&timeRuleEndOfThisYear{},
	&timeRuleStartOfPrevYear{},
	&timeRuleEndOfPrevYear{},
}

var defaultTimeFactory = NewTimeFactory()

// safeTimeFactory is a decorator which wraps a time factory for concurrent use
// by multiple goroutines.
type safeTimeFactory struct {
	f  TimeFactory
	rw sync.RWMutex
}

// Make implements the TimeFactory Make method.
func (f *safeTimeFactory) Make(pivot time.Time, sc TimeShortcut) (t Time, ok bool) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	return f.f.Make(pivot, sc)
}

// Require implements the TimeFactory Require method.
func (f *safeTimeFactory) Require(pivot time.Time, sc TimeShortcut) Time {
	f.rw.RLock()
	defer f.rw.RUnlock()

	return f.f.Require(pivot, sc)
}

// SetStringer implements the TimeFactory SetStringer method.
func (f *safeTimeFactory) SetStringer(s TimeStringer) {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.f.SetStringer(s)
}

// Extend implements the TimeFactory Extend method.
func (f *safeTimeFactory) Extend(rules []TimeRule) {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.f.Extend(rules)
}

// SetStartOfWeek implements the TimeFactory SetStartOfWeek method.
func (f *safeTimeFactory) SetStartOfWeek(s StartOfWeek) {
	f.rw.Lock()
	defer f.rw.Unlock()

	f.f.SetStartOfWeek(s)
}

func newSafeTimeFactory(f TimeFactory) TimeFactory {
	return &safeTimeFactory{f: f}
}

// NewTimeFactory creates a time factory which is ready to extend and
// safe for concurrent use by multiple goroutines.
func NewTimeFactory() TimeFactory {
	return newSafeTimeFactory(
		NewNonblockingTimeFactory())
}

// NewNonblockingTimeFactory creates an unsafe time factory which is ready to extend.
// It means that there is no synchronization mechanisms in it.
// On the one hand extending of this factory is not a thread-safe operation,
// but on the other hand it makes Make and Require methods to be non-blocking.
// (eliminates the cache coherency problem)
//
// It might be useful when your application is under high load and your time factory
// doesn't use Extend method at all or use it once during the init.
func NewNonblockingTimeFactory() TimeFactory {
	return newUnsafeTimeFactory(defaultRules, &defaultTimeStringer{})
}

// SetDefaultTimeFactory sets your own time factory as default.
//
// After that you can use NewTime or RequireTime functions
// without concreting a factory like that rdate.NewTime(...)
// or rdate.RequireTime(...))
func SetDefaultTimeFactory(f TimeFactory) {
	defaultTimeFactory = f
}

// SetDefaultStartOfWeek sets the start of the week for the default time factory.
// It can be Monday or Sunday.
func SetDefaultStartOfWeek(s StartOfWeek) {
	defaultTimeFactory.SetStartOfWeek(s)
}

type Time struct {
	t time.Time
	s TimeStringer
}

// NewTime calls Make method of the default time factory.
func NewTime(pivot time.Time, sc TimeShortcut) (t Time, ok bool) {
	return defaultTimeFactory.Make(pivot, sc)
}

// RequireTime calls Require method of the default time factory.
func RequireTime(pivot time.Time, sc TimeShortcut) Time {
	return defaultTimeFactory.Require(pivot, sc)
}

// Time is a getter of the internal time.Time value of the type.
func (t Time) Time() time.Time {
	return t.t
}

func (t Time) String() string {
	if t.s == nil {
		return ""
	}

	return t.s.String(t.t)
}

// IsZero reports if the value is a zero-value of the type
func (t Time) IsZero() bool {
	return t.s == nil && t.t.IsZero()
}
