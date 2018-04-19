package pools

import (
	"bytes"
	"sync"
)

// BytesPool is the []byte wrapper of sync.Pool
type BytesPool struct {
	size int
	pool *sync.Pool
}

// BufPool is the alias of BytesPool for backward compatibility.
type BufPool = BytesPool

// NewBufPool is for backward compatibility.
func NewBufPool(size int) BufPool {
	return NewBytesPool(size)
}

// NewBytesPool returns a new []byte pool.
//
// size is the size of the []byte.
func NewBytesPool(size int) BytesPool {
	pool := &sync.Pool{New: func() interface{} {
		return make([]byte, size)
	}}
	return BytesPool{pool: pool, size: size}
}

// Get returns a []byte.
func (p BytesPool) Get() []byte {
	x := p.pool.Get()
	if x == nil {
		return make([]byte, p.size)
	}
	return x.([]byte)
}

// Put places a []byte to the pool.
func (p BytesPool) Put(b []byte) {
	if b != nil {
		p.pool.Put(b)
	}
}

// BufferPool is the bytes.Buffer wrapper of sync.Pool
type BufferPool struct {
	pool *sync.Pool
}

func makeBuffer(size int) (b *bytes.Buffer) {
	b = bytes.NewBuffer(make([]byte, size))
	b.Reset()
	return
}

// NewBufferPool returns a new bytes.Buffer pool.
func NewBufferPool() BufferPool {
	bp := BufferPool{}
	pool := &sync.Pool{New: func() interface{} {
		return makeBuffer(1024)
	}}
	bp.pool = pool
	return bp
}

// Get returns a bytes.Buffer.
func (p BufferPool) Get() *bytes.Buffer {
	x := p.pool.Get()
	if x == nil {
		return makeBuffer(1024)
	}
	return x.(*bytes.Buffer)
}

// Put places a bytes.Buffer to the pool.
func (p BufferPool) Put(b *bytes.Buffer) {
	if b != nil {
		b.Reset()
		p.pool.Put(b)
	}
}
