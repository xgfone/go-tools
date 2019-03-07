// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package option is used to represent an optional value referring to Option in Rust.
package option

import (
	"errors"
	"fmt"

	"github.com/xgfone/go-tools/types"
)

// NONE is the global None value.
var NONE = None()

// option represents an optional value.
type option struct {
	value interface{}
}

// Some returns an Option.
//
// If v is nil, it will be a None value.
func Some(v interface{}) Option {
	return &option{value: v}
}

// None is equal to Some(nil).
func None() Option {
	return Some(nil)
}

// IsSome reports whether there is a value.
func (o *option) IsSome() bool {
	return o.value != nil
}

// IsNone reports whether the value is None.
func (o *option) IsNone() bool {
	return o.value == nil
}

// Value returns the inner value. Return nil if it's a None.
func (o *option) Value() interface{} {
	return o.value
}

// Reset resets the inner value.
func (o *option) Reset(value interface{}) {
	o.value = value
}

// Scan parses src and assigns to iteself.
//
// Notice: it does not convert src, only assign it to the inner value.
func (o *option) Scan(src interface{}) error {
	o.value = src
	return nil
}

// ConvertTo converts the value by convert then assigns the result to the inner.
func (o *option) ConvertTo(value interface{}, convert func(interface{}) (interface{}, error)) error {
	v, err := convert(value)
	if err == nil {
		o.value = v
	}
	return err
}

// Some returns the inner value, but panic if it's a None.
func (o *option) Some() interface{} {
	if o.value == nil {
		panic(errors.New("the value is None"))
	}
	return o.value
}

// None check whether the inner value is None and panic if it's not a None.
func (o *option) None() {
	if o.value != nil {
		panic(errors.New("the value is not None"))
	}
}

// SomeOr returns the inner value if it's not None. Or return v.
func (o *option) SomeOr(v interface{}) interface{} {
	if o.value == nil {
		return v
	}
	return o.value
}

// String implements the interface fmt.Stringer.
func (o *option) String() string {
	return fmt.Sprintf("Option(%v)", o.value)
}

// Str returns the inner string value. Or panic.
func (o *option) Str() string {
	return o.value.(string)
}

// Bytes returns the inner []byte value. Or panic.
func (o *option) Bytes() []byte {
	return o.value.([]byte)
}

// Bool returns the inner bool value. Or panic.
func (o *option) Bool() bool {
	return o.value.(bool)
}

// Byte returns the inner byte value. Or panic.
func (o *option) Byte() byte {
	return o.value.(byte)
}

// Rune returns the inner rune value. Or panic.
func (o *option) Rune() rune {
	return o.value.(rune)
}

// Int returns the inner int value. Or panic.
func (o *option) Int() int {
	return o.value.(int)
}

// Int8 returns the inner int8 value. Or panic.
func (o *option) Int8() int8 {
	return o.value.(int8)
}

// Int16 returns the inner int16 value. Or panic.
func (o *option) Int16() int16 {
	return o.value.(int16)
}

// Int32 returns the inner int32 value. Or panic.
func (o *option) Int32() int32 {
	return o.value.(int32)
}

// Int64 returns the inner int64 value. Or panic.
func (o *option) Int64() int64 {
	return o.value.(int64)
}

// Uint returns the inner uint value. Or panic.
func (o *option) Uint() uint {
	return o.value.(uint)
}

// Uint8 returns the inner uint8 value. Or panic.
func (o *option) Uint8() uint8 {
	return o.value.(uint8)
}

// Uint16 returns the inner uint16 value. Or panic.
func (o *option) Uint16() uint16 {
	return o.value.(uint16)
}

// Uint32 returns the inner uint32 value. Or panic.
func (o *option) Uint32() uint32 {
	return o.value.(uint32)
}

// Uint64 returns the inner uin64t value. Or panic.
func (o *option) Uint64() uint64 {
	return o.value.(uint64)
}

// Float32 returns the inner float32 value. Or panic.
func (o *option) Float32() float32 {
	return o.value.(float32)
}

// Float64 returns the inner float64 value. Or panic.
func (o *option) Float64() float64 {
	return o.value.(float64)
}

// Strs returns the inner []string value. Or panic.
func (o *option) Strs() []string {
	return o.value.([]string)
}

// Interfaces returns the inner []interface{} value. Or panic.
func (o *option) Interfaces() []interface{} {
	return o.value.([]interface{})
}

// Map returns the inner map[string]interface{} value. Or panic.
func (o *option) Map() map[string]interface{} {
	return o.value.(map[string]interface{})
}

// StrMap returns the inner map[string]string value. Or panic.
func (o *option) StrMap() map[string]string {
	return o.value.(map[string]string)
}

// IsString reports whether the type of the value is string.
func (o *option) IsString() bool {
	switch o.value.(type) {
	case string:
		return true
	}
	return false
}

// IsBytes reports whether the type of the value is []byte.
func (o *option) IsBytes() bool {
	switch o.value.(type) {
	case []byte:
		return true
	}
	return false
}

// IsBool reports whether the type of the value is bool.
func (o *option) IsBool() bool {
	switch o.value.(type) {
	case bool:
		return true
	}
	return false
}

// IsInt reports whether the type of the value is int.
func (o *option) IsInt() bool {
	switch o.value.(type) {
	case int:
		return true
	}
	return false
}

// IsInt8 reports whether the type of the value is int8.
func (o *option) IsInt8() bool {
	switch o.value.(type) {
	case int8:
		return true
	}
	return false
}

