// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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
