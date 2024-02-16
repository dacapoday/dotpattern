// Package dotpattern provides methods to convert bytes to dot matrix pattern.
package dotpattern

import (
	"strings"
)

// ReorderBits adjust the bits of a byte in the specified order.
func ReorderBits(src byte, order [8]uint8) (dst byte) {
	for i, pos := range order {
		if val := (src >> i) & 1; val > 0 {
			dst |= val << pos
		}
	}
	return
}

// Byte convert byte to braille character.
func Byte(src byte, order [8]uint8) rune {
	return rune(0x2800) + rune(ReorderBits(src, order))
}

// Bytes convert byte slice to dot matrix.
func Bytes(src []byte, order [8]uint8) string {
	var buffer strings.Builder
	buffer.Grow(len(src) * 4)
	for _, v := range src {
		buffer.WriteRune(Byte(v, order))
	}
	return buffer.String()
}

// DefaultOrder is ↓↓
var DefaultOrder = [8]uint8{0, 1, 2, 6, 3, 4, 5, 7}

// UshapeOrder is ↓↑
var UshapeOrder = [8]uint8{0, 1, 2, 6, 7, 5, 4, 3}
