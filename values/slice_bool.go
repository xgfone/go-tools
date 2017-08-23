package values

// Bool does the best to convert the value whose index is i to bool.
func (s Slice) Bool(i int) (v bool, ok bool) {
	if len(s) <= i {
		return
	}

	return !IsZero(s[i]), true
}
