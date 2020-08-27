package rdate

import "time"

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

type TimeFactory struct {
	rules map[TimeShortcut]TimeRule
	s     TimeStringer
}

func newTimeFactory(rules []TimeRule, s TimeStringer) *TimeFactory {
	f := &TimeFactory{
		rules: map[TimeShortcut]TimeRule{},
		s:     s,
	}

	f.Extend(rules)

	return f
}

// Make creates a new Time object using the rule which is found (or not)
// by the given TimeShortcut.
//
// If the rule is not found, ok will be false and t will be a zero-value of Time.
func (f *TimeFactory) Make(pivot time.Time, sc TimeShortcut) (t Time, ok bool) {
	r, ok := f.rules[sc]
	if !ok {
		return Time{}, false
	}

	return Time{
		t: r.Calculate(pivot),
		s: f.s,
	}, true
}

// Require creates new Time object using the rule which is found (or not)
// by the given TimeShortcut.
//
// If the rule is not found, the result will be a zero-value of Time.
//
// This method should be used only if you are sure about existance of given shortcut.
func (f *TimeFactory) Require(pivot time.Time, sc TimeShortcut) Time {
	t, ok := f.Make(pivot, sc)
	if !ok {
		t = Time{}
	}

	return t
}

// SetStringer sets your own TimeStringer implementation
// for every new Time object which is created by this factory.
func (f *TimeFactory) SetStringer(s TimeStringer) {
	f.s = s
}

// Extend appends new rules or replaces existing ones
// (if there were any rules with the same shortcuts) to the time factory.
func (f *TimeFactory) Extend(rules []TimeRule) {
	for _, r := range rules {
		f.rules[r.Shortcut()] = r
	}
}

// SetStartOfWeek sets the start of the week for this time factory.
// It can be Monday or Sunday.
//
// The default value is Monday.
func (f *TimeFactory) SetStartOfWeek(s StartOfWeek) {
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

func (f *TimeFactory) copy() *TimeFactory {
	rules := make([]TimeRule, 0, len(f.rules))
	for _, r := range f.rules {
		rules = append(rules, r)
	}

	return newTimeFactory(rules, f.s)
}

var defaultTimeFactory = newTimeFactory([]TimeRule{
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
}, DefaultTimeStringer)

// NewTimeFactory returns a copy of the default time factory
func NewTimeFactory() *TimeFactory {
	return defaultTimeFactory.copy()
}

// SetDefaultTimeFactory sets your own time factory as default.
//
// After that you can use NewTime or RequireTime functions
// without concreting a factory like that rdate.NewTime(...)
// or rdate.RequireTime(...))
func SetDefaultTimeFactory(f *TimeFactory) {
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
	return t.s.String(t.t)
}
