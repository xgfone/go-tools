package option

// Interface is an common Option interface, which is used to denote all the types of Option.
type Interface interface {
	IsSome() bool
	IsNone() bool

	// Value returns the inner value. Return nil if it's a None.
	Value() interface{}

	// Some returns the inner value, but panic if it's a None.
	Some() interface{}

	// None check whether the inner value is None and panic if it's not a None.
	None()

	// Reset resets the inner value to value.
	Reset(value interface{})

	// Scan parses src and assigns to iteself.
	Scan(src interface{}) error

	// ConvertTo converts the value by convert then assigns the result to the inner.
	ConvertTo(value interface{}, convert func(interface{}) (interface{}, error)) error

	// SomeOr returns the inner value if it's not None. Or return v.
	SomeOr(v interface{}) interface{}

	// String implements the interface fmt.Stringer.
	String() string

	// Return the inner value with the specific type. Or panic if the type is not right.
	Str() string
	Bytes() []byte
	Bool() bool
	Byte() byte
	Rune() rune
	Int() int
	Int8() int8
	Int16() int16
	Int32() int32
	Int64() int64
	Uint() uint
	Uint8() uint8
	Uint16() uint16
	Uint32() uint32
	Uint64() uint64
	Float32() float32
	Float64() float64
	Strs() []string
	Interfaces() []interface{}
	Map() map[string]interface{}
	StrMap() map[string]string

	// Report whether the inner value is the corresponding type.
	IsString() bool
	IsBytes() bool
	IsBool() bool
	IsInt() bool
	IsInt8() bool
	IsInt16() bool
	IsInt32() bool
	IsInt64() bool
	IsUint() bool
	IsUint8() bool
	IsUint16() bool
	IsUint32() bool
	IsUint64() bool
	IsFloat32() bool
	IsFloat64() bool
	IsSignedInteger() bool
	IsUnsignedInteger() bool
	IsInteger() bool
	IsFloat() bool
	IsNumber() bool
	IsStrs() bool
	IsInterfaces() bool
	IsMap() bool
	IsStrMap() bool

	// Convert the inner value to the specific type.
	ToString() (string, error)
	ToBool() (bool, error)
	ToInt() (int, error)
	ToInt8() (int8, error)
	ToInt16() (int16, error)
	ToInt32() (int32, error)
	ToInt64() (int64, error)
	ToUint() (uint, error)
	ToUint8() (uint8, error)
	ToUint16() (uint16, error)
	ToUint32() (uint32, error)
	ToUint64() (uint64, error)
	ToFloat32() (float32, error)
	ToFloat64() (float64, error)
}
