// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rdate_test

import (
	"fmt"
	"time"

	"github.com/petrunkodg/rdate"
)

type prevDecadePeriodRule struct{}

func (p *prevDecadePeriodRule) Calculate(pivot time.Time, tf *rdate.TimeFactory) (from, to rdate.Time) {
	offset := pivot.Year() % 10
	f := pivot.AddDate(-offset, 0, 0)

	return tf.Require(f.AddDate(-10, 0, 0), rdate.TimeStartOfThisYear),
		tf.Require(f, rdate.TimeEndOfThisYear)
}

func (p *prevDecadePeriodRule) Shortcut() rdate.PeriodShortcut { return "prev decade" }

func ExamplePeriodFactory_Extend() {
	pivot := time.Date(1998, 3, 1, 0, 2, 1, 6, time.UTC)

	f := rdate.NewPeriodFactory()

	p, ok := f.Make(pivot, "prev decade")
	if !ok {
		fmt.Println("'prev decade' shortcut is not implemented")
	}
	// 'prev decade' shortcut is not implemented

	f.Extend([]rdate.PeriodRule{&prevDecadePeriodRule{}})

	p, ok = f.Make(pivot, "prev decade")
	if !ok {
		fmt.Println("'prev decade' shortcut is not implemented")
	}

	fmt.Println(p)
	// 1980-01-01 00:00:00 — 1990-12-31 23:59:59
	fmt.Println(p.From().Time())
	// 1980-01-01 00:00:00 +0000 UTC
	fmt.Println(p.To().Time())
	// 1990-12-31 23:59:59.999999999 +0000 UTC

	// Output:
	// 'prev decade' shortcut is not implemented
	// 1980-01-01 00:00:00 — 1990-12-31 23:59:59
	// 1980-01-01 00:00:00 +0000 UTC
	// 1990-12-31 23:59:59.999999999 +0000 UTC
}