// IsInt16 reports whether the type of the value is int16.
func (o *option) IsInt16() bool {
	switch o.value.(type) {
	case int16:
		return true
	}
	return false
}

// IsInt32 reports whether the type of the value is int32.
func (o *option) IsInt32() bool {
	switch o.value.(type) {
	case int32:
		return true
	}
	return false
}

// IsInt64 reports whether the type of the value is int64.
func (o *option) IsInt64() bool {
	switch o.value.(type) {
	case int64:
		return true
	}
	return false
}

// IsUint reports whether the type of the value is uint.
func (o *option) IsUint() bool {
	switch o.value.(type) {
	case uint:
		return true
	}
	return false
}

// IsUint8 reports whether the type of the value is uint8.
func (o *option) IsUint8() bool {
	switch o.value.(type) {
	case uint8:
		return true
	}
	return false
}

// IsUint16 reports whether the type of the value is uint16.
func (o *option) IsUint16() bool {
	switch o.value.(type) {
	case uint16:
		return true
	}
	return false
}

// IsUint32 reports whether the type of the value is uint32.
func (o *option) IsUint32() bool {
	switch o.value.(type) {
	case uint32:
		return true
	}
	return false
}

// IsUint64 reports whether the type of the value is uint64
func (o *option) IsUint64() bool {
	switch o.value.(type) {
	case uint64:
		return true
	}
	return false
}

// IsFloat32 reports whether the type of the value is float32.
func (o *option) IsFloat32() bool {
	switch o.value.(type) {
	case float32:
		return true
	}
	return false
}

// IsFloat64 reports whether the type of the value is float64.
func (o *option) IsFloat64() bool {
	switch o.value.(type) {
	case float64:
		return true
	}
	return false
}

// IsSignedInteger reports whether the value is a signed integer.
func (o *option) IsSignedInteger() bool {
	switch o.value.(type) {
	case int, int8, int16, int32, int64:
		return true
	}
	return false
}

// IsUnsignedInteger reports whether the value is an unsigned integer.
func (o *option) IsUnsignedInteger() bool {
	switch o.value.(type) {
	case uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

// IsInteger reports whether the value is a signed or unsigned integer.
func (o *option) IsInteger() bool {
	switch o.value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

// IsFloat reports whether the value is a float32 or float64.
func (o *option) IsFloat() bool {
	switch o.value.(type) {
	case float32, float64:
		return true
	}
	return false
}

// IsNumber reports whether the value is an integer or float.
func (o *option) IsNumber() bool {
	switch o.value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return true
	}
	return false
}

// IsStrs reports whether the type of the value is []string.
func (o *option) IsStrs() bool {
	_, ok := o.value.([]string)
	return ok
}

// IsInterfaces reports whether the type of the value is []interface.
func (o *option) IsInterfaces() bool {
	_, ok := o.value.([]interface{})
	return ok
}

// IsMap reports whether the type of the value is map[string]interface{}.
func (o *option) IsMap() bool {
	_, ok := o.value.(map[string]interface{})
	return ok
}

// IsStrMap reports whether the type of the value is map[string]string.
func (o *option) IsStrMap() bool {
	_, ok := o.value.(map[string]string)
	return ok
}

// ToString converts the inner value to string.
func (o *option) ToString() (string, error) {
	return types.ToString(o.value)
}

// ToBool converts the inner value to bool.
func (o *option) ToBool() (bool, error) {
	return types.ToBool(o.value)
}

// ToInt converts the inner value to int.
func (o *option) ToInt() (int, error) {
	v, err := types.ToInt64(o.value)
	return int(v), err
}

// ToInt8 converts the inner value to int8.
func (o *option) ToInt8() (int8, error) {
	v, err := types.ToInt64(o.value)
	return int8(v), err
}

// ToInt16 converts the inner value to int16.
func (o *option) ToInt16() (int16, error) {
	v, err := types.ToInt64(o.value)
	return int16(v), err
}

// ToInt32 converts the inner value to int32.
func (o *option) ToInt32() (int32, error) {
	v, err := types.ToInt64(o.value)
	return int32(v), err
}

// ToInt64 converts the inner value to int64.
func (o *option) ToInt64() (int64, error) {
	return types.ToInt64(o.value)
}

// ToUint converts the inner value to uint.
func (o *option) ToUint() (uint, error) {
	v, err := types.ToUint64(o.value)
	return uint(v), err
}

// ToUint8 converts the inner value to uint8.
func (o *option) ToUint8() (uint8, error) {
	v, err := types.ToUint64(o.value)
	return uint8(v), err
}

// ToUint16 converts the inner value to uint16.
func (o *option) ToUint16() (uint16, error) {
	v, err := types.ToUint64(o.value)
	return uint16(v), err
}

// ToUint32 converts the inner value to uint32.
func (o *option) ToUint32() (uint32, error) {
	v, err := types.ToUint64(o.value)
	return uint32(v), err
}

// ToUint64 converts the inner value to uint64.
func (o *option) ToUint64() (uint64, error) {
	return types.ToUint64(o.value)
}

// ToFloat32 converts the inner value to float32.
func (o *option) ToFloat32() (float32, error) {
	v, err := types.ToFloat64(o.value)
	return float32(v), err
}

// ToFloat64 converts the inner value to float64.
func (o *option) ToFloat64() (float64, error) {
	return types.ToFloat64(o.value)
}
