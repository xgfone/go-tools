// Package errors supplies an error type implementation based on
// the type inheritance.
package errors

import (
	"fmt"
	"sync/atomic"
)

// ErrorT represents a error type.
type ErrorT interface {
	String() string

	SubType(name string) ErrorT
	IsParentOf(child ErrorT) bool
	IsChildOf(parent ErrorT) bool

	New(format string, args ...interface{}) Error
}

// Error represents a error.
type Error interface {
	Error() string
	IsType(err ErrorT) bool
}

// IsError reports whether err is an Error.
func IsError(err error) bool {
	_, ok := err.(Error)
	return ok
}

// IsType reports whether err is an Error and whose type is t.
func IsType(err error, t ErrorT) bool {
	if e, ok := err.(Error); ok {
		return e.IsType(t)
	}
	return false
}

type terror struct {
	t ErrorT
	e string
}

func (te terror) Error() string {
	return te.e
}

func (te terror) IsType(e ErrorT) bool {
	return te.t.(terrorT).index == e.(terrorT).index || te.t.IsChildOf(e)
}

var errorTindex uint64

// BaseErrorT The topmost ErrorT, which is the parent type of all the ErrorT.
var BaseErrorT ErrorT = terrorT{}

// NewType returns a new ErrorT type with the name.
func NewType(name string) ErrorT {
	return BaseErrorT.SubType(name)
}

// New returns a new Error.
func New(format string, args ...interface{}) Error {
	return BaseErrorT.New(format, args...)
}

type terrorT struct {
	types []uint64
	index uint64
	name  string
}

func (et terrorT) String() string {
	return fmt.Sprintf("ErrorT(%s)", et.name)
}

func (et terrorT) SubType(name string) ErrorT {
	types := make([]uint64, len(et.types)+1)
	copy(types, et.types)
	types[len(et.types)] = et.index
	return terrorT{name: name, types: types, index: atomic.AddUint64(&errorTindex, 1)}
}

func (et terrorT) IsParentOf(child ErrorT) bool {
	for _, index := range child.(terrorT).types {
		if index == et.index {
			return true
		}
	}
	return false
}

func (et terrorT) IsChildOf(parent ErrorT) bool {
	pindex := parent.(terrorT).index
	for _, index := range et.types {
		if index == pindex {
			return true
		}
	}
	return false
}

func (et terrorT) New(format string, args ...interface{}) Error {
	if len(args) == 0 {
		return terror{t: et, e: format}
	}
	return terror{t: et, e: fmt.Sprintf(format, args...)}
}
