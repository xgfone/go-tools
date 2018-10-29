package types

import (
	"fmt"
	"sync"

	"github.com/xgfone/go-tools/function"
)

// Nil Error
var (
	ErrKeyIsNil   = fmt.Errorf("the key is nil")
	ErrValueIsNil = fmt.Errorf("the value is nil")
)

// Series is used to store the key-value in series.
//
// DEPRECATED!!!
type Series interface {
	// Delete deletes the value by the key.
	// It does not delete the key-value pair in the parent.
	Delete(key interface{})

	// Set sets the key to the value.
	//
	// Notice: both key and value should not be nil. Or it maybe panic.
	Set(key interface{}, value interface{})

	// GetParent returns the parent Series.
	GetParent() Series

	// Get returns the value by the key.
	// If there is not the key in the current Series, it will find the key
	// in the parent. If there is not the key in the parent, it will return nil.
	Get(key interface{}) (value interface{})

	// These methods are the same as Get(), but returns the value of the
	// corresponding type by the key. If the key does not exist, ok is false.
	GetInt(key interface{}) (value int, ok bool)
	GetInt64(key interface{}) (value int64, ok bool)
	GetUint(key interface{}) (value uint, ok bool)
	GetUint64(key interface{}) (value uint64, ok bool)
	GetString(key interface{}) (value string, ok bool)
	GetBool(key interface{}) (value bool, ok bool)

	// These methods are the same as Get(), but returns the value of the
	// corresponding type by the key. If the key does not exist, they returns
	// the default.
	GetIntD(key interface{}, _default int) (value int)
	GetInt64D(key interface{}, _default int64) (value int64)
	GetUintD(key interface{}, _default uint) (value uint)
	GetUint64D(key interface{}, _default uint64) (value uint64)
	GetStringD(key interface{}, _default string) (value string)
	GetBoolD(key interface{}, _default bool) (value bool)

	// These methods are the same as Get(), but returns the value of the
	// corresponding type by the key. If the key does not exist, they will panic.
	MustGetInt(key interface{}) (value int)
	MustGetInt64(key interface{}) (value int64)
	MustGetUint(key interface{}) (value uint)
	MustGetUint64(key interface{}) (value uint64)
	MustGetString(key interface{}) (value string)
	MustGetBool(key interface{}) (value bool)
}

// MergeSeriesKeyP2C merges the values of the key in the Series s from parent
// to children. That's, the value in the parent is before that in the children.
func MergeSeriesKeyP2C(s Series, key interface{}) []interface{} {
	vlaues := MergeSeriesKeyC2P(s, key)
	if len(vlaues) > 1 {
		function.Reverse(vlaues)
	}
	return vlaues
}

// MergeSeriesKeyC2P merges the values of the key in the Series s from children
// to parent. That's, the values in the children are before that in the parent.
func MergeSeriesKeyC2P(s Series, key interface{}) []interface{} {
	var values []interface{}
	for cur := s; cur != nil; cur = cur.GetParent() {
		if value := cur.Get(key); value != nil {
			values = append(values, value)
		}
	}
	return values
}

// NamedSeries is a named Series.
type NamedSeries struct {
	Series
	name string
}

// NewNamedSeries returns a new NamedSeries.
func NewNamedSeries(name string, parent ...Series) NamedSeries {
	s := NewSeries(parent...)
	return NamedSeries{Series: s, name: name}
}

// GetName returns the name of Series.
func (ns NamedSeries) GetName() string {
	return ns.name
}

type series struct {
	maps   *sync.Map
	parent Series
}

// NewSeries returns a new Series.
func NewSeries(parent ...Series) Series {
	var p Series
	if len(parent) > 0 {
		p = parent[0]
	}
	return series{parent: p, maps: new(sync.Map)}
}

func (s series) Delete(key interface{}) {
	s.maps.Delete(key)
}

func (s series) Set(key, value interface{}) {
	if key == nil {
		panic(ErrKeyIsNil)
	}
	if value == nil {
		panic(ErrValueIsNil)
	}
	s.maps.Store(key, value)
}

func (s series) GetParent() Series {
	return s.parent
}

func (s series) Get(key interface{}) interface{} {
	if v, ok := s.maps.Load(key); ok {
		return v
	}
	if s.parent != nil {
		return s.parent.Get(key)
	}
	return nil
}

func (s series) GetInt(key interface{}) (int, bool) {
	if v := s.Get(key); v != nil {
		return v.(int), true
	}
	return 0, false
}

func (s series) GetInt64(key interface{}) (int64, bool) {
	if v := s.Get(key); v != nil {
		return v.(int64), true
	}
	return 0, false
}

func (s series) GetUint(key interface{}) (uint, bool) {
	if v := s.Get(key); v != nil {
		return v.(uint), true
	}
	return 0, false
}

func (s series) GetUint64(key interface{}) (uint64, bool) {
	if v := s.Get(key); v != nil {
		return v.(uint64), true
	}
	return 0, false
}

func (s series) GetString(key interface{}) (string, bool) {
	if v := s.Get(key); v != nil {
		return v.(string), true
	}
	return "", false
}

func (s series) GetBool(key interface{}) (bool, bool) {
	if v := s.Get(key); v != nil {
		return v.(bool), true
	}
	return false, false
}

func (s series) GetIntD(key interface{}, _default int) (value int) {
	if v, ok := s.GetInt(key); ok {
		return v
	}
	return _default
}

func (s series) GetInt64D(key interface{}, _default int64) (value int64) {
	if v, ok := s.GetInt64(key); ok {
		return v
	}
	return _default
}

func (s series) GetUintD(key interface{}, _default uint) (value uint) {
	if v, ok := s.GetUint(key); ok {
		return v
	}
	return _default
}

func (s series) GetUint64D(key interface{}, _default uint64) (value uint64) {
	if v, ok := s.GetUint64(key); ok {
		return v
	}
	return _default
}

func (s series) GetStringD(key interface{}, _default string) (value string) {
	if v, ok := s.GetString(key); ok {
		return v
	}
	return _default
}

func (s series) GetBoolD(key interface{}, _default bool) (value bool) {
	if v, ok := s.GetBool(key); ok {
		return v
	}
	return _default
}

func (s series) MustGetInt(key interface{}) (value int) {
	if v, ok := s.GetInt(key); ok {
		return v
	}
	panic(fmt.Errorf("no key '%v'", key))
}

func (s series) MustGetInt64(key interface{}) (value int64) {
	if v, ok := s.GetInt64(key); ok {
		return v
	}
	panic(fmt.Errorf("no key '%v'", key))
}

func (s series) MustGetUint(key interface{}) (value uint) {
	if v, ok := s.GetUint(key); ok {
		return v
	}
	panic(fmt.Errorf("no key '%v'", key))
}

func (s series) MustGetUint64(key interface{}) (value uint64) {
	if v, ok := s.GetUint64(key); ok {
		return v
	}
	panic(fmt.Errorf("no key '%v'", key))
}

func (s series) MustGetString(key interface{}) (value string) {
	if v, ok := s.GetString(key); ok {
		return v
	}
	panic(fmt.Errorf("no key '%v'", key))
}

func (s series) MustGetBool(key interface{}) (value bool) {
	if v, ok := s.GetBool(key); ok {
		return v
	}
	panic(fmt.Errorf("no key '%v'", key))
}
