// Package option is used to represent an optional value referring to Option in Rust.
package option

import (
	"errors"
	"fmt"

	"github.com/xgfone/go-tools/types"
)

// NONE is the global None value.
var NONE = None()

// Option represents an optional value.
type Option struct {
	value interface{}
}

// Some returns an Option.
//
// If v is nil, it will be a None value.
func Some(v interface{}) Option {
	return Option{value: v}
}

// None is equal to Some(nil).
func None() Option {
	return Some(nil)
}

// IsSome reports whether there is a value.
func (o Option) IsSome() bool {
	return o.value != nil
}

// IsNone reports whether the value is None.
func (o Option) IsNone() bool {
	return o.value == nil
}

// Value returns the inner value. Return nil if it's a None.
func (o Option) Value() interface{} {
	return o.value
}

// Reset resets the inner value.
func (o *Option) Reset(value interface{}) {
	o.value = value
}

// ConvertTo converts the value to the inner by convert, which `convert` will
// convert `value` then assign the result to the inner.
func (o *Option) ConvertTo(value interface{}, convert func(interface{}) (interface{}, error)) error {
	v, err := convert(value)
	if err == nil {
		o.value = v
	}
	return err
}

// Some returns the inner value, but panic if it's a None.
func (o Option) Some() interface{} {
	if o.value == nil {
		panic(errors.New("the value is None"))
	}
	return o.value
}

// None check whether the inner value is None and panic if it's not a None.
func (o Option) None() {
	if o.value != nil {
		panic(errors.New("the value is not None"))
	}
}

// SomeOr returns the inner value if it's not None. Or return v.
func (o Option) SomeOr(v interface{}) interface{} {
	if o.value == nil {
		return v
	}
	return o.value
}

// String implements the interface fmt.Stringer.
func (o Option) String() string {
	return fmt.Sprintf("Option(%v)", o.value)
}

// Str returns the inner string value. Or panic.
func (o Option) Str() string {
	return o.value.(string)
}

// Bytes returns the inner []byte value. Or panic.
func (o Option) Bytes() []byte {
	return o.value.([]byte)
}

// Bool returns the inner bool value. Or panic.
func (o Option) Bool() bool {
	return o.value.(bool)
}

// Byte returns the inner byte value. Or panic.
func (o Option) Byte() byte {
	return o.value.(byte)
}

// Rune returns the inner rune value. Or panic.
func (o Option) Rune() rune {
	return o.value.(rune)
}

// Int returns the inner int value. Or panic.
func (o Option) Int() int {
	return o.value.(int)
}

// Int8 returns the inner int8 value. Or panic.
func (o Option) Int8() int8 {
	return o.value.(int8)
}

// Int16 returns the inner int16 value. Or panic.
func (o Option) Int16() int16 {
	return o.value.(int16)
}

// Int32 returns the inner int32 value. Or panic.
func (o Option) Int32() int32 {
	return o.value.(int32)
}

// Int64 returns the inner int64 value. Or panic.
func (o Option) Int64() int64 {
	return o.value.(int64)
}

// Uint returns the inner uint value. Or panic.
func (o Option) Uint() uint {
	return o.value.(uint)
}

// Uint8 returns the inner uint8 value. Or panic.
func (o Option) Uint8() uint8 {
	return o.value.(uint8)
}

// Uint16 returns the inner uint16 value. Or panic.
func (o Option) Uint16() uint16 {
	return o.value.(uint16)
}

// Uint32 returns the inner uint32 value. Or panic.
func (o Option) Uint32() uint32 {
	return o.value.(uint32)
}

// Uint64 returns the inner uin64t value. Or panic.
func (o Option) Uint64() uint64 {
	return o.value.(uint64)
}

// Float32 returns the inner float32 value. Or panic.
func (o Option) Float32() float32 {
	return o.value.(float32)
}

// Float64 returns the inner float64 value. Or panic.
func (o Option) Float64() float64 {
	return o.value.(float64)
}

// Strs returns the inner []string value. Or panic.
func (o Option) Strs() []string {
	return o.value.([]string)
}

// Interfaces returns the inner []interface{} value. Or panic.
func (o Option) Interfaces() []interface{} {
	return o.value.([]interface{})
}

// Map returns the inner map[string]interface{} value. Or panic.
func (o Option) Map() map[string]interface{} {
	return o.value.(map[string]interface{})
}

// StrMap returns the inner map[string]string value. Or panic.
func (o Option) StrMap() map[string]string {
	return o.value.(map[string]string)
}

// IsString reports whether the type of the value is string.
func (o Option) IsString() bool {
	switch o.value.(type) {
	case string:
		return true
	}
	return false
}

// IsBytes reports whether the type of the value is []byte.
func (o Option) IsBytes() bool {
	switch o.value.(type) {
	case []byte:
		return true
	}
	return false
}

// IsBool reports whether the type of the value is bool.
func (o Option) IsBool() bool {
	switch o.value.(type) {
	case bool:
		return true
	}
	return false
}

// IsInt reports whether the type of the value is int.
func (o Option) IsInt() bool {
	switch o.value.(type) {
	case int:
		return true
	}
	return false
}

// IsInt8 reports whether the type of the value is int8.
func (o Option) IsInt8() bool {
	switch o.value.(type) {
	case int8:
		return true
	}
	return false
}

