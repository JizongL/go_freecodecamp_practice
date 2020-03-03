package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	// print memory address
	fmt.Println(&a, b)
	fmt.Println(*&a, *b)
}
