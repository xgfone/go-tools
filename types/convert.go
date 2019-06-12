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

package types

import (
	"fmt"
	"time"
)

// DateTimeLayout is the DateTime layout to parse the value to the time.Time.
const DateTimeLayout = "2006-01-02 15:04:05"

var (
	// ErrNotSliceOrArray is returned when the value is not a slice.
	ErrNotSliceOrArray = fmt.Errorf("the value is not a slice or array")

	// ErrNotMap is returned when the value is not a map.
	ErrNotMap = fmt.Errorf("the value is not a map")

	// ErrNotString is returned when the type of the key is not string.
	ErrNotString = fmt.Errorf("the type of the key is not string")

	// ErrKindNotExist is returned when a certain kind does not exist.
	ErrKindNotExist = fmt.Errorf("no kind")

	// ErrUnknownType is returned when not to identify a data type.
	ErrUnknownType = fmt.Errorf("unknown type")
)

var converters = make(map[Kind]func(interface{}) (interface{}, error))

// Predefine some kinds.
const (
	Unknown Kind = iota
	Nil
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
	Bytes
	Time        // For the format "YYYY-MM-DD HH:MM:SS"
	RFC3339Time // For the format time.RFC3339
)

// Kind represents the kind of the converter.
type Kind int

func init() {
	RegisterConverter(Bool, func(v interface{}) (interface{}, error) { return ToBool(v) })
	RegisterConverter(Int, func(v interface{}) (interface{}, error) { return ToInt(v) })
	RegisterConverter(Int32, func(v interface{}) (interface{}, error) { return ToInt32(v) })
	RegisterConverter(Int64, func(v interface{}) (interface{}, error) { return ToInt64(v) })
	RegisterConverter(Uint, func(v interface{}) (interface{}, error) { return ToUint(v) })
	RegisterConverter(Uint32, func(v interface{}) (interface{}, error) { return ToUint32(v) })
	RegisterConverter(Uint64, func(v interface{}) (interface{}, error) { return ToUint64(v) })
	RegisterConverter(Float64, func(v interface{}) (interface{}, error) { return ToFloat64(v) })
	RegisterConverter(String, func(v interface{}) (interface{}, error) { return ToString(v) })
	RegisterConverter(RFC3339Time, func(v interface{}) (interface{}, error) { return ToTime(v, time.RFC3339) })
	RegisterConverter(Time, func(v interface{}) (interface{}, error) {
		if s, _ := v.(string); s == "0000-00-00 00:00:00" {
			return time.Time{}, nil
		}
		return ToTime(v, "2006-01-02 15:04:05")
	})
}

// RegisterConverter registers a converter of the kind k.
//
// By default it has registered the kinds as follow:
//
//     Bool
//     String
//     Float64
//     Int, Int32, Int64
//     Uint, Uint32, Uint64
//     Time, RFC3339Time
func RegisterConverter(k Kind, converter func(interface{}) (interface{}, error)) {
	converters[k] = converter
}

// Convert calls the converter of the kind k to convert the value v.
//
// If the converter of the kind k does not exists, it returns ErrKindNotExist.
func Convert(k Kind, v interface{}) (interface{}, error) {
	if c, ok := converters[k]; ok {
		return c(v)
	}
	return nil, ErrKindNotExist
}

// Converter is used to convert the value by the Scan method. So you can use it
// as the argument of Rows.Scan() in sql.
type Converter struct {
	kind  Kind
	value interface{}
}

// NewConverter returns a Converter to convert a value to the type kind.
func NewConverter(kind Kind) Converter {
	return Converter{kind: kind}
}

// Scan converts the value src.
func (c *Converter) Scan(src interface{}) error {
	value, err := Convert(c.kind, src)
	if err != nil {
		return err
	}
	c.value = value
	return nil
}

// Value returns the inner converted result.
func (c Converter) Value() interface{} {
	return c.value
}
