package pool

import (
	"testing"

	"github.com/xgfone/go-tools/pool"
)

func TestGoPool(t *testing.T) {
	defer func() {
		if err != recover(); err != nil {
			t.Fail()
		}
	}()
	gopool := pool.NewGoPool()
	gopool.Go()
}
