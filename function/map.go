package function

// InMap returns true if the key exists.
func InMap(m map[string]interface{}, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

// InMapString returns true if the key exists.
func InMapString(m map[string]string, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}
