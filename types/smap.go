package types

import (
	"fmt"
	"time"
)

// SMap is the proxy of map[string]interface{}.
type SMap map[string]interface{}

// NewSMap returns a new SMap with the capacity.
func NewSMap(capacity int) SMap {
	return make(SMap, capacity)
}

// ToMap predicates itself to map[string]interface{}.
func (m SMap) ToMap() map[string]interface{} {
	return map[string]interface{}(m)
}

// Get returns the value by the key, which will return the default instead
// if no the key and panic if the type of the vlaue is not string.
func (m SMap) Get(key string, _default interface{}) interface{} {
	if v, ok := m[key]; ok {
		return v.(interface{})
	}
	return _default
}

// GetString returns the string value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not string.
func (m SMap) GetString(key string, _default string) string {
	if v, ok := m[key]; ok {
		return v.(string)
	}
	return _default
}

// GetBool returns the bool value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not bool.
func (m SMap) GetBool(key string, _default bool) bool {
	if v, ok := m[key]; ok {
		return v.(bool)
	}
	return _default
}

// GetInt returns the int value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not int.
func (m SMap) GetInt(key string, _default int) int {
	if v, ok := m[key]; ok {
		return v.(int)
	}
	return _default
}

// GetInt8 returns the int8 value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not int8.
func (m SMap) GetInt8(key string, _default int8) int8 {
	if v, ok := m[key]; ok {
		return v.(int8)
	}
	return _default
}

// GetInt16 returns the int16 value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not int16.
func (m SMap) GetInt16(key string, _default int16) int16 {
	if v, ok := m[key]; ok {
		return v.(int16)
	}
	return _default
}

// GetInt32 returns the int32 value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not int32.
func (m SMap) GetInt32(key string, _default int32) int32 {
	if v, ok := m[key]; ok {
		return v.(int32)
	}
	return _default
}

// GetInt64 returns the int64 value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not int64.
func (m SMap) GetInt64(key string, _default int64) int64 {
	if v, ok := m[key]; ok {
		return v.(int64)
	}
	return _default
}

// GetUint returns the uint value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not uint.
func (m SMap) GetUint(key string, _default uint) uint {
	if v, ok := m[key]; ok {
		return v.(uint)
	}
	return _default
}

// GetUint8 returns the uint8 value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not uint8.
func (m SMap) GetUint8(key string, _default uint8) uint8 {
	if v, ok := m[key]; ok {
		return v.(uint8)
	}
	return _default
}

// GetUint16 returns the uint16 value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not uint16.
func (m SMap) GetUint16(key string, _default uint16) uint16 {
	if v, ok := m[key]; ok {
		return v.(uint16)
	}
	return _default
}

// GetUint32 returns the uint32 value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not uint32.
func (m SMap) GetUint32(key string, _default uint32) uint32 {
	if v, ok := m[key]; ok {
		return v.(uint32)
	}
	return _default
}

// GetUint64 returns the uint64 value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not uint64.
func (m SMap) GetUint64(key string, _default uint64) uint64 {
	if v, ok := m[key]; ok {
		return v.(uint64)
	}
	return _default
}

// GetFloat32 returns the float32 value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not float32.
func (m SMap) GetFloat32(key string, _default float32) float32 {
	if v, ok := m[key]; ok {
		return v.(float32)
	}
	return _default
}

// GetFloat64 returns the float64 value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not float64.
func (m SMap) GetFloat64(key string, _default float64) float64 {
	if v, ok := m[key]; ok {
		return v.(float64)
	}
	return _default
}

// GetDuration returns the time.Duration value by the key, which will return
// the default instead if no the key and panic if the type of the vlaue is
// not time.Duration.
func (m SMap) GetDuration(key string, _default time.Duration) time.Duration {
	if v, ok := m[key]; ok {
		return v.(time.Duration)
	}
	return _default
}

// GetTime returns the time.Time value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not time.Time.
func (m SMap) GetTime(key string, _default time.Time) time.Time {
	if v, ok := m[key]; ok {
		return v.(time.Time)
	}
	return _default
}

// GetInterfaceSlice returns the []interface{} value by the key, which will
// return the default instead if no the key and panic if the type of the vlaue
// is not []interface{}.
func (m SMap) GetInterfaceSlice(key string, _default []interface{}) []interface{} {
	if v, ok := m[key]; ok {
		return v.([]interface{})
	}
	return _default
}

// GetStringSlice returns the []string value by the key, which will return
// the default instead if no the key and panic if the type of the vlaue is
// not []string.
func (m SMap) GetStringSlice(key string, _default []string) []string {
	if v, ok := m[key]; ok {
		return v.([]string)
	}
	return _default
}

// GetIntSlice returns the []int value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not []int.
func (m SMap) GetIntSlice(key string, _default []int) []int {
	if v, ok := m[key]; ok {
		return v.([]int)
	}
	return _default
}

// GetUintSlice returns the []uint value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not []uint.
func (m SMap) GetUintSlice(key string, _default []uint) []uint {
	if v, ok := m[key]; ok {
		return v.([]uint)
	}
	return _default
}

