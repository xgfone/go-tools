// Package signal2 is the supplement of the standard library of `signal`,
// such as `HandleSignal`.
package signal2

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/xgfone/go-tools/lifecycle"
)

// DefaultSignals is the default signals to be handled.
var DefaultSignals = []os.Signal{
	syscall.SIGTERM,
	syscall.SIGQUIT,
	syscall.SIGABRT,
	syscall.SIGINT,
}

// HandleSignal is the same as HandleSignalWithLifecycle, but using the global
// default lifecycle manager.
//
// It's equal to
//   HandleSignalWithLifecycle(lifecycle.GetDefaultManager(), signals...)
//
// Notice: If the signals are empty, it will be equal to
//   HandleSignal(DefaultSignals...)
func HandleSignal(signals ...os.Signal) {
	if len(signals) == 0 {
		signals = DefaultSignals
	}
	HandleSignalWithLifecycle(lifecycle.GetDefaultManager(), signals...)
}

// HandleSignalWithLifecycle wraps and handles the signals.
//
// The default wraps os.Interrupt. And you can pass the extra signals,
// syscall.SIGTERM, syscall.SIGQUIT, etc, such as
//   m := lifecycle.GetDefaultManager()
//   HandleSignalWithLifecycle(m, syscall.SIGTERM, syscall.SIGQUIT)
//
// For running it in a goroutine, use
//   go HandleSignalWithLifecycle(m, syscall.SIGTERM, syscall.SIGQUIT)
func HandleSignalWithLifecycle(m *lifecycle.Manager, signals ...os.Signal) {
	HandleSignalWithFunc(func() { m.Stop() }, os.Interrupt, signals...)
}

// HandleSignalWithFunc calls the function f when the signals are received.
func HandleSignalWithFunc(f func(), sig os.Signal, signals ...os.Signal) {
	ss := make(chan os.Signal, 1)
	signals = append(signals, sig)
	signal.Notify(ss, signals...)
	for {
		<-ss
		f()
	}
}
