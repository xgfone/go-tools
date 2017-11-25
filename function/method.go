package function

import (
	"errors"
	"reflect"
)

var (
	// ErrNotHaveMethod is returned when a certain type doesn't have the method.
	ErrNotHaveMethod = errors.New("Don't have the method")
)

// HasMethod returns true if `t` has the method of `method`.
func HasMethod(t interface{}, method string) bool {
	_, b := reflect.TypeOf(t).MethodByName(method)
	if b {
		return true
	}
	return false
}

func getMethod(t interface{}, method string) reflect.Value {
	m, b := reflect.TypeOf(t).MethodByName(method)
	if !b {
		return reflect.ValueOf(nil)
	}
	return m.Func
}

// GetMethod returns the method, `method`, of `t`. If not, return nil.
//
// Notice: The first argument of the returned function is the receiver. That's,
// when calling the function, you must pass the receiver as the first argument
// of that, but, which the passed receiver needs not be identical to t.
func GetMethod(t interface{}, method string) interface{} {
	m := getMethod(t, method)
	if !m.IsValid() || m.IsNil() {
		return nil
	}
	return m.Interface()
}

// CallMethod calls the method 'method' of 't', and return (ReturnedValue, nil)
// if calling successfully, which ReturnedValue is the result which that method
// returned. Or return (nil, Error).
func CallMethod(t interface{}, method string, args ...interface{}) ([]interface{}, error) {
	m := GetMethod(t, method)
	if m == nil {
		return nil, ErrNotHaveMethod
	}
	_args := make([]interface{}, len(args)+1)
	_args[0] = t
	copy(_args[1:], args)
	return Call(m, _args...)
}
