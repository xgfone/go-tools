package method

import "reflect"

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
	return getMethod(t, method).Interface()
}

// Call the method, `method`, of `t`, and return the result which that method
// returned. It will panic if `t` does not have the method of `method`.
func CallMethod(t interface{}, method string, args ...interface{}) []interface{} {
	in := []reflect.Value{reflect.ValueOf(t)}
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}
	out := getMethod(t, method).Call(in)
	result := make([]interface{}, 0)
	for _, arg := range out {
		result = append(result, arg.Interface())
	}
	return result
}
