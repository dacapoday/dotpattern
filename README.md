# dotpattern

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dacapoday/dotpattern)

Implement dot matrix pattern effect based on [Braille characters](https://en.wikipedia.org/wiki/Braille_Patterns).

## Example
```go
package main

import (
	"fmt"

	"github.com/dacapoday/dotpattern"
)

func bitmap(data []byte) string {
	return dotpattern.Bytes(data, dotpattern.DefaultOrder)
}

func main() {
	test := []byte{1, 2, 4, 8, 16, 32, 64, 128, 15, 240, 102, 6, 96}
	result := bitmap(test)
	fmt.Println("[", result, "]")
	// Output: [ ⠁⠂⠄⡀⠈⠐⠠⢀⡇⢸⠶⠆⠰ ]
}

```