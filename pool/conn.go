package pool

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/xgfone/go-tools/io2"
)

// AddrTCPConnPool is the connection pool based on the address, that's, when you
// need a connection, you only get the connection by the address.
type AddrTCPConnPool struct {
	lock *sync.Mutex

	size  int
	pools map[string]*ResourcePool

	// The connection timeout, and the default is 3s.
	timeout time.Duration
}

// NewAddrTCPConnPool returns a new AddrTCPConnPool.
//
// options supports two arguments: size and timeout. size is 1 by default, and
// timeout is 3 by default, the unit of which is second. So you can call it
// by the three ways below:
//
//    NewAddrTCPConnPool()
//    NewAddrTCPConnPool(size)
//    NewAddrTCPConnPool(size, timeout)
//
// Notice: the first argument is size, and the second is timeout.
func NewAddrTCPConnPool(options ...int) AddrTCPConnPool {
	size := 1
	timeout := 3

	_len := len(options)
	if _len == 1 {
		size = options[0]
	} else if _len > 1 {
		size = options[0]
		timeout = options[1]
	}

	return AddrTCPConnPool{
		size:    size,
		lock:    new(sync.Mutex),
		pools:   make(map[string]*ResourcePool),
		timeout: time.Duration(timeout) * time.Second,
	}
}

// Put places the TCP connection into the pool relating to the addr.
//
// For every successful Get, a corresponding Put is required. If you no longer
// need a resource, you will need to call Put(nil) instead of returning the
// closed resource.
func (p AddrTCPConnPool) Put(addr string, c *net.TCPConn) {
	p.lock.Lock()
	rp := p.pools[addr]
	p.lock.Unlock()
	rp.Put(io2.NewClose(c))
}

// Get returns a TCP connection by the addr from the pool.
func (p AddrTCPConnPool) Get(addr string) (c *net.TCPConn, err error) {
	var rp *ResourcePool

	p.lock.Lock()
	rp, ok := p.pools[addr]
	if !ok {
		rp = NewResourcePool(func() (Resource, error) {
			c, err := net.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			return io2.NewClose(c.(*net.TCPConn)), nil
		}, p.size, p.size, p.timeout)
		p.pools[addr] = rp
	}
	p.lock.Unlock()

	r, err := rp.Get(context.TODO())
	if err != nil {
		return
	}
	c = r.(io2.Close).Value.(*net.TCPConn)
	return
}
