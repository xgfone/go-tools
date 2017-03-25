package mac_test

import (
	"fmt"
	"testing"

	"github.com/xgfone/go-tools/nets/mac"
)

func TestValid(t *testing.T) {
	mac1 := "Aa:bB:01:2:00:0"
	if mac.StandardizeUU(mac1) != "AA:BB:01:02:00:00" {
		t.Fail()
	}
	if mac.StandardizeUu(mac1) != "AA:BB:1:2:0:0" {
		t.Fail()
	}
	if mac.StandardizeuU(mac1) != "aa:bb:01:02:00:00" {
		t.Fail()
	}
	if mac.Standardizeuu(mac1) != "aa:bb:1:2:0:0" {
		t.Fail()
	}

	mac2 := "AA:BB"
	if mac.StandardizeUU(mac2) != "" {
		t.Fail()
	}

	mac3 := "Aa:ZZ:00:00:00:00"
	if mac.StandardizeUU(mac3) != "" {
		t.Fail()
	}
}

func ExampleStandard_Standardize() {
	standard := mac.NewStandard(true, true)
	fmt.Println(standard.Standardize("Aa:bB:01:2:00:0"))
	// Output:
	// AA:BB:01:02:00:00
}
