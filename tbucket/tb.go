// The Simple Token Bucket like HTB in Linux TC.
// [Deprecated!] Please use golang.org/x/time/rate.
package tbucket

import "time"

type TokenBucket struct {
	bucket  chan bool
	sleep   time.Duration
	started bool
	stoped  bool
	num     int64
	tick    time.Duration
	cache   uint64
}

// NewTB creates a new token bucket.
// The default size of the token bucket is 1024.
func NewTokenBucket(rate uint64) *TokenBucket {
	t := &TokenBucket{cache: rate}
	return t.SetMinTick(time.Millisecond * 2).SetBucketSize(1024)
}

// Set the size of the token bucket. The default is 1024.
//
// Please set it up according to the real case. If you need the more tokens,
// you maybe set it up to a larger value. But the larger the size is, the more
// the burst quantity.
//
// If the token bucket server has been started, calling this method will panic.
func (t *TokenBucket) SetBucketSize(size uint) *TokenBucket {
	if t.started {
		panic("The token bucket server has been started")
	}
	t.bucket = make(chan bool, size)
	return t
}

// Set the rate to produce the token. The unit is token/s.
//
// Allow that adjust the rate in running. But please don't calling it concurrently.
// If needing to do this, please use the synchrolock, such as sync.Mutex, by yourself.
//
// If rate * tick > time.Second, recommend to use that value, which multiply
// rate by tick is divisibled by time.Second, as the rate. If tick is 10ms and
// rate is greater than 100, for example, please use the multiple of 100 as
// the rate. Of course, you don't have to use the multiple of 100, and can use
// the arbitrary value, but it will be truncated to the multiple of 100.
// Also see SetMinTick(tick).
func (t *TokenBucket) SetRate(rate uint64) *TokenBucket {
	t.num = 1
	t.sleep = time.Second / time.Duration(rate)
	if t.sleep < t.tick {
		t.num = int64(t.tick / t.sleep)
		t.sleep = t.tick
	}
	return t
}

// Set the minimal time granularity. The default is 2ms. Don't suggest to set
// it to a smaller number, unless you known what to happen. You can regard it
// as the clock tick in OS.
//
// When calling this method, it will recalculate the real clock tick and the
// number of the tokens which are produced in one clock tick.
//
// Notice:
// The larger the value of tick is, the higher the load of OS is. Based on a
// simple network rate test, which one token stands for the rate of 1KB/s, 2ms
// is a suitable tick. If it's bad for your case, you can adjust it smaller or
// bigger (recommend), such as 4ms, 8ms, 10ms, which had better be a multiple
// of 2 or 10.
//
// The final real clock tick is decided by both the default tick and the rate.
// See SetRate(rate). You can get the final real clock tick and the number of
// the tokens which are produced in one clock tick. See Audit().
func (t *TokenBucket) SetMinTick(tick time.Duration) *TokenBucket {
	t.tick = tick
	if t.cache > 0 {
		t.SetRate(t.cache)
	}
	return t
}

// Audit returns the real clock tick, and the number of the tokens which are
// produced in one clock tick.
func (t TokenBucket) Audit() (tick time.Duration, num int64) {
	return t.sleep, t.num
}

// Get the token from the bucket.
//
// This method doesn't return a value. That it returns is indicating that you
// have got the token.
//
// If the token bucket server has not been started, calling this method will panic.
func (t *TokenBucket) Get() {
	if !t.started {
		panic("The token bucket server isn't started")
	}
	<-t.bucket
	return
}

// Return true if the token bucket server has been started. Or false.
func (t TokenBucket) IsStart() bool {
	return t.started
}

// Start to produce the token and put it to the bucket. Then you can get
// the token from the bucket by calling t.Get().
//
// If the token bucket server has been started, calling this method will panic.
func (t *TokenBucket) Start() {
	if t.started {
		panic("The token bucket server has been started")
	}

	t.started = true
	go t.start()
}

// Stop the token bucket server. Later you can start it again.
//
// If the token bucket server has not been started, calling this method will panic.
func (t *TokenBucket) Stop() {
	if !t.started {
		panic("The token bucket server isn't started")
	}
	t.started = false

	// In order to let the for loop ends in t.start().
	time.Sleep(t.sleep)
}

func (t *TokenBucket) start() {
	for t.started {
		for i := int64(0); i < t.num; i++ {
			t.bucket <- true
		}

		time.Sleep(t.sleep)
	}
}
