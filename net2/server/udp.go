// +build go1.9

package server

import (
	"github.com/xgfone/go-tools/net2"
)

type (
	// UHandle is the type alias of net2.UHandle.
	//
	// DEPRECATED!!! Please the package net2.
	UHandle = net2.UHandle

	// UHandleFunc is the type alias of net2.UHandleFunc.
	//
	// DEPRECATED!!! Please the package net2.
	UHandleFunc = net2.UHandleFunc
)

var (
	// UDPWithError is the alias of net2.UDPWithError.
	//
	// DEPRECATED!!! Please the package net2.
	UDPWithError = net2.UDPWithError

	// UDPServerForever is the alias of net2.UDPServerForever.
	//
	// DEPRECATED!!! Please the package net2.
	UDPServerForever = net2.UDPServerForever

	// DialUDP is the alias of net2.DialUDP.
	//
	// DEPRECATED!!! Please the package net2.
	DialUDP = net2.DialUDP

	// DialUDPWithAddr is the alias of net2.DialUDPWithAddr.
	//
	// DEPRECATED!!! Please the package net2.
	DialUDPWithAddr = net2.DialUDPWithAddr
)
