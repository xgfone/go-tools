package values

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
