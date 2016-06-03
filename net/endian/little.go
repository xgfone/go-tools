package endian

var Little littleEndian

type littleEndian struct{}

func (e littleEndian) From16(v uint16) []byte {
	b := make([]byte, 2)

	b[0] = byte(v)
	b[1] = byte(v >> 8)

	return b
}

func (e littleEndian) From32(v uint32) []byte {
	b := make([]byte, 4)

	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)

	return b
}

func (e littleEndian) From64(v uint64) []byte {
	b := make([]byte, 8)

	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)

	return b
}

func (e littleEndian) To16(b []byte) uint16 {
	if len(b) < 2 {
		return 0
	}

	var v uint16
	v += uint16(b[0])
	v += uint16(b[1]) << 8
	return v
}

func (e littleEndian) To32(b []byte) uint32 {
	if len(b) < 4 {
		return 0
	}

	var v uint32
	v += uint32(b[0])
	v += uint32(b[1]) << 8
	v += uint32(b[2]) << 16
	v += uint32(b[3]) << 24
	return v
}

func (e littleEndian) To64(b []byte) uint64 {
	if len(b) < 8 {
		return 0
	}

	var v uint64
	v += uint64(b[0])
	v += uint64(b[1]) << 8
	v += uint64(b[2]) << 16
	v += uint64(b[3]) << 24
	v += uint64(b[4]) << 32
	v += uint64(b[5]) << 40
	v += uint64(b[6]) << 48
	v += uint64(b[7]) << 56
	return v
}
