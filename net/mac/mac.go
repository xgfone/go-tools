// Standardize the mac address.
package mac

import (
	"fmt"
	"strconv"
	"strings"
)

type Standard struct {
	// If true, output the upper character, such as AA:BB:11:22:33:44.
	// Or, output the lower character.
	Upper bool

	// If true, pad with leading zero, such as 01:02:03:04:05:06.
	Unified bool
}

var (
	StandardUU = NewStandard(true, true)
	StandardUu = NewStandard(true, false)
	StandarduU = NewStandard(false, true)
	Standarduu = NewStandard(false, false)
)

func NewStandard(upper, unified bool) Standard {
	return Standard{Upper: upper, Unified: unified}
}

// Convert the argument of mac to the specifical standard mac address.
//
// Return the empty string if the argument of mac is not the legal mac address.
func (m Standard) Standardize(mac string) string {
	macs := strings.Split(mac, ":")
	if len(macs) != 6 {
		return ""
	}

	width := ""
	upper := "x"
	if m.Upper {
		upper = "X"
	}
	if m.Unified {
		width = "2"
	}
	formatter := fmt.Sprintf("%%0%s%s", width, upper)

	for i := 0; i < 6; i++ {
		if _v, err := strconv.ParseUint(macs[i], 16, 64); err != nil {
			return ""
		} else {
			macs[i] = fmt.Sprintf(formatter, _v)
		}
	}

	return strings.Join(macs, ":")
}

// Same as NewStandard(true, true).Standardize(mac)
func StandardizeUU(mac string) string {
	return StandardUU.Standardize(mac)
}

// Same as NewStandard(true, false).Standardize(mac)
func StandardizeUu(mac string) string {
	return StandardUu.Standardize(mac)
}

// Same as NewStandard(false, true).Standardize(mac)
func StandardizeuU(mac string) string {
	return StandarduU.Standardize(mac)
}

// Same as NewStandard(false, false).Standardize(mac)
func Standardizeuu(mac string) string {
	return Standarduu.Standardize(mac)
}
