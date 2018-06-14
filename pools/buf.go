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
//
// bufSize is the initializing size of the buffer. If the size is equal to
// or less than 0, it will be ignored, and use the default size, 1024.
func NewBufferPool(bufSize ...int) BufferPool {
	size := 1024
	if len(bufSize) > 0 && bufSize[0] > 0 {
		size = bufSize[0]
	}

	bp := BufferPool{}
	pool := &sync.Pool{New: func() interface{} {
		return makeBuffer(size)
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
