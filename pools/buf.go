// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pools

import (
	"bytes"
	"sync"
)

var (
	// BufferPool1k is the buffer pool with 1K initialized capacity.
	BufferPool1k = NewBufferPool(1024)

	// BufferPool2k is the buffer pool with 2K initialized capacity.
	BufferPool2k = NewBufferPool(2048)

	// BufferPool4k is the buffer pool with 4K initialized capacity.
	BufferPool4k = NewBufferPool(4096)

	// BufferPool8k is the buffer pool with 8K initialized capacity.
	BufferPool8k = NewBufferPool(8192)

	// BytesPool1K the bytes pool with 1K buffer.
	BytesPool1K = NewBytesPool(1024)

	// BytesPool2K the bytes pool with 2K buffer.
	BytesPool2K = NewBytesPool(2048)

	// BytesPool4K the bytes pool with 4K buffer.
	BytesPool4K = NewBytesPool(4096)

	// BytesPool8K the bytes pool with 8K buffer.
	BytesPool8K = NewBytesPool(8192)
)

// BytesPool is the []byte wrapper of sync.Pool
type BytesPool struct {
	size int
	pool sync.Pool
}

// NewBytesPool returns a new []byte pool.
//
// size is the size of the []byte.
func NewBytesPool(size int) *BytesPool {
	newf := func() interface{} { return make([]byte, size) }
	return &BytesPool{pool: sync.Pool{New: newf}, size: size}
}

// Get returns a []byte.
func (p *BytesPool) Get() []byte {
	x := p.pool.Get()
	if x == nil {
		return make([]byte, p.size)
	}
	return x.([]byte)
}

// Put places a []byte to the pool.
func (p *BytesPool) Put(b []byte) {
	if b != nil {
		p.pool.Put(b)
	}
}

// BufferPool is the bytes.Buffer wrapper of sync.Pool
type BufferPool struct {
	pool sync.Pool
	size int
}

func makeBuffer(size int) (b *bytes.Buffer) {
	b = bytes.NewBuffer(make([]byte, size))
	b.Reset()
	return
}

// NewBufferPool returns a new bytes.Buffer pool.
func NewBufferPool(size int) *BufferPool {
	newf := func() interface{} { return makeBuffer(size) }
	return &BufferPool{pool: sync.Pool{New: newf}, size: size}
}

// Get returns a bytes.Buffer.
func (p *BufferPool) Get() *bytes.Buffer {
	x := p.pool.Get()
	if x == nil {
		return makeBuffer(p.size)
	}
	return x.(*bytes.Buffer)
}

// Put places a bytes.Buffer to the pool.
func (p *BufferPool) Put(b *bytes.Buffer) {
	if b != nil {
		b.Reset()
		p.pool.Put(b)
	}
}
