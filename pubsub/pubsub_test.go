package pubsub

import (
	"fmt"
)

func ExamplePubSub() {
	ps := NewMemoryPubSub()

	u1, _ := ps.Subcribe("topic", func(msg interface{}) {
		fmt.Printf("topic: %v\n", msg)
	})
	defer u1()

	u2, _ := ps.Subcribe("topic*", func(msg interface{}) {
		fmt.Printf("topic*: %v\n", msg)
	})
	defer u2()

	ps.Publish("topic", "message")
	ps.Publish("non-topic", "non-message")

	// Output:
	// topic: message
	// topic*: message
}
