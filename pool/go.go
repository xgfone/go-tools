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
	"sync/atomic"

	"github.com/xgfone/go-tools/function"
)

var (
	MaxGoroutineError = errors.New("More than the goroutine")

	total gonum
)

// Get the number of all the goroutines that are managered by this package.
//
// If you want to known the total number of all the goroutines in the current
// process, please use runtime.NumGoroutine().
func GetTotalNum() uint32 {
	return total.get()
}

type gonum uint32

func (g gonum) get() uint32 {
	return atomic.LoadUint32((*uint32)(&g))
}

func (g gonum) add() {
	atomic.AddUint32((*uint32)(&g), 1)
}

func (g gonum) del() {
	atomic.AddUint32((*uint32)(&g), ^uint32(0))
}

type GoPool struct {
	sync.Mutex
	num   uint
	total uint
}

// Get a goroutine pool, also get it by &GoPool{} directly.
//
// The default doesn't limit the number of goroutine. If it's not what you want,
// you can set up the limit by GoPool.SetMaxLimit(n).
func NewGoPool() *GoPool {
	return &GoPool{}
}

// Set the maximal limit num of goroutine, and return the old.
//
// If the num is 0, it won't limit the number of goroutine.
//
// Don't suggest to set it to too large number. See
//
//     (1) https://golang.org/pkg/runtime/debug/#SetMaxThreads
//     (2) https://github.com/golang/go/issues/4056
//
func (p *GoPool) SetMaxLimit(num uint) (old uint) {
	p.Lock()
	defer p.Unlock()

	old = p.total
	p.total = num
	return
}

// Get the number of all the current running goroutines.
func (p *GoPool) GetNum() uint {
	p.Lock()
	defer p.Unlock()
	return p.num
}

func (p *GoPool) reduce() {
	p.Lock()
	defer p.Unlock()
	p.num -= 1
	total.del()
}

func (p GoPool) test() bool {
	if p.total == 0 {
		return true
	} else if p.num < p.total {
		return true
	} else {
		return false
	}
}

// Call the function with the arguments in a new goroutine.
//
// Return an error when the first argument f is not a function, the arguments
// is incorrect, or it can not start a new goroutine. Return nil when starting
// a new goroutine. Even though the arguments is correct, it don't guarantee
// that the function can be called successfully.
//
// It is the same as the keyword, go, which discards the returned values.
// But the difference of both is that this method will capture the panic which
// occurs in the called function and don't allow it to panic.
func (p *GoPool) Go(f interface{}, args ...interface{}) error {
	p.Lock()
	defer p.Unlock()

	if !p.test() {
		return MaxGoroutineError
	}
	p.num += 1
	total.add()

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

	defer p.reduce()

	f.Call(args)
}
