package pool_test

import (
	"fmt"
	"time"

	"github.com/xgfone/go-tools/pool"
)

func ExampleGoPool_Go() {
	gopool := pool.NewGoPool()

	f1 := func(p *pool.GoPool) {
		fmt.Println("IN", p.GetNum()) // print: IN 1
	}
	if err := gopool.Go(f1, gopool); err != nil {
		fmt.Println(err)
	} else {
		time.Sleep(time.Millisecond * 10)   // For excuting the f1 firstly
		fmt.Println("OUT", gopool.GetNum()) // print: OUT 0
	}

	f2 := func(i int, j int) {
		// The result is output in other goroutine.
		// So it is not the output of the example.
		fmt.Println("ADD", i+j) // print: ADD 3
	}
	if err := gopool.Go(f2, 1, 2); err != nil {
		fmt.Println(err)
	} else {
		time.Sleep(time.Millisecond * 10) // For excuting the f2 firstly
		fmt.Println("OK")
	}

	// Output:
	// IN 1
	// OUT 0
	// ADD 3
	// OK
}
