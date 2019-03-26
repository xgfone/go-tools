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

package option

import (
	"encoding/json"

	"github.com/xgfone/go-tools/types"
)

// BoolOption is an Option of the bool type.
type BoolOption struct {
	Option
}

// NewBoolOption returns a new BoolOption.
func NewBoolOption(o Option) BoolOption {
	if o == nil {
		o = None()
	}
	return BoolOption{Option: o}
}

// Scan converts src as bool to the inner value.
func (o BoolOption) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToBool(v) })
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o BoolOption) UnmarshalJSON(src []byte) (err error) {
	var v bool
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// StringOption is an Option of the string type.
type StringOption struct {
	Option
}

// NewStringOption returns a new StringOption.
func NewStringOption(o Option) StringOption {
	if o == nil {
		o = None()
	}
	return StringOption{Option: o}
}

// Scan converts src as string to the inner value.
func (o StringOption) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToString(v) })
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o StringOption) UnmarshalJSON(src []byte) (err error) {
	o.Reset(string(src))
	return
}

// IntOption is an Option of the int type.
type IntOption struct {
	Option
}

// NewIntOption returns a new IntOption.
func NewIntOption(o Option) IntOption {
	if o == nil {
		o = None()
	}
	return IntOption{Option: o}
}

// Scan converts src as int to the inner value.
func (o IntOption) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToInt(v) })
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o IntOption) UnmarshalJSON(src []byte) (err error) {
	var v int
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// Int8Option is an Option of the int8 type.
type Int8Option struct {
	Option
}

// NewInt8Option returns a new Int8Option.
func NewInt8Option(o Option) Int8Option {
	if o == nil {
		o = None()
	}
	return Int8Option{Option: o}
}

// Scan converts src as int8 to the inner value.
func (o Int8Option) Scan(src interface{}) error {
	v, err := types.ToInt64(src)
	if err == nil {
		o.Reset(int8(v))
	}
	return err
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o Int8Option) UnmarshalJSON(src []byte) (err error) {
	var v int8
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// Int16Option is an Option of the int16 type.
type Int16Option struct {
	Option
}

// NewInt16Option returns a new Int16Option.
func NewInt16Option(o Option) Int16Option {
	if o == nil {
		o = None()
	}
	return Int16Option{Option: o}
}

// Scan converts src as int16 to the inner value.
func (o Int16Option) Scan(src interface{}) error {
	v, err := types.ToInt64(src)
	if err == nil {
		o.Reset(int16(v))
	}
	return err
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o Int16Option) UnmarshalJSON(src []byte) (err error) {
	var v int16
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// Int32Option is an Option of the int32 type.
type Int32Option struct {
	Option
}

// NewInt32Option returns a new Int32Option.
func NewInt32Option(o Option) Int32Option {
	if o == nil {
		o = None()
	}
	return Int32Option{Option: o}
}

// Scan converts src as int32 to the inner value.
func (o Int32Option) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToInt32(v) })
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o Int32Option) UnmarshalJSON(src []byte) (err error) {
	var v int32
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// Int64Option is an Option of the int64 type.
type Int64Option struct {
	Option
}

// NewInt64Option returns a new Int64Option.
func NewInt64Option(o Option) Int64Option {
	if o == nil {
		o = None()
	}
	return Int64Option{Option: o}
}

// Scan converts src as int64 to the inner value.
func (o Int64Option) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToInt64(v) })
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o Int64Option) UnmarshalJSON(src []byte) (err error) {
	var v int64
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// UintOption is an Option of the uint type.
type UintOption struct {
	Option
}

// NewUintOption returns a new UintOption.
func NewUintOption(o Option) UintOption {
	if o == nil {
		o = None()
	}
	return UintOption{Option: o}
}

// Scan converts src as uint to the inner value.
func (o UintOption) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToUint(v) })
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o UintOption) UnmarshalJSON(src []byte) (err error) {
	var v uint
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// Uint8Option is an Option of the uint8 type.
type Uint8Option struct {
	Option
}

// NewUint8Option returns a new Uint8Option.
func NewUint8Option(o Option) Uint8Option {
	if o == nil {
		o = None()
	}
	return Uint8Option{Option: o}
}

// Scan converts src as uint8 to the inner value.
func (o Uint8Option) Scan(src interface{}) error {
	v, err := types.ToUint64(src)
	if err == nil {
		o.Reset(uint8(v))
	}
	return err
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o Uint8Option) UnmarshalJSON(src []byte) (err error) {
	var v uint8
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// Uint16Option is an Option of the uint16 type.
type Uint16Option struct {
	Option
}

// NewUint16Option returns a new Uint16Option.
func NewUint16Option(o Option) Uint16Option {
	if o == nil {
		o = None()
	}
	return Uint16Option{Option: o}
}

// Scan converts src as uint16 to the inner value.
func (o Uint16Option) Scan(src interface{}) error {
	v, err := types.ToUint64(src)
	if err == nil {
		o.Reset(uint16(v))
	}
	return err
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o Uint16Option) UnmarshalJSON(src []byte) (err error) {
	var v uint16
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// Uint32Option is an Option of the uint32 type.
type Uint32Option struct {
	Option
}

// NewUint32Option returns a new Uint32Option.
func NewUint32Option(o Option) Uint32Option {
	if o == nil {
		o = None()
	}
	return Uint32Option{Option: o}
}

// Scan converts src as uint32 to the inner value.
func (o Uint32Option) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToUint32(v) })
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o Uint32Option) UnmarshalJSON(src []byte) (err error) {
	var v uint32
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// Uint64Option is an Option of the uint64 type.
type Uint64Option struct {
	Option
}

// NewUint64Option returns a new Uint64Option.
func NewUint64Option(o Option) Uint64Option {
	if o == nil {
		o = None()
	}
	return Uint64Option{Option: o}
}

// Scan converts src as uint64 to the inner value.
func (o Uint64Option) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToUint64(v) })
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o Uint64Option) UnmarshalJSON(src []byte) (err error) {
	var v uint64
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// Float64Option is an Option of the float64 type.
type Float64Option struct {
	Option
}

// NewFloat64Option returns a new Float64Option.
func NewFloat64Option(o Option) Float64Option {
	if o == nil {
		o = None()
	}
	return Float64Option{Option: o}
}

// Scan converts src as float64 to the inner value.
func (o Float64Option) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToFloat64(v) })
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o Float64Option) UnmarshalJSON(src []byte) (err error) {
	var v float64
	if err = json.Unmarshal(src, &v); err == nil {
		o.Reset(v)
	}
	return
}

// TimeOption is an Option of the time.Time type.
type TimeOption struct {
	Option
}

// NewTimeOption returns a new TimeOption.
func NewTimeOption(o Option) TimeOption {
	if o == nil {
		o = None()
	}
	return TimeOption{Option: o}
}

// Scan converts src as float64 to the inner value.
func (o TimeOption) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) {
		return types.ToTime(v, types.DateTimeLayout)
	})
}

// UnmarshalJSON implements the interface json.Unmarshaler.
func (o TimeOption) UnmarshalJSON(src []byte) (err error) {
	v, err := types.ToTime(src, types.DateTimeLayout)
	if err == nil {
		o.Reset(v)
	}
	return
}