// IsInt16 reports whether the type of the value is int16.
func (o Option) IsInt16() bool {
	switch o.value.(type) {
	case int16:
		return true
	}
	return false
}

// IsInt32 reports whether the type of the value is int32.
func (o Option) IsInt32() bool {
	switch o.value.(type) {
	case int32:
		return true
	}
	return false
}

// IsInt64 reports whether the type of the value is int64.
func (o Option) IsInt64() bool {
	switch o.value.(type) {
	case int64:
		return true
	}
	return false
}

// IsUint reports whether the type of the value is uint.
func (o Option) IsUint() bool {
	switch o.value.(type) {
	case uint:
		return true
	}
	return false
}

// IsUint8 reports whether the type of the value is uint8.
func (o Option) IsUint8() bool {
	switch o.value.(type) {
	case uint8:
		return true
	}
	return false
}

// IsUint16 reports whether the type of the value is uint16.
func (o Option) IsUint16() bool {
	switch o.value.(type) {
	case uint16:
		return true
	}
	return false
}

// IsUint32 reports whether the type of the value is uint32.
func (o Option) IsUint32() bool {
	switch o.value.(type) {
	case uint32:
		return true
	}
	return false
}

// IsUint64 reports whether the type of the value is uint64
func (o Option) IsUint64() bool {
	switch o.value.(type) {
	case uint64:
		return true
	}
	return false
}

// IsFloat32 reports whether the type of the value is float32.
func (o Option) IsFloat32() bool {
	switch o.value.(type) {
	case float32:
		return true
	}
	return false
}

// IsFloat64 reports whether the type of the value is float64.
func (o Option) IsFloat64() bool {
	switch o.value.(type) {
	case float64:
		return true
	}
	return false
}

// IsSignedInteger reports whether the value is a signed integer.
func (o Option) IsSignedInteger() bool {
	switch o.value.(type) {
	case int, int8, int16, int32, int64:
		return true
	}
	return false
}

// IsUnsignedInteger reports whether the value is an unsigned integer.
func (o Option) IsUnsignedInteger() bool {
	switch o.value.(type) {
	case uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

// IsInteger reports whether the value is a signed or unsigned integer.
func (o Option) IsInteger() bool {
	switch o.value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

// IsFloat reports whether the value is a float32 or float64.
func (o Option) IsFloat() bool {
	switch o.value.(type) {
	case float32, float64:
		return true
	}
	return false
}

// IsNumber reports whether the value is an integer or float.
func (o Option) IsNumber() bool {
	switch o.value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return true
	}
	return false
}

// IsStrs reports whether the type of the value is []string.
func (o Option) IsStrs() bool {
	_, ok := o.value.([]string)
	return ok
}

// IsInterfaces reports whether the type of the value is []interface.
func (o Option) IsInterfaces() bool {
	_, ok := o.value.([]interface{})
	return ok
}

// IsMap reports whether the type of the value is map[string]interface{}.
func (o Option) IsMap() bool {
	_, ok := o.value.(map[string]interface{})
	return ok
}

// IsStrMap reports whether the type of the value is map[string]string.
func (o Option) IsStrMap() bool {
	_, ok := o.value.(map[string]string)
	return ok
}

// ToString converts the inner value to string.
func (o Option) ToString() (string, error) {
	return types.ToString(o.value)
}

// ToBool converts the inner value to bool.
func (o Option) ToBool() (bool, error) {
	return types.ToBool(o.value)
}

// ToInt converts the inner value to int.
func (o Option) ToInt() (int, error) {
	v, err := types.ToInt64(o.value)
	return int(v), err
}

// ToInt8 converts the inner value to int8.
func (o Option) ToInt8() (int8, error) {
	v, err := types.ToInt64(o.value)
	return int8(v), err
}

// ToInt16 converts the inner value to int16.
func (o Option) ToInt16() (int16, error) {
	v, err := types.ToInt64(o.value)
	return int16(v), err
}

// ToInt32 converts the inner value to int32.
func (o Option) ToInt32() (int32, error) {
	v, err := types.ToInt64(o.value)
	return int32(v), err
}

// ToInt64 converts the inner value to int64.
func (o Option) ToInt64() (int64, error) {
	return types.ToInt64(o.value)
}

// ToUint converts the inner value to uint.
func (o Option) ToUint() (uint, error) {
	v, err := types.ToUint64(o.value)
	return uint(v), err
}

// ToUint8 converts the inner value to uint8.
func (o Option) ToUint8() (uint8, error) {
	v, err := types.ToUint64(o.value)
	return uint8(v), err
}

// ToUint16 converts the inner value to uint16.
func (o Option) ToUint16() (uint16, error) {
	v, err := types.ToUint64(o.value)
	return uint16(v), err
}

// ToUint32 converts the inner value to uint32.
func (o Option) ToUint32() (uint32, error) {
	v, err := types.ToUint64(o.value)
	return uint32(v), err
}

// ToUint64 converts the inner value to uint64.
func (o Option) ToUint64() (uint64, error) {
	return types.ToUint64(o.value)
}

// ToFloat32 converts the inner value to float32.
func (o Option) ToFloat32() (float32, error) {
	v, err := types.ToFloat64(o.value)
	return float32(v), err
}

// ToFloat64 converts the inner value to float64.
func (o Option) ToFloat64() (float64, error) {
	return types.ToFloat64(o.value)
}
