package pool

import "sync"

// The []byte wrapper of sync.Pool
type BufPool struct {
	size int
	pool *sync.Pool
}

// New a buffer pool.
//
// size is the size of the buffer.
func NewBufPool(size int) *BufPool {
	_pool := &sync.Pool{New: func() interface{} {
		return make([]byte, size)
	}}
	return &BufPool{pool: _pool, size: size}
}

// Get a buffer.
func (p *BufPool) Get() []byte {
	x := p.pool.Get()
	if x == nil {
		return make([]byte, p.size)
	}
	return x.([]byte)
}

// Put a buffer to the pool.
func (p *BufPool) Put(buf []byte) {
	p.pool.Put(buf)
}
