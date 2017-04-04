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
)

// RunForever runs for ever.
func RunForever() {
	locked.Lock()
	if shutdowned {
		locked.Unlock()
		panic("The server has been shutdowned")
	}
	locked.Unlock()

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
	shouldShutdown <- true
}

// IsShutdowned returns whether the server has been shutdowned.
func IsShutdowned() (yes bool) {
	locked.Lock()
	yes = shutdowned
	locked.Unlock()
	return
}

// RegisterManager replaces the default lifecycle manager.
// The default manager is the default global manager in the package lifecycle.
func RegisterManager(m *lifecycle.Manager) {
	locked.Lock()
	manager = m
	locked.Unlock()
}
