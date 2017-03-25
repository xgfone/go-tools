package parse

import "fmt"

// String converts v to string.
func String(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
