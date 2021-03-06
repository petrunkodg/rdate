[![GoDoc](https://godoc.org/github.com/petrunkodg/rdate?status.svg)](https://godoc.org/github.com/petrunkodg/rdate)
[![coverage](https://img.shields.io/codecov/c/github/petrunkodg/rdate)](https://codecov.io/gh/petrunkodg/rdate)

# rdate

A golang package which implements a few primitives to work with time and periods conveniently.
It's especially useful for creating reports.

The package provides some presets of time and period calculation. Also it can be extended.

## Getting

```
go get -u github.com/petrunkodg/rdate
```

## Overview

- 25 default rules (presets) of time calculation
- 12 default rules (presets) of period calculation
- You can add new ones or replace any of them
- You can set your own stringer for Time or Period types or decorate the default ones

## Examples

Simple example (Time):

	ts := time.Date(2020, 8, 11, 0, 2, 1, 6, time.UTC)

	t, ok := rdate.NewTime(ts, "start prev week")
	if !ok {
		fmt.Println("the shortcut doesn't exist")
	}

	fmt.Println(t)

	// If you are sure the shortcut exists, you don't need to check ok.
	// you can just use RequireTime.
	t = rdate.RequireTime(ts, rdate.TimeStartOfPrevMonth)
	fmt.Println(t)

	t = rdate.RequireTime(ts, rdate.TimeEndOfPrevYear)
	fmt.Println(t)

	// Output:
	// 2020-08-03 00:00:00
	// 2020-07-01 00:00:00
	// 2019-12-31 23:59:59


Simple example (Period):

	ts := time.Date(2020, 8, 11, 0, 2, 1, 6, time.UTC)

	p, ok := rdate.NewPeriod(ts, "prev week")
	if !ok {
		fmt.Println("the shortcut doesn't exist")
	}
	fmt.Println(p)

	p = rdate.RequirePeriod(ts, rdate.PeriodPrevQuart)
	fmt.Println(p)

	p = rdate.RequirePeriod(ts, rdate.PeriodThisMonth)
	fmt.Println(p)

	// Output:
	// 2020-08-03 00:00:00 — 2020-08-09 23:59:59
	// 2020-04-01 00:00:00 — 2020-06-30 23:59:59
	// 2020-08-01 00:00:00 — 2020-08-31 23:59:59

For more information and examples see [godoc](https://godoc.org/github.com/petrunkodg/rdate).


## Contact

Danila Petrunko <petrunkodg@gmail.com>

## License

Source code is available under the MIT [License](/LICENSE).