package values

// Interface returns the interface of the ith value.
func (s Slice) Interface(i int) (v interface{}, ok bool) {
	if len(s) <= i {
		return nil, false
	}

	return s[i], true
}

// Slice does the best to convert the value whose index is i to Slice.
func (s Slice) Slice(i int) (v Slice, ok bool) {
	if len(s) <= i {
		return nil, false
	}
	return toSlice(s[i])
}

// SMap does the best to convert the value whose index is i to SMap.
func (s Slice) SMap(i int) (v SMap, ok bool) {
	if len(s) <= i {
		return nil, false
	}
	return toSMap(s[i])
}
