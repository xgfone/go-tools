package values

type SMap map[string]interface{}

func ToSMap(v interface{}) (SMap, error) {
	switch v.(type) {
	case map[string]interface{}:
		return SMap(v.(map[string]interface{})), nil
	case SMap:
		return v.(SMap), nil
	default:
		return nil, ErrTypeOrIndex
	}
}
