// Package parse converts something from a string to `bool`, `int`, `uint`,
// `float`, or from a certain type to string, etc.
//
// If failed, return their ZERO value.
//
package parse

import "strconv"

// ToB is short for ToBool.
func ToB(v string) bool {
	if x, err := strconv.ParseBool(v); err != nil {
		return false
	} else {
		return x
	}
}

// ToBool converts v to bool.
func ToBool(v string) bool {
	return ToB(v)
}

func toFloat(v string, bit int) float64 {
	if x, err := strconv.ParseFloat(v, bit); err != nil {
		return 0.0
	} else {
		return float64(x)
	}
}

// ToF32 converts v to float32.
func ToF32(v string) float32 {
	return float32(toFloat(v, 32))
}

// ToF64 converts v to float64.
func ToF64(v string) float64 {
	return toFloat(v, 64)
}

func toInt(v string, base int, bit int) int64 {
	if x, err := strconv.ParseInt(v, base, bit); err != nil {
		return 0
	} else {
		return int64(x)
	}
}

// ToI is short for ToInt.
func ToI(v string, base int) int {
	return int(toInt(v, base, 0))
}

// ToInt converts v to int in base.
func ToInt(v string, base int) int {
	return ToI(v, base)
}

// ToI8 converts v to int8 in base.
func ToI8(v string, base int) int8 {
	return int8(toInt(v, base, 8))
}

// ToI16 converts v to int16 in base.
func ToI16(v string, base int) int16 {
	return int16(toInt(v, base, 16))
}

// ToI32 converts v to int32 in base.
func ToI32(v string, base int) int32 {
	return int32(toInt(v, base, 32))
}

// ToI64 converts v to int64 in base.
func ToI64(v string, base int) int64 {
	return int64(toInt(v, base, 64))
}

func toUint(v string, base int, bit int) uint64 {
	if x, err := strconv.ParseInt(v, base, bit); err != nil {
		return 0
	} else {
		return uint64(x)
	}
}

// ToU is short for ToUint.
func ToU(v string, base int) uint {
	return uint(toUint(v, base, 0))
}

// ToUint converts v to uint in base.
func ToUint(v string, base int) uint {
	return ToU(v, base)
}

// ToU8 converts v to uint8 in base.
func ToU8(v string, base int) uint8 {
	return uint8(toUint(v, base, 8))
}

// ToU16 converts v to uint16 in base.
func ToU16(v string, base int) uint16 {
	return uint16(toUint(v, base, 16))
}

// ToU32 converts v to uint32 in base.
func ToU32(v string, base int) uint32 {
	return uint32(toUint(v, base, 32))
}

// ToU64 converts v to uint64 in base.
func ToU64(v string, base int) uint64 {
	return uint64(toUint(v, base, 64))
}
