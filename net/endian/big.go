// Convert between int/uint and the big/little endian.
package endian

var Big bigEndian

type bigEndian struct{}

func (e bigEndian) From16(v uint16) []byte {
	b := make([]byte, 2)

	b[0] = byte(v >> 8)
	b[1] = byte(v)

	return b
}

func (e bigEndian) From32(v uint32) []byte {
	b := make([]byte, 4)

	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)

	return b
}

func (e bigEndian) From64(v uint64) []byte {
	b := make([]byte, 8)

	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)

	return b
}

func (e bigEndian) To16(b []byte) uint16 {
	if len(b) < 2 {
		return 0
	}

	var v uint16
	v += uint16(b[0]) << 8
	v += uint16(b[1])
	return v
}

func (e bigEndian) To32(b []byte) uint32 {
	if len(b) < 4 {
		return 0
	}

	var v uint32
	v += uint32(b[0]) << 24
	v += uint32(b[1]) << 16
	v += uint32(b[2]) << 8
	v += uint32(b[3])
	return v
}

func (e bigEndian) To64(b []byte) uint64 {
	if len(b) < 8 {
		return 0
	}

	var v uint64
	v += uint64(b[0]) << 56
	v += uint64(b[1]) << 48
	v += uint64(b[2]) << 40
	v += uint64(b[3]) << 32
	v += uint64(b[4]) << 24
	v += uint64(b[5]) << 16
	v += uint64(b[6]) << 8
	v += uint64(b[7])
	return v
}
