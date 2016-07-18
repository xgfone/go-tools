// The Simple Token Bucket like HTB in Linux TC.
package tb

import "time"

type TB struct {
	bucket  chan bool
	sleep   time.Duration
	started bool
	stoped  bool
}

// NewTB creates a new token bucket.
// The default size of the token bucket is 1024.
func NewTB(rate uint64) *TB {
	t := &TB{}
	return t.SetRate(rate).SetBucketSize(1024)
}

func (t TB) calcSleep(rate uint64) time.Duration {
	return time.Second / time.Duration(rate)
	//return time.Duration(uint64(time.Second) / rate)
}

// Set the size of the token bucket.
//
// If the token bucket server has been started, calling this method will panic.
func (t *TB) SetBucketSize(size uint) *TB {
	if t.started {
		panic("The token bucket server has been started")
	}
	t.bucket = make(chan bool, size)
	return t
}

// Set the rate to produce the token. The unit is token/s.
//
// Allow that adjust the rate in running.
func (t *TB) SetRate(rate uint64) *TB {
	t.sleep = t.calcSleep(rate)
	return t
}

// Get the token from the bucket.
//
// This method isn't the returned value. That it returns is indicating that you
// have got the token.
//
// If the token bucket server has not been started, calling this method will panic.
func (t *TB) Get() {
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
func (t *TB) Start() {
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
func (t *TB) Stop() {
	if !t.started {
		panic("The token bucket server isn't started")
	}
	t.stoped = true
	t.started = false

	// In order to let the for loop ends in t.start().
	time.Sleep(t.sleep)
}

func (t *TB) start() {
	for !t.stoped {
		t.bucket <- true
		time.Sleep(t.sleep)
	}
}
