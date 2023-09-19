package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	fmt.Println(1 + 1)
	fmt.Println(getNumber())
}

func getNumber() int {
	return 1 << 2
}
