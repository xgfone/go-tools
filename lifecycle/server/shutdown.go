package server

import "github.com/xgfone/go-tools/atomics"

var (
	shutdowned     = atomics.NewBool()
	shouldShutdown = make(chan bool, 1)
)

// RunForever runs for ever.
func RunForever() {
	if shutdowned.Get() {
		panic("The server has been shutdowned")
	}
	<-shouldShutdown
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
