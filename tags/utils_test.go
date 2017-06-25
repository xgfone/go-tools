package tags

import (
	"fmt"
)

func ExampleGetAllTags() {
	tags := GetAllTags(` tag1:"'123'" tag2:"456" \n`)
	fmt.Printf("tag1=|%s|\n", tags["tag1"])
	fmt.Printf("tag2=|%s|\n", tags["tag2"])

	// Output:
	// tag1=|'123'|
	// tag2=|456|
	//
}
