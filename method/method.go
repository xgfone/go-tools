// Call the method of a type dynamically.
//
// The constraint will be checked at the runtime, it may panic if CallMustPanic
// is true.
//
package method

import (
	"fmt"
	"reflect"
)

var CallMustPanic bool = true

// The short for HasMethod
func Has(t interface{}, method string) bool {
	return HasMethod(t, method)
}

// The short for GetMethod
func Get(t interface{}, method string) interface{} {
	return GetMethod(t, method)
}

// The short for CallMethod
func Call(t interface{}, method string, args ...interface{}) []interface{} {
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
func GetMethod(t interface{}, method string) interface{} {
	m := getMethod(t, method)
	if !m.IsValid() || m.IsNil() {
		return nil
	}
	return m.Interface()
}

// Call the method, `method`, of `t`, and return the result which that method
// returned.
//
// Exception:
//     (1) It will panic, if `t` does not have the method of `method`,
//         or it's failed to call the method.
//     (2) If `CallMustPanic` is false, it won't panic, but return nil.
func CallMethod(t interface{}, method string, args ...interface{}) []interface{} {
	isNil := false
	defer func() {
		if isNil {
			return
		}
		if err := recover(); err != nil && CallMustPanic {
			panic(err)
		}
	}()
	m := getMethod(t, method)
	if !m.IsValid() || m.IsNil() {
		if CallMustPanic {
			isNil = true
			panic(fmt.Sprintf("Can't find the method: %v", method))
		}
		return nil
	}

	in := []reflect.Value{reflect.ValueOf(t)}
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}
	out := m.Call(in)
	result := make([]interface{}, 0)
	for _, arg := range out {
		result = append(result, arg.Interface())
	}
	return result
}
