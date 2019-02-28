package option

import (
	"github.com/xgfone/go-tools/types"
)

// BoolOption is an Option of the bool type.
type BoolOption struct {
	Interface
}

// NewBoolOption returns a new BoolOption.
func NewBoolOption(o Interface) BoolOption {
	if o == nil {
		o = None()
	}
	return BoolOption{Interface: o}
}

// Scan converts src as bool to the inner value.
func (o BoolOption) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToBool(v) })
}

// StringOption is an Option of the string type.
type StringOption struct {
	Interface
}

// NewStringOption returns a new StringOption.
func NewStringOption(o Interface) StringOption {
	if o == nil {
		o = None()
	}
	return StringOption{Interface: o}
}

// Scan converts src as string to the inner value.
func (o StringOption) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToString(v) })
}

// IntOption is an Option of the int type.
type IntOption struct {
	Interface
}

// NewIntOption returns a new IntOption.
func NewIntOption(o Interface) IntOption {
	if o == nil {
		o = None()
	}
	return IntOption{Interface: o}
}

// Scan converts src as int to the inner value.
func (o IntOption) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToInt(v) })
}

// Int8Option is an Option of the int8 type.
type Int8Option struct {
	Interface
}

// NewInt8Option returns a new Int8Option.
func NewInt8Option(o Interface) Int8Option {
	if o == nil {
		o = None()
	}
	return Int8Option{Interface: o}
}

// Scan converts src as int8 to the inner value.
func (o Int8Option) Scan(src interface{}) error {
	v, err := types.ToInt64(src)
	if err == nil {
		o.Reset(int8(v))
	}
	return err
}

// Int16Option is an Option of the int16 type.
type Int16Option struct {
	Interface
}

// NewInt16Option returns a new Int16Option.
func NewInt16Option(o Interface) Int16Option {
	if o == nil {
		o = None()
	}
	return Int16Option{Interface: o}
}

// Scan converts src as int16 to the inner value.
func (o Int16Option) Scan(src interface{}) error {
	v, err := types.ToInt64(src)
	if err == nil {
		o.Reset(int16(v))
	}
	return err
}

// Int32Option is an Option of the int32 type.
type Int32Option struct {
	Interface
}

// NewInt32Option returns a new Int32Option.
func NewInt32Option(o Interface) Int32Option {
	if o == nil {
		o = None()
	}
	return Int32Option{Interface: o}
}

// Scan converts src as int32 to the inner value.
func (o Int32Option) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToInt32(v) })
}

// Int64Option is an Option of the int64 type.
type Int64Option struct {
	Interface
}

// NewInt64Option returns a new Int64Option.
func NewInt64Option(o Interface) Int64Option {
	if o == nil {
		o = None()
	}
	return Int64Option{Interface: o}
}

// Scan converts src as int64 to the inner value.
func (o Int64Option) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToInt64(v) })
}

// UintOption is an Option of the uint type.
type UintOption struct {
	Interface
}

// NewUintOption returns a new UintOption.
func NewUintOption(o Interface) UintOption {
	if o == nil {
		o = None()
	}
	return UintOption{Interface: o}
}

// Scan converts src as uint to the inner value.
func (o UintOption) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToUint(v) })
}

// Uint8Option is an Option of the uint8 type.
type Uint8Option struct {
	Interface
}

// NewUint8Option returns a new Uint8Option.
func NewUint8Option(o Interface) Uint8Option {
	if o == nil {
		o = None()
	}
	return Uint8Option{Interface: o}
}

// Scan converts src as uint8 to the inner value.
func (o Uint8Option) Scan(src interface{}) error {
	v, err := types.ToUint64(src)
	if err == nil {
		o.Reset(uint8(v))
	}
	return err
}

// Uint16Option is an Option of the uint16 type.
type Uint16Option struct {
	Interface
}

// NewUint16Option returns a new Uint16Option.
func NewUint16Option(o Interface) Uint16Option {
	if o == nil {
		o = None()
	}
	return Uint16Option{Interface: o}
}

// Scan converts src as uint16 to the inner value.
func (o Uint16Option) Scan(src interface{}) error {
	v, err := types.ToUint64(src)
	if err == nil {
		o.Reset(uint16(v))
	}
	return err
}

// Uint32Option is an Option of the uint32 type.
type Uint32Option struct {
	Interface
}

// NewUint32Option returns a new Uint32Option.
func NewUint32Option(o Interface) Uint32Option {
	if o == nil {
		o = None()
	}
	return Uint32Option{Interface: o}
}

// Scan converts src as uint32 to the inner value.
func (o Uint32Option) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToUint32(v) })
}

// Uint64Option is an Option of the uint64 type.
type Uint64Option struct {
	Interface
}

// NewUint64Option returns a new Uint64Option.
func NewUint64Option(o Interface) Uint64Option {
	if o == nil {
		o = None()
	}
	return Uint64Option{Interface: o}
}

// Scan converts src as uint64 to the inner value.
func (o Uint64Option) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToUint64(v) })
}

// Float64Option is an Option of the float64 type.
type Float64Option struct {
	Interface
}

// NewFloat64Option returns a new Float64Option.
func NewFloat64Option(o Interface) Float64Option {
	if o == nil {
		o = None()
	}
	return Float64Option{Interface: o}
}

// Scan converts src as float64 to the inner value.
func (o Float64Option) Scan(src interface{}) error {
	return o.ConvertTo(src, func(v interface{}) (interface{}, error) { return types.ToFloat64(v) })
}
