package strings2

import (
	"fmt"
)

func ExampleKvFmt() {
	tpl := "{key1}, {key2}, {key1}, {key2}"
	args := map[string]interface{}{"key1": "abc", "key2": "123"}
	fmt.Println(KvFmt(tpl, args))

	// Output:
	// abc, 123, abc, 123
}

func ExampleSetFmtDelimiter() {
	SetFmtDelimiter("{{", "}}")

	tpl := "{{key1}}, {{key2}}, {{key1}}, {{key2}}"
	args := map[string]interface{}{"key1": "abc", "key2": "123"}
	fmt.Println(KvFmt(tpl, args))

	// Output:
	// abc, 123, abc, 123
}
