// Call a function dynamically.
//
// The constraint will be checked at the runtime.
//
package function

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	NotFuncError  = errors.New("The first argument is not the function")
	ArgsNumError  = errors.New("The number of the argument is incorrect")
	ArgsTypeError = errors.New("The type of the argument is incorrect")
)

func Valid(f interface{}, args ...interface{}) (vf reflect.Value, vargs []reflect.Value, err error) {
	vf = reflect.ValueOf(f)
	if vf.Kind() != reflect.Func {
		return reflect.ValueOf(nil), nil, NotFuncError
	}

	tf := vf.Type()
	_len := len(args)
	if tf.NumIn() != _len {
		return reflect.ValueOf(nil), nil, ArgsNumError
	}

	vargs = make([]reflect.Value, _len)
	for i := 0; i < _len; i++ {
		typ := tf.In(i).Kind()
		if (typ != reflect.Interface) && (typ != reflect.TypeOf(args[i]).Kind()) {
			return reflect.ValueOf(nil), nil, ArgsTypeError
		}
		vargs[i] = reflect.ValueOf(args[i])
	}
	return vf, vargs, nil
}

func Call(f interface{}, args ...interface{}) (results []interface{}, err error) {
	defer func() {
		if _err := recover(); _err != nil {
			err = errors.New(fmt.Sprintf("%v", _err))
			results = nil
		}
	}()

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
