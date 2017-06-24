package values

// Keys returns all the keys.
func (m SMap) Keys() []string {
	r := make([]string, 0)
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Values returns all the values.
func (m SMap) Values() []interface{} {
	r := make([]interface{}, 0)
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// In returns true if k is a key of SMap, or false.
func (m SMap) In(k string) bool {
	_, ok := m[k]
	return ok
}
