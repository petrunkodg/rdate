/*
Package rdate implements a few primitives to work with dates or times conveniently.
It's especially useful for creating reports.

If you need some new rules for time or period calculation,
you can just extend time or period factories for that.

Also if you need a custom String implementation of Time or Period types,
you can replace the default ones to yours by writing PeriodStringer or
TimeStringer implementation.
*/
package rdate
