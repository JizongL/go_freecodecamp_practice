package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 4000000000.
	// 2f 2 decimal places
	fmt.Printf("%.2fGB", fileSize/GB)
}
