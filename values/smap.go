package values

import "reflect"

// SMap is a type based on the string key.
type SMap map[string]interface{}

// NewSMap returns a new SMap.
func NewSMap() SMap {
	return make(SMap)
}

// ToSMap converts the type of map[string]interface{} or SMap to SMap.
//
// Return nil if the type is not either.
func ToSMap(v interface{}) SMap {
	_v, _ := toSMap(v)
	return _v
}

func toSMap(v interface{}) (SMap, bool) {
	switch v.(type) {
	case map[string]interface{}:
		return SMap(v.(map[string]interface{})), true
	case SMap:
		return v.(SMap), true
	default:
		return ConvertToSMap(v)
	}
}

// ConvertToSMap converts any map, whose key is the type of string, to SMap.
//
// Return nil if it's not a map, or it's nil or has no elements.
//
// Notice: SMap(nil) is not a valid SMap.
func ConvertToSMap(v interface{}) (SMap, bool) {
	_v := reflect.ValueOf(v)
	if !_v.IsValid() || _v.Kind() != reflect.Map {
		return nil, false
	}

	results := make(SMap, _v.Len())
	for _, key := range _v.MapKeys() {
		if key.Kind() != reflect.String {
			return nil, false
		}
		results[key.String()] = _v.MapIndex(key).Interface()
	}
	return results, true
}
