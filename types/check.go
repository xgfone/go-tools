package types

import (
	"fmt"
)

// CheckMapType checks the type of m[k] is t, and return the value if yes,
// or return an error if not.
//
// If m is nil, return an error.
//
// The function uses VerifyType to verify the type, that's, VerifyType(m[k], t),
// so for t, see VerifyType.
func CheckMapType(m map[string]interface{}, k, t string) (interface{}, error) {
	if m == nil {
		return nil, fmt.Errorf("the map is nil")
	}
	value := m[k]
	if value == nil {
		return nil, fmt.Errorf("the value of the key[%s] in map is nil", k)
	}
	if !VerifyType(value, t) {
		return nil, fmt.Errorf("the value of the key[%s] in map is not %s",
			k, stype2type[k])
	}
	return value, nil
}
