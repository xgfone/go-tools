package pools

import (
	"fmt"
	"unsafe"
)

func ExampleAddrTCPConnPool() {
	addr := "www.baidu.com:80"
	var v1, v2 uint64

	p := NewAddrTCPConnPool(1)
	c1, err := p.Get(addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	v1 = *(*uint64)(unsafe.Pointer(c1))
	p.Put(addr, c1)

	c2, err := p.Get(addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	v2 = *(*uint64)(unsafe.Pointer(c2))
	p.Put(addr, c2)

	if v1 != 0 && v1 == v2 {
		fmt.Println("OK")
	}

	// Output:
	// OK
}
