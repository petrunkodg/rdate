// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rdate

import "time"

type PeriodRule interface {
	Calculate(pivot time.Time, tf TimeFactory) (from, to Time)
	Shortcut() PeriodShortcut
}

type periodRuleThisDay struct{}

func (p *periodRuleThisDay) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfThisDay),
		tf.Require(pivot, TimeEndOfThisDay)
}

func (p *periodRuleThisDay) Shortcut() PeriodShortcut { return PeriodThisDay }

type periodRulePrevDay struct{}

func (p *periodRulePrevDay) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfPrevDay),
		tf.Require(pivot, TimeEndOfPrevDay)
}

func (p *periodRulePrevDay) Shortcut() PeriodShortcut { return PeriodPrevDay }

type periodRulePrevWeek struct{}

func (p *periodRulePrevWeek) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfPrevWeek),
		tf.Require(pivot, TimeEndOfPrevWeek)
}

func (p *periodRulePrevWeek) Shortcut() PeriodShortcut { return PeriodPrevWeek }

type periodRulePrevMonth struct{}

func (p *periodRulePrevMonth) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfPrevMonth),
		tf.Require(pivot, TimeEndOfPrevMonth)
}

func (p *periodRulePrevMonth) Shortcut() PeriodShortcut { return PeriodPrevMonth }

type periodRulePrevQuart struct{}

func (p *periodRulePrevQuart) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfPrevQuart),
		tf.Require(pivot, TimeEndOfPrevQuart)
}

func (p *periodRulePrevQuart) Shortcut() PeriodShortcut { return PeriodPrevQuart }

type periodRulePrevHalfYear struct{}

func (p *periodRulePrevHalfYear) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfPrevHalfYear),
		tf.Require(pivot, TimeEndOfPrevHalfYear)
}

func (p *periodRulePrevHalfYear) Shortcut() PeriodShortcut { return PeriodPrevHalfYear }

type periodRulePrevYear struct{}

func (p *periodRulePrevYear) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfPrevYear),
		tf.Require(pivot, TimeEndOfPrevYear)
}

func (p *periodRulePrevYear) Shortcut() PeriodShortcut { return PeriodPrevYear }

type periodRuleThisWeek struct{}

func (p *periodRuleThisWeek) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfThisWeek),
		tf.Require(pivot, TimeEndOfThisWeek)
}

func (p *periodRuleThisWeek) Shortcut() PeriodShortcut { return PeriodThisWeek }

type periodRuleThisMonth struct{}

func (p *periodRuleThisMonth) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfThisMonth),
		tf.Require(pivot, TimeEndOfThisMonth)
}

func (p *periodRuleThisMonth) Shortcut() PeriodShortcut { return PeriodThisMonth }

type periodRuleThisQuart struct{}

func (p *periodRuleThisQuart) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfThisQuart),
		tf.Require(pivot, TimeEndOfThisQuart)
}

func (p *periodRuleThisQuart) Shortcut() PeriodShortcut { return PeriodThisQuart }

type periodRuleThisHalfYear struct{}

func (p *periodRuleThisHalfYear) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfThisHalfYear),
		tf.Require(pivot, TimeEndOfThisHalfYear)
}

func (p *periodRuleThisHalfYear) Shortcut() PeriodShortcut { return PeriodThisHalfYear }

type periodRuleThisYear struct{}

func (p *periodRuleThisYear) Calculate(pivot time.Time, tf TimeFactory) (from, to Time) {
	return tf.Require(pivot, TimeStartOfThisYear),
		tf.Require(pivot, TimeEndOfThisYear)
}

func (p *periodRuleThisYear) Shortcut() PeriodShortcut { return PeriodThisYear }
