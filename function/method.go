// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	return b
}

func getMethod(t interface{}, method string) reflect.Value {
	if m, b := reflect.TypeOf(t).MethodByName(method); b {
		return m.Func
	}
	return reflect.ValueOf(nil)
}

// GetMethod returns the method, `method`, of `t`. If not, return nil.
//
// Notice: The first argument of the returned function is the receiver. That's,
// when calling the function, you must pass the receiver as the first argument
// of that, but, which the passed receiver needs not be identical to t.
func GetMethod(t interface{}, method string) interface{} {
	if m := getMethod(t, method); m.IsValid() {
		return m.Interface()
	}
	return nil
}

// CallMethod calls the method 'method' of 't', and return (ReturnedValue, nil)
// if calling successfully, which ReturnedValue is the result which that method
// returned. Or return (nil, Error).
func CallMethod(t interface{}, method string, args ...interface{}) (
	results []interface{}, err error) {
	if m := GetMethod(t, method); m != nil {
		_args := make([]interface{}, len(args)+1)
		_args[0] = t
		copy(_args[1:], args)
		results, err = Call(m, _args...)
	}
	return nil, ErrNotHaveMethod
}
