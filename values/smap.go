package values

import "reflect"

// SMap is a type based on the string key.
type SMap map[string]interface{}

// ToSMap converts the type of map[string]interface{} or SMap to SMap.
//
// Return nil if the type is not either.
func ToSMap(v interface{}) SMap {
	switch v.(type) {
	case map[string]interface{}:
		return SMap(v.(map[string]interface{}))
	case SMap:
		return v.(SMap)
	default:
		return nil
	}
}

// ConvertToSMap converts any map, whose key is the type of string, to SMap.
//
// Return nil if it's not a map, or it's nil or has no elements.
func ConvertToSMap(v interface{}) SMap {
	_v := reflect.ValueOf(v)
	if !_v.IsValid() || _v.Kind() != reflect.Map {
		return nil
	}

	results := make(SMap, _v.Len())
	for _, key := range _v.MapKeys() {
		if key.Kind() != reflect.String {
			return nil
		}
		results[key.String()] = _v.MapIndex(key).Interface()
	}
	return results
}
