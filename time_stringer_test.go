package rdate_test

import (
	"testing"
	"time"

	"gitlab.com/petrunkodg/rdate"
)

func TestDefaultTimeStringer(t *testing.T) {
	ts := time.Date(2019, 12, 11, 0, 2, 1, 6, time.UTC)

	expected := "2019-12-11 00:02:01"

	actual := rdate.DefaultTimeStringer.String(ts)
	if actual != expected {
		t.Errorf("expected: '%s', but actual: '%s'", expected, actual)
	}
}
