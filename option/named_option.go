package option

import "fmt"

// NamedOption represents a named Option.
type NamedOption struct {
	Option

	name string
}

// NewNamedOption returns a new NamedOption based on option.
//
// If option is nil, it is None() by default.
func NewNamedOption(name string, option Option) NamedOption {
	if option == nil {
		option = None()
	}
	return NamedOption{name: name, Option: option}
}

// NamedSome returns an NamedOption named name.
//
// If v is nil, it will be a None value.
func NamedSome(name string, v interface{}) NamedOption {
	return NewNamedOption(name, Some(v))
}

// NamedNone is equal to NamedSome(name, nil).
func NamedNone(name string) NamedOption {
	return NamedSome(name, nil)
}

// Name returns the name of the option.
func (o NamedOption) Name() string {
	return o.name
}

// String implements the interface fmt.Stringer.
func (o NamedOption) String() string {
	return fmt.Sprintf("Option(name='%s', value=%v)", o.name, o.Option.Value())
}

// Named returns a proxy to generate some named options with the same name n.
func Named(n string) func(interface{}) NamedOption {
	return func(v interface{}) NamedOption {
		return NamedSome(n, v)
	}
}
