package tbucket_test

import (
	"fmt"

	"github.com/xgfone/go-tools/tbucket"
)

func ExampleTokenBucket() {
	// Get a token from the bucket per second.
	tb := tbucket.NewTokenBucket(1)
	//tb.SetMinTick(time.Millisecond).SetRate(150)
	tb.Start()
	go func(tb *tbucket.TokenBucket) {
		for {
			tb.Get()
			fmt.Println("Get a token") // You can see it prints once per second.
		}
	}(tb)

	for {
	}
}
