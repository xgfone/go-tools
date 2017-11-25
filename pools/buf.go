package pools

import "sync"

// BufPool is the []byte wrapper of sync.Pool
type BufPool struct {
	size int
	pool *sync.Pool
}

// NewBufPool returns a new buffer pool.
//
// size is the size of the buffer.
func NewBufPool(size int) BufPool {
	pool := &sync.Pool{New: func() interface{} {
		return make([]byte, size)
	}}
	return BufPool{pool: pool, size: size}
}

// Get returns a buffer.
func (p BufPool) Get() []byte {
	x := p.pool.Get()
	if x == nil {
		return make([]byte, p.size)
	}
	return x.([]byte)
}

// Put places a buffer to the pool.
func (p BufPool) Put(buf []byte) {
	p.pool.Put(buf)
}
