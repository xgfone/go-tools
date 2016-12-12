package values

func (m SMap) Keys() []string {
	r := make([]string, 0)
	for k, _ := range m {
		r = append(r, k)
	}
	return r
}

func (m SMap) Values() []interface{} {
	r := make([]interface{}, 0)
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
