// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
