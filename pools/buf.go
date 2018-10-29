package pools

import (
	"bytes"
	"sync"
)

var (
	// DefaultBufferPool is the default global buffer pool.
	DefaultBufferPool = NewBufferPool()

	// BytesPool1K the bytes buffer with 1K buffer.
	BytesPool1K = NewBytesPool(1024)

	// BytesPool2K the bytes buffer with 2K buffer.
	BytesPool2K = NewBytesPool(2048)

	// BytesPool4K the bytes buffer with 4K buffer.
	BytesPool4K = NewBytesPool(4096)

	// BytesPool8K the bytes buffer with 8K buffer.
	BytesPool8K = NewBytesPool(8192)
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
	size int
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

	pool := &sync.Pool{New: func() interface{} {
		return makeBuffer(size)
	}}
	return BufferPool{pool: pool, size: size}
}

// Get returns a bytes.Buffer.
func (p BufferPool) Get() *bytes.Buffer {
	x := p.pool.Get()
	if x == nil {
		return makeBuffer(p.size)
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
