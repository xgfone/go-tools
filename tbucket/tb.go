// The Simple Token Bucket like HTB in Linux TC.
package tbucket

import "time"

type TokenBucket struct {
	bucket  chan bool
	sleep   time.Duration
	started bool
	stoped  bool
}

// NewTB creates a new token bucket.
// The default size of the token bucket is 1024.
func NewTokenBucket(rate uint64) *TokenBucket {
	t := &TokenBucket{}
	return t.SetRate(rate).SetBucketSize(1024)
}

func (t TokenBucket) calcSleep(rate uint64) time.Duration {
	return time.Second / time.Duration(rate)
	//return time.Duration(uint64(time.Second) / rate)
}

// Set the size of the token bucket.
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
// Allow that adjust the rate in running.
func (t *TokenBucket) SetRate(rate uint64) *TokenBucket {
	t.sleep = t.calcSleep(rate)
	return t
}

// Get the token from the bucket.
//
// This method isn't the returned value. That it returns is indicating that you
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

// Start to produce the token and put it to the bucket. Then you can get
// the token from the bucket by calling t.Get().
//
// If the token bucket server has been started, calling this method will panic.
func (t *TokenBucket) Start() {
	if t.started {
		panic("The token bucket server has been started")
	}

	go t.start()
	t.started = true
	t.stoped = false
}

// Stop the token bucket server. Later you can start it again.
//
// If the token bucket server has not been started, calling this method will panic.
func (t *TokenBucket) Stop() {
	if !t.started {
		panic("The token bucket server isn't started")
	}
	t.stoped = true
	t.started = false

	// In order to let the for loop ends in t.start().
	time.Sleep(t.sleep)
}

func (t *TokenBucket) start() {
	for !t.stoped {
		t.bucket <- true
		time.Sleep(t.sleep)
	}
}
