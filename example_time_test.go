// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rdate_test

import (
	"fmt"
	"time"

	"github.com/petrunkodg/rdate"
)

type myBirthdayTimeRule struct{}

func (r *myBirthdayTimeRule) Calculate(pivot time.Time) time.Time {
	return time.Date(pivot.Year(), 12, 13, 0, 0, 0, 0, pivot.Location())
}

func (r *myBirthdayTimeRule) Shortcut() rdate.TimeShortcut {
	return "my birthday this year"
}

func ExampleTimeFactory_extend() {
	pivot := time.Date(2004, 3, 1, 0, 2, 1, 6, time.UTC)

	f := rdate.NewTimeFactory()

	d, ok := f.Make(pivot, "my birthday this year")
	if !ok {
		fmt.Println("'my birthday this year' shortcut is not implemented")
	}
	// 'my birthday this year' shortcut is not implemented

	f.Extend([]rdate.TimeRule{&myBirthdayTimeRule{}})

	d, ok = f.Make(pivot, "my birthday this year")
	if !ok {
		fmt.Println("'my birthday this year' shortcut is not implemented")
	}

	fmt.Println(d)
	// 2004-12-13 00:00:00
	fmt.Println(d.Time())
	// 2004-12-13 00:00:00 +0000 UTC

	// Output:
	// 'my birthday this year' shortcut is not implemented
	// 2004-12-13 00:00:00
	// 2004-12-13 00:00:00 +0000 UTC
}
