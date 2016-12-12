package values

type SMap map[string]interface{}

// ToSMap converts v to the type of SMap.
//
// Return nil if failed.
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
