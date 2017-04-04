package server

import (
	"github.com/xgfone/go-tools/atomics"
	"github.com/xgfone/go-tools/lifecycle"
)

var (
	manager        = lifecycle.GetDefaultManager()
	shutdowned     = atomics.NewBool()
	shouldShutdown = make(chan bool, 1)
)

// RunForever runs for ever.
func RunForever() {
	if shutdowned.Get() {
		panic("The server has been shutdowned")
	}
	<-shouldShutdown
	manager.Stop()
}

// Shutdown shutdowns the server gracefully.
func Shutdown() {
	shutdowned.SetTrue()
	shouldShutdown <- true
}

// IsShutdowned returns whether the server has been shutdowned.
func IsShutdowned() bool {
	return shutdowned.Get()
}

// RegisterManager replaces the default lifecycle manager.
// The default manager is the default global manager in the package lifecycle.
func RegisterManager(m *lifecycle.Manager) {
	manager = m
}
