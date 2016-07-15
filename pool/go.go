package pool

import (
	"errors"
	"reflect"
)

var (
	TOTALGO uint = 9900

	MaxGoroutineError = errors.New("More than the goroutine")
	NotFuncError      = errors.New("The first argument is not the function")
	ArgsNumError      = errors.New("The number of the argument is incorrect")
	ArgsTypeError     = errors.New("The type of the argument is incorrect")
)

type GoPool struct {
	num int
}

func NewGoPool() *GoPool {
	return &GoPool{}
}

func (p GoPool) GetNum() int {
	return p.num
}

func (p *GoPool) Go(f, args ...interface{}) error {
	if p.num > TOTALGO {
		return MaxGoroutineError
	}

	vf := reflect.ValueOf(f)
	if vf.Kind() != reflect.Func {
		return NotFuncError
	}

	tf := vf.Type()
	_len := len(args)
	if tf.NumIn() != _len {
		return ArgsNumError
	}

	_args := make([]reflect.Value, _len)
	for i := 0; i < _len; i++ {
		if tf.In(i).Kind() != reflect.TypeOf(args[i]).Kind() {
			return ArgsTypeError
		}
		_args[i] = reflect.ValueOf(args[i])
	}

	go vf.Call(_args)

	return nil
}

func (p *GoPool) run() {

}
