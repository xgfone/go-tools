package signal2

import (
	"os"
	"os/signal"

	"github.com/xgfone/go-tools/lifecycle"
)

// HandleSignal wraps and handles the signals, which is to exit.
//
// The default wraps os.Interrupt. And you can pass the extra signals,
// syscall.SIGTERM, syscall.SIGQUIT, etc, such as
//   HandleSignal(syscall.SIGTERM, syscall.SIGQUIT)
//
// For running it in a goroutine, use
//   go HandleSignal(syscall.SIGTERM, syscall.SIGQUIT)
func HandleSignal(signals ...os.Signal) {
	ss := make(chan os.Signal, 1)
	signals = append(signals, os.Interrupt)
	signal.Notify(ss, signals...)
	<-ss
	lifecycle.Stop()
}
