// Copyright © 2020 Danila Petrunko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rdate

import (
	"time"
)

type TimeStringer interface {
	String(t time.Time) string
}

var DefaultTimeStringer = TimeStringer(&defaultTimeStringer{})

type defaultTimeStringer struct{}

func (s *defaultTimeStringer) String(t time.Time) string {
	const format = "2006-01-02 15:04:05"

	return t.Format(format)
}
