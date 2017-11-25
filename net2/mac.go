package net2

import (
	"fmt"
	"strconv"
	"strings"
)

// NormalizeMac normalizes the mac.
//
// If fill is true, pad with leading zero, such as 01:02:03:04:05:06.
//
// If upper is true, output the upper character, such as AA:BB:11:22:33:44.
// Or, output the lower character, such as aa:bb:cc:11:22:33.
//
// Return "" if the mac is a invalid mac.
func NormalizeMac(mac string, fill, upper bool) string {
	macs := strings.Split(mac, ":")
	if len(macs) != 6 {
		return ""
	}

	width := ""
	_upper := "x"
	if upper {
		_upper = "X"
	}
	if fill {
		width = "2"
	}
	formatter := fmt.Sprintf("%%0%s%s", width, _upper)

	for i, m := range macs {
		v, err := strconv.ParseUint(m, 16, 8)
		if err != nil {
			return ""
		}
		macs[i] = fmt.Sprintf(formatter, v)
	}

	return strings.Join(macs, ":")
}

// NormalizeMacFU is equal to NormalizeMac(mac, true, true).
func NormalizeMacFU(mac string) string {
	return NormalizeMac(mac, true, true)
}

// NormalizeMacFu is equal to NormalizeMac(mac, true, false).
func NormalizeMacFu(mac string) string {
	return NormalizeMac(mac, true, false)
}

// NormalizeMacfU is equal to NormalizeMac(mac, false, true).
func NormalizeMacfU(mac string) string {
	return NormalizeMac(mac, false, true)
}

// NormalizeMacfu is equal to NormalizeMac(mac, false, false).
func NormalizeMacfu(mac string) string {
	return NormalizeMac(mac, false, false)
}
