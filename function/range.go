package function

import (
	"fmt"
)

// Ranges returns a integer range between start and stop, which progressively
// increase or descrease by step.
//
// If step is positive, r[i] = start + step*i when i>0 and r[i]<stop.
//
// If step is negative, r[i] = start + step*i but when i>0 and r[i]>stop.
//
// If step is 0, it will panic.
func Ranges(start, stop, step int) (r []int) {
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

	panic(fmt.Errorf("The step must not be 0"))
}

// Range collects three kinds of the using of Range.
//
//     Range(stop)              ==> Ranges(0, num, 1)
//     Range(start, stop)       ==> Ranges(start, stop, 1)
//     Range(start, stop, step) ==> Ranges(start, stop, step)
//
// Notice: it is equal to range in Python.
func Range(num int, others ...int) []int {
	switch len(others) {
	case 0:
		return Ranges(0, num, 1)
	case 1:
		return Ranges(num, others[0], 1)
	case 2:
		return Ranges(num, others[0], others[1])
	default:
		panic(fmt.Errorf("too many arguments to call Ranges"))
	}
}
