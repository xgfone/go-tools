package pools

import (
	"context"
	"net"
	"sync"
	"time"
)

type tcpConnResource struct {
	c *net.TCPConn
}

func (t tcpConnResource) Close() {
	t.c.Close()
}

// AddrTCPConnPool is the connection pool based on the address, that's, when you
// need a connection, you only get the connection by the address.
type AddrTCPConnPool struct {
	lock *sync.Mutex

	size  int
	pools map[string]*ResourcePool

	// The connection timeout, and the default is 3s.
	Timeout time.Duration
}

// NewAddrTCPConnPool returns a new AddrTCPConnPool.
func NewAddrTCPConnPool(size int) AddrTCPConnPool {
	return AddrTCPConnPool{
		size:  size,
		lock:  new(sync.Mutex),
		pools: make(map[string]*ResourcePool),

		Timeout: 3 * time.Second,
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

	rp.Put(tcpConnResource{c: c})
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
			return tcpConnResource{c: c.(*net.TCPConn)}, nil
		}, p.size, p.size, p.Timeout)
		p.pools[addr] = rp
	}
	p.lock.Unlock()

	r, err := rp.Get(context.TODO())
	if err != nil {
		return
	}
	c = r.(tcpConnResource).c
	return
}
