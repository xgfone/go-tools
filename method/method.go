// Call the method of a type dynamically.
//
// The constraint will be checked at the runtime.
//
package method

import (
	"errors"
	"reflect"

	"github.com/xgfone/go-tools/function"
)

var (
	NotHaveMethod = errors.New("Don't have the method")
)

// The short for HasMethod
func Has(t interface{}, method string) bool {
	return HasMethod(t, method)
}

// The short for GetMethod
func Get(t interface{}, method string) interface{} {
	return GetMethod(t, method)
}

// The short for CallMethod
func Call(t interface{}, method string, args ...interface{}) ([]interface{}, error) {
	return CallMethod(t, method, args...)
}

// Return true if `t` has the method of `method`.
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

// Return the method, `method`, of `t`. If not, return nil.
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

// Call the method 'method' of 't', and return (ReturnedValue, nil) if calling
// successfully, which ReturnedValue is the result which that method returned.
// Or return (nil, Error).
func CallMethod(t interface{}, method string, args ...interface{}) ([]interface{}, error) {
	if m := GetMethod(t, method); m == nil {
		return nil, NotHaveMethod
	} else {
		_args := make([]interface{}, len(args)+1)
		_args[0] = t
		copy(_args[1:], args)
		return function.Call(m, _args...)
	}
}
