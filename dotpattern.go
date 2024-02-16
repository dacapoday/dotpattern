package dotpattern

import (
	"strings"
)

func ReorderBits(src byte, order [8]uint8) (dst byte) {
	for i, pos := range order {
		if val := (src >> i) & 1; val > 0 {
			dst |= val << pos
		}
	}
	return
}

func Byte(src byte, order [8]uint8) rune {
	return rune(0x2800) + rune(ReorderBits(src, order))
}

func Bytes(src []byte, order [8]uint8) string {
	var buffer strings.Builder
	buffer.Grow(len(src) * 4)
	for _, v := range src {
		buffer.WriteRune(Byte(v, order))
	}
	return buffer.String()
}

var DefaultOrder = [8]uint8{0, 1, 2, 6, 3, 4, 5, 7}
