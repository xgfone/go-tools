package server

import (
	"fmt"
	"log"
	"os"
)

var (
	// If true, output the verbose information.(deprecated)
	Debug   bool
	_logger *logger
)

func init() {
	_logger = &logger{_logger: log.New(os.Stderr, "[TCP Server] ", log.LstdFlags)}
}

type logger struct {
	_logger *log.Logger
}

func (l logger) Debug(format string, args ...interface{}) {
	l.Output(10, format, args...)
}

func (l logger) Output(level int, format string, args ...interface{}) {
	var prefix string
	if level <= 10 {
		prefix = "Debug"
	} else if level <= 20 {
		prefix = "Info"
	} else if level <= 30 {
		prefix = "Warning"
	} else if level <= 40 {
		prefix = "Error"
	}

	f := fmt.Sprintf("[%v] %v\n", prefix, format)
	l._logger.Printf(f, args...)
}

func (l logger) Info(format string, args ...interface{}) {
	l.Output(20, format, args...)
}

func (l logger) Warn(format string, args ...interface{}) {
	l.Output(30, format, args...)
}

func (l logger) Error(format string, args ...interface{}) {
	l.Output(40, format, args...)
}

func SetLogger(logger *log.Logger) {
	_logger._logger = logger
}

func GetLogger() *log.Logger {
	return _logger._logger
}
