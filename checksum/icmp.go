// Package checksum supplies some ways to calculate the checksum.
package checksum

// ICMP calculates the checksum of ICMP package.
func ICMP(data []byte) uint16 {
	var (
		sum    uint32
		index  int
		length = len(data)
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}
