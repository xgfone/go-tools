package server

import (
	"log"
	"os"
)

var (
	_logger *log.Logger
)

func init() {
	_logger = log.New(os.Stderr, "[TCP Server] ", log.LstdFlags)
}

func SetLogger(logger *log.Logger) {
	_logger = logger
}

func GetLogger() *log.Logger {
	return _logger
}
