package function

// Range returns a integer range between start and stop, which progressively
// increase or descrease by step.
//
// It is equal to range(start, stop, step) in Python.
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

	panic("The step is 0")
}

// RangeWithStep is the closure function for step in Range.
//
// RangeWithStep(1)(start, stop) is equal to range(start, stop) in Python.
func RangeWithStep(step int) func(start, stop int) []int {
	return func(start, stop int) []int {
		return Range(start, stop, step)
	}
}

// RangeByStop is the closure function for start and step in Range.
//
// RangeByStop(0, 1)(stop) is equal to range(stop) in Python.
func RangeByStop(start, step int) func(stop int) []int {
	return func(stop int) []int {
		return Range(start, stop, step)
	}
}
