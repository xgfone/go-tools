// +build go1.9

package server

import (
	"github.com/xgfone/go-tools/net2"
)

type (
	// THandle is the type alias of net2.THandle.
	//
	// DEPRECATED!!! Please the package net2.
	THandle = net2.THandle

	// THandleFunc is the type alias of net2.THandleFunc.
	//
	// DEPRECATED!!! Please the package net2.
	THandleFunc = net2.THandleFunc
)

var (
	// TCPWrapError is the alias of net2.TCPWrapError.
	//
	// DEPRECATED!!! Please the package net2.
	TCPWrapError = net2.TCPWrapError

	// TCPServerForever is the alias of net2.TCPServerForever.
	//
	// DEPRECATED!!! Please the package net2.
	TCPServerForever = net2.TCPServerForever

	// DialTCP is the alias of net2.DialTCP.
	//
	// DEPRECATED!!! Please the package net2.
	DialTCP = net2.DialTCP

	// DialTCPWithAddr is the alias of net2.DialTCPWithAddr.
	//
	// DEPRECATED!!! Please the package net2.
	DialTCPWithAddr = net2.DialTCPWithAddr
)
