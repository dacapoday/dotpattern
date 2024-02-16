package dotpattern

import "fmt"

func ExampleReorderBits() {
	var x byte = 1
	fmt.Println(ReorderBits(x, [8]uint8{7, 1, 2, 3, 4, 5, 6, 0}))
	// Output: 128
}

func ExampleBytes() {
	fmt.Println(Bytes([]byte{1, 2, 3, 4, 8, 16, 32, 64, 255}, DefaultOrder))
	// Output: ⠁⠂⠃⠄⡀⠈⠐⠠⣿
}