// GetFloat64Slice returns the []float64 value by the key, which will return
// the default instead if no the key and panic if the type of the vlaue is
// not []float64.
func (m SMap) GetFloat64Slice(key string, _default []float64) []float64 {
	if v, ok := m[key]; ok {
		return v.([]float64)
	}
	return _default
}

// GetStringMap returns the SMap value by the key, which will return the default
// instead if no the key and panic if the type of the vlaue is not SMap or
// map[string]interface{}.
func (m SMap) GetStringMap(key string, _default SMap) SMap {
	if v, ok := m[key]; ok {
		switch m := v.(type) {
		case SMap:
			return m
		case map[string]interface{}:
			return m
		default:
			panic(fmt.Errorf("unknown type '%T' to map[string]interface{}", v))
		}
	}
	return _default
}

/////////////////////////////////////////////////////////////////////////////

// Must is the same as Get(key), but panic if no the key.
func (m SMap) Must(key string) interface{} {
	if v, ok := m[key]; ok {
		return v
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustString is the same as GetString, but panic if no the key.
func (m SMap) MustString(key string) string {
	if v, ok := m[key]; ok {
		return v.(string)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustBool is the same as GetBool, but panic if no the key.
func (m SMap) MustBool(key string) bool {
	if v, ok := m[key]; ok {
		return v.(bool)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustInt is the same as GetInt, but panic if no the key.
func (m SMap) MustInt(key string) int {
	if v, ok := m[key]; ok {
		return v.(int)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustInt8 is the same as GetInt8, but panic if no the key.
func (m SMap) MustInt8(key string) int8 {
	if v, ok := m[key]; ok {
		return v.(int8)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustInt16 is the same as GetInt16, but panic if no the key.
func (m SMap) MustInt16(key string) int16 {
	if v, ok := m[key]; ok {
		return v.(int16)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustInt32 is the same as GetInt32, but panic if no the key.
func (m SMap) MustInt32(key string) int32 {
	if v, ok := m[key]; ok {
		return v.(int32)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustInt64 is the same as GetInt64, but panic if no the key.
func (m SMap) MustInt64(key string) int64 {
	if v, ok := m[key]; ok {
		return v.(int64)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustUint is the same as GetUint, but panic if no the key.
func (m SMap) MustUint(key string) uint {
	if v, ok := m[key]; ok {
		return v.(uint)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustUint8 is the same as GetUint8, but panic if no the key.
func (m SMap) MustUint8(key string) uint8 {
	if v, ok := m[key]; ok {
		return v.(uint8)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustUint16 is the same as GetUint16, but panic if no the key.
func (m SMap) MustUint16(key string) uint16 {
	if v, ok := m[key]; ok {
		return v.(uint16)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustUint32 is the same as GetUint32, but panic if no the key.
func (m SMap) MustUint32(key string) uint32 {
	if v, ok := m[key]; ok {
		return v.(uint32)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustUint64 is the same as GetUint64, but panic if no the key.
func (m SMap) MustUint64(key string) uint64 {
	if v, ok := m[key]; ok {
		return v.(uint64)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustFloat32 is the same as GetFloat32, but panic if no the key.
func (m SMap) MustFloat32(key string) float32 {
	if v, ok := m[key]; ok {
		return v.(float32)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustFloat64 is the same as GetFloat64, but panic if no the key.
func (m SMap) MustFloat64(key string) float64 {
	if v, ok := m[key]; ok {
		return v.(float64)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustDuration is the same as GetDuration, but panic if no the key.
func (m SMap) MustDuration(key string) time.Duration {
	if v, ok := m[key]; ok {
		return v.(time.Duration)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustTime is the same as GetTime, but panic if no the key.
func (m SMap) MustTime(key string) time.Time {
	if v, ok := m[key]; ok {
		return v.(time.Time)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustInterfaceSlice is the same as GetInterfaceSlice, but panic if no the key.
func (m SMap) MustInterfaceSlice(key string) []interface{} {
	if v, ok := m[key]; ok {
		return v.([]interface{})
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustStringSlice is the same as GetStringSlice, but panic if no the key.
func (m SMap) MustStringSlice(key string) []string {
	if v, ok := m[key]; ok {
		return v.([]string)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustIntSlice is the same as GetIntSlice, but panic if no the key.
func (m SMap) MustIntSlice(key string) []int {
	if v, ok := m[key]; ok {
		return v.([]int)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustUintSlice is the same as GetUintSlice, but panic if no the key.
func (m SMap) MustUintSlice(key string) []uint {
	if v, ok := m[key]; ok {
		return v.([]uint)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustFloat64Slice is the same as GetFloat64Slice, but panic if no the key.
func (m SMap) MustFloat64Slice(key string) []float64 {
	if v, ok := m[key]; ok {
		return v.([]float64)
	}
	panic(fmt.Errorf("no map key '%s'", key))
}

// MustStringMap is the same as GetStringMap, but panic if no the key.
func (m SMap) MustStringMap(key string) SMap {
	if v, ok := m[key]; ok {
		switch m := v.(type) {
		case SMap:
			return m
		case map[string]interface{}:
			return m
		default:
			panic(fmt.Errorf("unknown type '%T' to map[string]interface{}", v))
		}
	}
	panic(fmt.Errorf("no map key '%s'", key))
}
