package convert

import (
	"encoding/binary"
)

type Number interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | int | uint
}

// eg: ConvertToBytes(1) -> []byte{0, 0, 0, 1}
func ConvertToBytes[T Number](num T) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(num))
	return buf
}

// eg: ParseBytes([]byte{0, 0, 0, 1}) -> 1
func ParseBytes[T Number](buf []byte) T {
	// ljust 8 bytes
	if len(buf) < 8 {
		var newbuf = make([]byte, 8)
		copy(newbuf, buf)
		buf = newbuf
	}
	num := T(binary.LittleEndian.Uint64(buf))
	return num
}
