package rdate

import (
	"fmt"
)

type PeriodStringer interface {
	String(from, to Time, sc PeriodShortcut) string
}

var DefaultPeriodStringer = PeriodStringer(&defaultPeriodStringer{})

type defaultPeriodStringer struct{}

func (s *defaultPeriodStringer) String(from, to Time, sc PeriodShortcut) string {
	return fmt.Sprintf("%s — %s", from, to)
}
