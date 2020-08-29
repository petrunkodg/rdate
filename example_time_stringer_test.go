// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rdate_test

import (
	"fmt"
	"time"

	"github.com/petrunkodg/rdate"
)

type customTimeStringer struct{}

func (s *customTimeStringer) String(t time.Time) string {
	if t.Day() == 13 && t.Month() == time.December {
		return "Happy birthday, Daniel!"
	}

	return t.Format("02 Jan 06 15:04 MST")
}

func ExampleTimeStringer_replacing() {
	tf := rdate.NewTimeFactory()
	tf.SetStringer(&customTimeStringer{})

	d := tf.Require(time.Date(2020, 8, 11, 0, 2, 1, 6, time.UTC), rdate.TimeAsIs)
	fmt.Println(d)
	// 11 Aug 20 00:02 UTC

	d = tf.Require(time.Date(2020, 12, 13, 0, 2, 1, 6, time.UTC), rdate.TimeAsIs)
	fmt.Println(d)
	// Happy birthday, Daniel!

	d = tf.Require(time.Date(1999, 2, 11, 0, 2, 1, 6, time.UTC), rdate.TimeAsIs)
	fmt.Println(d)
	// 11 Feb 99 00:02 UTC

	// Output:
	// 11 Aug 20 00:02 UTC
	// Happy birthday, Daniel!
	// 11 Feb 99 00:02 UTC
}
