package function

import (
	"fmt"
)

// Range returns a integer range between start and stop, which progressively
// increase or descrease by step.
//
// If step is positive, r[i] = start + step*i when i>0 and r[i]<stop.
//
// If step is negative, r[i] = start + step*i but when i>0 and r[i]>stop.
//
// If step is 0, it will panic.
func Range(start, stop, step int) (r []int) {
	if step > 0 {
		for start < stop {
			r = append(r, start)
			start += step
		}
		return
	} else if step < 0 {
		for start > stop {
			r = append(r, start)
			start += step
		}
		return
	}

	panic(fmt.Errorf("The step is not 0"))
}

// Ranges collects three kinds of the using of Range.
//
//     Ranges(stop)              ==> Range(0, num, 1)
//     Ranges(start, stop)       ==> Range(start, stop, 1)
//     Ranges(start, stop, step) ==> Range(start, stop, step)
//
// Notice: it is equal to range in Python.
func Ranges(num int, others ...int) []int {
	switch len(others) {
	case 0:
		return Range(0, num, 1)
	case 1:
		return Range(num, others[0], 1)
	case 2:
		return Range(num, others[0], others[1])
	default:
		panic(fmt.Errorf("too many arguments to call Ranges"))
	}
}

// RangeStepOne is equal to Range(start, stop, 1). [DEPRECATED]
//
// That's, range(start, stop) in Python.
func RangeStepOne(start, stop int) []int {
	return Range(start, stop, 1)
}

// RangeStop is equal to Range(0, stop, 1). [DEPRECATED]
//
// That's, range(stop) in Python.
func RangeStop(stop int) []int {
	return Range(0, stop, 1)
}

// RangeWithStep is the closure function for step in Range. [DEPRECATED]
//
// RangeWithStep(1)(start, stop) is equal to range(start, stop) in Python.
func RangeWithStep(step int) func(start, stop int) []int {
	return func(start, stop int) []int {
		return Range(start, stop, step)
	}
}

// RangeByStop is the closure function for start and step in Range. [DEPRECATED]
//
// RangeByStop(0, 1)(stop) is equal to range(stop) in Python.
func RangeByStop(start, step int) func(stop int) []int {
	return func(stop int) []int {
		return Range(start, stop, step)
	}
}
