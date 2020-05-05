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
	// ErrNotFunc is returned when the callee is not a function.
	ErrNotFunc = errors.New("The first argument is not the function")

	// ErrArgsNum is returned when the number of the arguments is incorrect
	// when calling the callee.
	ErrArgsNum = errors.New("The number of the argument is incorrect")

	// ErrArgsType is returned when the type of the arguments is incorrect
	// when calling the callee.
	ErrArgsType = errors.New("The type of the argument is incorrect")
)

// Valid valids whether the callee is a function, and the number and the type
// of the arguments is correct, then return the valid function and arguments.
func Valid(callee interface{}, args ...interface{}) (vf reflect.Value,
	vargs []reflect.Value, err error) {
	vf = reflect.ValueOf(callee)
	if vf.Kind() != reflect.Func {
		return reflect.ValueOf(nil), nil, ErrNotFunc
	}

	tf := vf.Type()
	_len := len(args)
	if tf.NumIn() != _len {
		return reflect.ValueOf(nil), nil, ErrArgsNum
	}

	vargs = make([]reflect.Value, _len)
	for i := 0; i < _len; i++ {
		typ := tf.In(i).Kind()
		if (typ != reflect.Interface) && (typ != reflect.TypeOf(args[i]).Kind()) {
			return reflect.ValueOf(nil), nil, ErrArgsType
		}
		vargs[i] = reflect.ValueOf(args[i])
	}
	return vf, vargs, nil
}

// Call calls a function dynamically.
func Call(f interface{}, args ...interface{}) (results []interface{}, err error) {
	vf, vargs, _err := Valid(f, args...)
	if _err != nil {
		return nil, _err
	}
	ret := vf.Call(vargs)
	_len := len(ret)
	results = make([]interface{}, _len)
	for i := 0; i < _len; i++ {
		results[i] = ret[i].Interface()
	}
	return
}
