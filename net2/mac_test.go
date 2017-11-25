package net2

import (
	"fmt"
)

func ExmapleNormalizeMac() {
	mac := "Aa:bB:01:2:00:0"
	fmt.Println(NormalizeMac(mac, true, true))
	fmt.Println(NormalizeMac(mac, true, false))
	fmt.Println(NormalizeMac(mac, false, true))
	fmt.Println(NormalizeMac(mac, false, false))

	fmt.Println(NormalizeMac("AA:BB", false, false))
	fmt.Println(NormalizeMac("Aa:ZZ:01:02:03:0467", true, true))

	// Output:
	// AA:BB:01:02:00:00
	// aa:bb:01:02:00:00
	// AA:BB:1:2:0:0
	// aa:bb:1:2:0:0
	//
	//
	// aa
}

func ExmapleNormalizeMacFU() {
	fmt.Println(NormalizeMacFU("Aa:bB:01:2:00:0"))

	// Output:
	// AA:BB:01:02:00:00
}

func ExmapleNormalizeMacFu() {
	fmt.Println(NormalizeMacFu("Aa:bB:01:2:00:0"))

	// Output:
	// aa:bb:01:02:00:00
}

func ExmapleNormalizeMacfU() {
	fmt.Println(NormalizeMacfU("Aa:bB:01:2:00:0"))

	// Output:
	// AA:BB:1:2:0:0
}

func ExmapleNormalizeMacfu() {
	fmt.Println(NormalizeMacfu("Aa:bB:01:2:00:0"))

	// Output:
	// aa:bb:1:2:0:0
}
