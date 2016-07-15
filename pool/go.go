// Goroutine pool, the wrapper of the keyword, go.
//
// If the number of the running goroutine exceeds the maximal, it refuses to
// execute the goroutine.
//
package pool

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/xgfone/go-tools/function"
)

const (
	maxLimit = 9999
)

var (
	goLimit int = 9000

	MaxGoroutineError = errors.New("More than the goroutine")
)

// Set the maximal limit num of goroutine, and return the old.
//
// Return -1 and don't set it if num is less than 1, or too big,
// such as more than the thread limit, see
//
//     (1) https://golang.org/pkg/runtime/debug/#SetMaxThreads
//     (2) https://github.com/golang/go/issues/4056
//
// At present, the maximal limit is 9999.
func SetMaxLimit(num int) (old int) {
	if num < 1 || num > maxLimit {
		return -1
	}
	old = goLimit
	goLimit = num
	return
}

type GoPool struct {
	sync.Mutex
	num int
}

// Get a goroutine pool, also get it by &GoPool{} directly.
func NewGoPool() *GoPool {
	return &GoPool{}
}

// Get the number of all the current running goroutines.
func (p *GoPool) GetNum() int {
	p.Lock()
	defer p.Unlock()
	return p.num
}

func (p *GoPool) del() {
	p.Lock()
	defer p.Unlock()
	p.num -= 1
}

// Call the function with the arguments in a new goroutine.
//
// Return an error when the first argument f is not a function, the arguments
// is incorrect, or it can not start a new goroutine. Return nil when starting
// a new goroutine. Even though the arguments is correct, it don't guarantee '
// that the function can be called successfully.
//
// It is the same as the keyword, go, which discards the returned values.
// But the difference of both is that this method will capture the panic which
// occurs in the called function and don't allow it to panic.
func (p *GoPool) Go(f interface{}, args ...interface{}) error {
	p.Lock()
	defer p.Unlock()
	if p.num > goLimit {
		return MaxGoroutineError
	}
	p.num += 1

	vf, vargs, err := function.Valid(f, args...)
	if err != nil {
		return err
	}

	go p.run(vf, vargs)

	return nil
}

func (p *GoPool) run(f reflect.Value, args []reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[Goroutine] Failed to call the %v(%v): %v\n", f.Type().Name(), err)
		}
	}()

	defer p.del()

	f.Call(args)
}
