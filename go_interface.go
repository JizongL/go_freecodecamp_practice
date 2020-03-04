package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

// only contain behaviors not content.
type Writer interface {
	Write([]byte) (int, error)
}
