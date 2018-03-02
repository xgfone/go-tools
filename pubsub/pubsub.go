package pubsub

import (
	"container/list"
	"fmt"
	"strings"
	"sync"
	"time"
)

// PubSub is an Publish-Subcribe interface.
type PubSub interface {
	// Publish publishes a msg to a topic.
	//
	// It does not return until all the subcribers are notified or there is
	// an error.
	Publish(topic string, msg interface{}) (err error)

	// Subscribe subscribes a topic named pattern with a callback function cb,
	// and returns a unsubscribe function and an error.
	//
	// When the message is published, the callback will be called with the message
	// if the topic matches the pattern.
	//
	// If the pattern ends with "*", it enable the fuzzy matching, that's,
	// it will match all that starts with "pattern" except "*".
	Subcribe(pattern string, cb func(interface{})) (unsubscribe func() error, err error)
}

type patternCB struct {
	pattern  string
	isExact  bool
	callback func(interface{})
}

type memoryPubSub struct {
	sync.Mutex
	orders    *list.List
	subcribes map[string]patternCB
}

// NewMemoryPubSub returns a new PubSub implementation based on the memory.
func NewMemoryPubSub() PubSub {
	return &memoryPubSub{
		orders:    list.New(),
		subcribes: make(map[string]patternCB, 8),
	}
}

func (m *memoryPubSub) Publish(topic string, msg interface{}) error {
	if topic == "" {
		panic(fmt.Errorf("The pattern cannot be empty"))
	}

	m.Lock()
	defer m.Unlock()

	for element := m.orders.Front(); element != nil; element = element.Next() {
		pcb := m.subcribes[element.Value.(string)]

		if pcb.isExact {
			if pcb.pattern == topic {
				pcb.callback(msg)
			}
		} else if strings.HasPrefix(topic, pcb.pattern) {
			pcb.callback(msg)
		}
	}
	return nil
}

func (m *memoryPubSub) Subcribe(pattern string, cb func(interface{})) (func() error, error) {
	if pattern == "" {
		panic(fmt.Errorf("The pattern cannot be empty"))
	}

	key := time.Now().String()
	isExact := true
	maxLen := len(pattern) - 1
	if pattern[maxLen] == '*' {
		isExact = false
		pattern = pattern[:maxLen]
	}

	m.Lock()
	element := m.orders.PushBack(key)
	m.subcribes[key] = patternCB{pattern: pattern, isExact: isExact, callback: cb}
	m.Unlock()

	unsubscribe := func() error {
		m.Lock()
		delete(m.subcribes, key)
		m.orders.Remove(element)
		m.Unlock()
		return nil
	}

	return unsubscribe, nil
}
