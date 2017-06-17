package server

import (
	"sync"

	"github.com/xgfone/go-tools/lifecycle"
)

var (
	manager        = lifecycle.GetDefaultManager()
	locked         = new(sync.Mutex)
	shutdowned     = false
	shouldShutdown = make(chan bool, 1)
	wait           = lifecycle.NewManager()
)

// RunForever runs for ever.
func RunForever() {
	if IsShutdowned() {
		panic("The server has been shutdowned")
	}

	<-shouldShutdown
	manager.Stop()
}

// Shutdown shutdowns the server gracefully.
func Shutdown() {
	locked.Lock()
	defer locked.Unlock()
	if shutdowned {
		return
	}

	shutdowned = true
	wait.Stop()
	shouldShutdown <- true
}

// IsShutdowned returns whether the server has been shutdowned.
func IsShutdowned() (yes bool) {
	locked.Lock()
	yes = shutdowned
	locked.Unlock()
	return
}

func waitShutdown() {
	in := make(chan interface{})
	out := make(chan interface{})
	wait.RegisterChannel(in, out)
	<-in
	out <- struct{}{}
}

// WaitShutdown will wait until it is shutdowned.
// Return immediately if it has been shutdowned.
func WaitShutdown() {
	if IsShutdowned() {
		return
	}

	waitShutdown()
}

// RegisterManager replaces the default lifecycle manager.
// The default manager is the default global manager in the package lifecycle.
func RegisterManager(m *lifecycle.Manager) {
	locked.Lock()
	manager = m
	locked.Unlock()
}
