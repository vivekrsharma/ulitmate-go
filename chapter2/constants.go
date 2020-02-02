package main

import "time"

const (
	ui = 12345
	uf = 3.145
)

const (
	ti int     = 12345
	tf float64 = 3.145
)

const answer = 3 * 0.333 // KindFloat = Kindfloat() * KindFloat() promotion of 3 from KindInt to KindFloat
const third = 1 / 3.0    // KindFloat = KindFloat() / KindFloat(); promotion of 1 to KindFloat.
const zero = 1 / 3       // KindInt = KindInt / KindInt; zero retains the KindInt.

const one int8 = 1
const two = 2 * one // two is of type int8; it is promoted by implicit conversion of TypeInt8 * KindInt to TypeInt8

const (
	maxInt              = 93333333333333      // max value of int64
	biggerInt           = 9999999999999999999 // much bigger than int64
	muchbiggerInt int64 = 9999999999999999999 // error; exceeded the size of int64
)

func main() {
	now := time.Now()
	now.Add(5)         // Implicit converstion from KindInt to TypeDuration
	minusFive := -5    // type in64
	now.Add(minusFive) // Error - Cannot use minusFive (type int64) as (type Duration)
}
