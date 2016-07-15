// Call a function dynamically.
//
// The constraint will be checked at the runtime.
//
package function

import (
	"errors"
	"reflect"
)

var (
	TOTALGO uint = 9900

	NotFuncError  = errors.New("The first argument is not the function")
	ArgsNumError  = errors.New("The number of the argument is incorrect")
	ArgsTypeError = errors.New("The type of the argument is incorrect")
)

func Call(f interface{}, args ...interface{}) ([]interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			return nil, errors.New(err)
		}
	}()

	vf := reflect.ValueOf(f)
	if vf.Kind() != reflect.Func {
		return nil, NotFuncError
	}

	tf := vf.Type()
	_len := len(args)
	if tf.NumIn() != _len {
		return nil, ArgsNumError
	}

	_args := make([]reflect.Value, _len)
	for i := 0; i < _len; i++ {
		if tf.In(i).Kind() != reflect.TypeOf(args[i]).Kind() {
			return nil, ArgsTypeError
		}
		_args[i] = reflect.ValueOf(args[i])
	}

	ret := vf.Call(_args)
	_len = len(ret)
	results := make([]interface{}, _len)
	for i := 0; i < _len; i++ {
		results[i] = ret[i].Interface()
	}

	return results, nil
}
