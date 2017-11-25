package sort2

import (
	"fmt"
)

func ExampleKvSliceByKey() {
	m := map[string]interface{}{"b": 2, "c": 3, "a": 1}

	kv1 := MapToKvSliceByKey(m)
	Sort(kv1)

	kv2 := MapToKvSliceByKey(m)
	Sort(kv2, true)

	fmt.Printf("%+v\n", kv1)
	fmt.Printf("%+v\n", kv2)

	// Output:
	// [{Key:a Value:1} {Key:b Value:2} {Key:c Value:3}]
	// [{Key:c Value:3} {Key:b Value:2} {Key:a Value:1}]
}

func ExampleKvSliceByValue() {
	m := map[string]interface{}{"b": 2, "c": 3, "a": 1}

	kv1 := MapToKvSliceByValue(m)
	Sort(kv1)

	kv2 := MapToKvSliceByValue(m)
	Sort(kv2, true)

	fmt.Printf("%+v\n", kv1)
	fmt.Printf("%+v\n", kv2)

	// Output:
	// [{Key:a Value:1} {Key:b Value:2} {Key:c Value:3}]
	// [{Key:c Value:3} {Key:b Value:2} {Key:a Value:1}]
}
