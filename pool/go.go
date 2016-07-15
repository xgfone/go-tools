package pool

import "reflect"

var (
	TOTALGO uint = 9900
)

type GoPool struct {
	num int
}

func (p GoPool) GetNum() int {
	return p.num
}

func (p *GoPool) Go(f, args ...interface{}) bool {
	if p.num > TOTALGO {
		return false
	}

	vf := reflect.ValueOf(f)
	if vf.Kind() != reflect.Func {
		return false
	}
	_len := len(args)
	_args := make([]reflect.Value, _len)

	tf := vf.Type()

	return true
}
