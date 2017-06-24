package server

import (
	"fmt"
	"log"
	"os"
)

var (
	// NotOutputLog is on/off. If true, don't output the any log.
	NotOutputLog bool

	// NotOutputNewLine is on/off. If true, don't append a new line to the log.
	NotOutputNewLine bool

	_logger logger
)

func init() {
	_logger = logger{_logger: log.New(os.Stderr, "[TCP Server] ", log.LstdFlags)}
}

// Logger is a interface to implement a logger.
type Logger interface {
	Printf(format string, v ...interface{})
}

type logger struct {
	_logger Logger
}

func (l logger) Printf(format string, v ...interface{}) {
	if !NotOutputLog {
		if !NotOutputNewLine {
			format = fmt.Sprintln(format)
		}
		l._logger.Printf(format, v...)
	}
}

// SetLogger replaces the default logger.
func SetLogger(logger Logger) {
	_logger._logger = logger
}

// GetLogger returns the current logger.
func GetLogger() Logger {
	return _logger._logger
}
